package binarystar

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

type Daemon struct {
	tree            *Tree
	pending, recent *changeSetBuffer
}

func NewDaemon(ctx context.Context, tree *Tree) (*Daemon, error) {
	d := &Daemon{
		tree:    tree,
		pending: newChangeSetBuffer(),
		recent:  newChangeSetBuffer(),
	}

	go func() {
		for change := range tree.Changes {
			d.pending.Add(change)
			d.recent.Add(change)
		}
	}()

	if err := tree.Watch(ctx); err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Daemon) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Printf("Listening on %s", listener.Addr().String())
	for {
		conn, err := listener.Accept()
		if err != nil {
			return errors.Wrap(err, "Error accepting listener")
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr().String())

		// Setup server side of yamux
		session, err := yamux.Server(conn, nil)
		if err != nil {
			return err
		}

		// Accept a stream that we will use for duplex communication
		stream, err := session.Accept()
		if err != nil {
			log.Printf("Error accepting stream: %v", err)
			continue
		}

		connection := d.createConnection(session, stream)

		if err := connection.Process(); err != nil {
			log.Printf("Connection failed: %v", err)
			continue
		}
	}
}

func (d *Daemon) Connect(address string) error {
	log.Printf("Connecting to tcp://%s", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	// Setup client side of yamux
	session, err := yamux.Client(conn, nil)
	if err != nil {
		return err
	}

	// Open the duplex stream we will use for communication
	stream, err := session.Open()
	if err != nil {
		return err
	}

	connection := d.createConnection(session, stream)

	// Fire off an initial synchronization
	if err := connection.sendPullSync(); err != nil {
		return errors.Wrap(err, "Initial sync failed")
	}

	return connection.Process()
}

func (d *Daemon) createConnection(session *yamux.Session, stream net.Conn) *connection {
	return &connection{
		tree:    d.tree,
		session: session,
		stream:  stream,
		reader:  msgp.NewReader(stream),
		writer:  msgp.NewWriter(stream),
		recent:  d.recent,
		pending: d.pending,
	}
}

type connection struct {
	tree            *Tree
	streamMutex     sync.Mutex
	stream          net.Conn
	session         *yamux.Session
	reader          *msgp.Reader
	writer          *msgp.Writer
	pending, recent *changeSetBuffer
}

func (c *connection) Process() error {
	// process pending changes from the tree
	go func() {
		for {
			c.pending.L.Lock()
			if c.pending.Len() == 0 {
				// blocks until we get changes into the tree, which then get copied into
				// the change buffer
				c.pending.Wait()
			}
			if c.pending.Len() > 0 {
				log.Printf("Found %d changes in pending change buffer", c.pending.Len())
				pending := c.pending.ChangeSet
				c.pending.ChangeSet = ChangeSet{}

				// send the pending changes to the other side
				if err := c.sendPushSync(pending); err != nil {
					log.Printf("pushSync failed: %#v", err)
				}
			}
			c.pending.L.Unlock()
			time.Sleep(time.Second)
		}
	}()

	// receive inbound sync requests
	for {
		var syncReq SyncMessage

		// receive the initial sync message with remote files
		if err := c.receive(&syncReq); err != nil {
			return errors.Wrap(err, "Expected SyncMessage")
		}

		if syncReq.Changes.Len() > 0 {
			log.Printf("Receiving push sync (%d changes)", syncReq.Changes.Len())
			if err := c.receivePushSync(syncReq); err != nil {
				return err
			}
		} else {
			log.Printf("Receiving pull sync (%d files)", len(syncReq.Files))
			if err := c.receivePullSync(syncReq); err != nil {
				return err
			}
		}
	}
}

func (c *connection) sendChangeSet(changes ChangeSet) error {
	t := time.Now()
	log.Printf("Sending changeset: ADD %d, MODIFY %d, DELETE %d",
		len(changes.Add), len(changes.Modify), len(changes.Delete))

	// send ADD changes first, streaming the whole file over the wire
	for _, add := range changes.Add {
		fullPath := filepath.Join(add.Dir, add.Path)

		log.Printf("Streaming %s", fullPath)
		file, err := os.Open(fullPath)
		if err != nil {
			return err
		}
		defer file.Close()

		errCh := c.sendFileStream(file)
		log.Printf("Waiting for stream to finish")
		if err = <-errCh; err != nil {
			return err
		}

		log.Printf("Finished streaming bytes")
	}

	// send MODIFY changes next
	for _, mod := range changes.Modify {
		log.Printf("Sending %d blocks of change", len(mod.Patch.Blocks))
		if err := c.send(&mod.Patch); err != nil {
			return err
		}
	}

	log.Printf("Finished sending changeset in %v", time.Now().Sub(t))
	return nil
}

func (c *connection) receiveChangeSet(changes ChangeSet) error {
	t := time.Now()
	log.Printf("Receiving changeset: ADD %d, MODIFY %d, DELETE %d",
		len(changes.Add), len(changes.Modify), len(changes.Delete))

	for _, add := range changes.Add {
		tmpfile, err := ioutil.TempFile("", "changeset")
		if err != nil {
			return err
		}

		tf := time.Now()
		rc, err := c.receiveFileStream()
		if err != nil {
			_ = rc.Close()
			return err
		}

		n, err := io.Copy(tmpfile, rc)
		if err != nil {
			return err
		}

		log.Printf("Received %d bytes in %v", n, time.Now().Sub(tf))
		if err = rc.Close(); err != nil {
			return errors.Wrap(err, "error closing receive file stream")
		}

		if err := tmpfile.Close(); err != nil {
			return err
		}

		add.Dir = c.tree.Dir
		add.File = tmpfile

		if err := add.Apply(c.tree); err == ErrTreeAlreadyUpToDate {
			log.Printf("Ignoring ADD(%s): Tree already up to date", add.Path)
		} else if err != nil {
			return errors.Wrap(err, "error applying add changes to tree")
		}
	}

	for _, mod := range changes.Modify {
		var patch Patch

		if err := c.receive(&patch); err != nil {
			return err
		}

		log.Printf("Received %d blocks of change", len(patch.Blocks))
		mod.Patch = patch
		mod.To.Dir = c.tree.Dir
		if err := mod.Apply(c.tree); err == ErrTreeAlreadyUpToDate {
			log.Printf("Ignoring MODIFY(%s): Tree already up to date", mod.To.Path)
		} else if err != nil {
			return errors.Wrap(err, "error applying modify changes to tree")
		}
	}

	for _, del := range changes.Delete {
		del.To.Dir = c.tree.Dir
		if err := del.Apply(c.tree); err == ErrTreeAlreadyUpToDate {
			log.Printf("Ignoring DELETE(%s): Tree already up to date", del.To.Path)
		} else if err != nil {
			return errors.Wrap(err, "error applying delete changes to tree")
		}
	}

	log.Printf("Finished receiving changeset in %v", time.Now().Sub(t))
	return nil
}

// sendPushSync triggers a one-way sync, where a change set is pushed to the remote site
func (c *connection) sendPushSync(changes ChangeSet) error {
	log.Printf("Waiting for stream lock for push sync")
	c.streamMutex.Lock()
	defer c.streamMutex.Unlock()

	// Send a list of files that we have
	if err := c.send(&SyncMessage{Changes: changes}); err != nil {
		return err
	}

	if err := c.sendChangeSet(changes); err != nil {
		return err
	}

	return nil
}

// sendPullSync triggers a two-legged sync, where we send a filetree, then get back changes
// and apply them to our tree
func (c *connection) sendPullSync() error {
	log.Printf("Waiting for stream lock for pull sync")
	c.streamMutex.Lock()
	defer c.streamMutex.Unlock()
	log.Printf("Acquired stream lock")

	// Send a list of files that we have
	if err := c.send(&SyncMessage{Files: c.tree.FileSet}); err != nil {
		return err
	}

	var syncResp SyncResponseMessage

	// and get back a list of changes
	if err := c.receive(&syncResp); err != nil {
		return errors.Wrap(err, "Expected SyncResponseMessage")
	}

	if err := c.receiveChangeSet(syncResp.Changes); err != nil {
		return err
	}

	return nil
}

func (c *connection) receivePullSync(syncReq SyncMessage) error {
	var changes = Diff(
		FileSet(syncReq.Files),
		FileSet(c.tree.FileSet),
		c.tree.MatcherSet,
	)

	log.Printf("Remote side sent %d files", len(syncReq.Files))

	// generate patches
	for idx, mod := range changes.Modify {
		if len(mod.Patch.Blocks) == 0 && !fingerprintEqual(mod.From.Fingerprint, mod.To.Fingerprint) {
			log.Printf("Generating patch for %s", mod.To.Path)
			changes.Modify[idx].Patch = generatePatch(filepath.Join(mod.To.Dir, mod.To.Path), mod.From.Fingerprint, mod.To.Fingerprint)
		}
	}

	log.Printf("Waiting for stream lock for push sync")
	c.streamMutex.Lock()
	defer c.streamMutex.Unlock()
	log.Printf("Acquired stream lock")

	// generate a fileset of the diff with our filesystem and send it back
	if err := c.send(&SyncResponseMessage{Changes: changes}); err != nil {
		return err
	}

	if err := c.sendChangeSet(changes); err != nil {
		return err
	}

	return nil
}

func (c *connection) receivePushSync(syncReq SyncMessage) error {
	log.Printf("Waiting for stream lock for receiving pull sync")
	c.streamMutex.Lock()
	defer c.streamMutex.Unlock()
	log.Printf("Acquired stream lock")

	return c.receiveChangeSet(syncReq.Changes)
}

func (c *connection) sendFileStream(f *os.File) chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)
		stream, err := c.session.OpenStream()
		if err != nil {
			ch <- err
			return
		}
		if _, err = io.Copy(stream, f); err != nil {
			ch <- err
			return
		}
		if err = stream.Close(); err != nil {
			ch <- err
			return
		}
	}()
	return ch
}

func (c *connection) receiveFileStream() (io.ReadCloser, error) {
	stream, err := c.session.AcceptStream()
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *connection) send(req msgp.Encodable) error {
	log.Printf("Sending %T", req)
	if err := req.EncodeMsg(c.writer); err != nil {
		return err
	}

	if err := c.writer.Flush(); err != nil {
		return err
	}

	return nil
}

func (c *connection) receive(resp msgp.Decodable) error {
	log.Printf("Receiving %T", resp)
	return resp.DecodeMsg(c.reader)
}

type changeSetBuffer struct {
	*sync.Cond
	ChangeSet
}

func newChangeSetBuffer() *changeSetBuffer {
	m := sync.Mutex{}
	return &changeSetBuffer{Cond: sync.NewCond(&m)}
}

func (cs *changeSetBuffer) Add(c ChangeSet) {
	cs.L.Lock()
	cs.ChangeSet = cs.ChangeSet.Merge(c)
	cs.Broadcast()
	cs.L.Unlock()
}
