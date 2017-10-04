package binarystar

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"github.com/tinylib/msgp/msgp"
)

type Daemon struct {
	dir        string
	files      FileSet
	matcherSet *MatcherSet
}

func NewDaemon(dir string, m *MatcherSet) (*Daemon, error) {
	files, err := Scan(dir, m)
	if err != nil {
		return nil, err
	}
	log.Printf("Found %d files", len(files))
	return &Daemon{dir: dir, files: files, matcherSet: m}, nil
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

		if err = d.startStream(session, stream); err != nil {
			log.Printf("startSteam failed: %v", err)
		}
	}
}

func (d *Daemon) startStream(session *yamux.Session, stream net.Conn) error {
	connection := &connection{
		files:      d.files,
		matcherSet: d.matcherSet,
		dir:        d.dir,
		session:    session,
		stream:     stream,
		reader:     msgp.NewReader(stream),
		writer:     msgp.NewWriter(stream),
	}

	log.Printf("Accepted a stream, waiting for sync requests")
	if err := connection.receiveSync(); err != nil {
		return err
	}

	return nil
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

	connection := &connection{
		files:      d.files,
		dir:        d.dir,
		matcherSet: d.matcherSet,
		session:    session,
		stream:     stream,
		reader:     msgp.NewReader(stream),
		writer:     msgp.NewWriter(stream),
	}

	// Fire off an initial synchronization
	if err := connection.sendSync(); err != nil {
		return errors.Wrap(err, "Initial sync failed")
	}

	return connection.receiveSync()
}

type connection struct {
	dir        string
	files      FileSet
	matcherSet *MatcherSet
	stream     net.Conn
	session    *yamux.Session
	reader     *msgp.Reader
	writer     *msgp.Writer
}

func (c *connection) sendSync() error {
	// Send a list of files that we have
	if err := c.send(&SyncMessage{Files: c.files}); err != nil {
		return err
	}

	var syncResp SyncResponseMessage

	// and get back a list of changes
	if err := c.receive(&syncResp); err != nil {
		return errors.Wrap(err, "Expected SyncResponseMessage")
	}

	log.Printf("Changes: Adds %d, Modifies %d, Deletes %d",
		len(syncResp.Changes.Add),
		len(syncResp.Changes.Modify),
		len(syncResp.Changes.Delete),
	)

	// Process Deletes locally
	// for _, delete := range syncResp.Changes.Delete {
	// 	log.Printf("Deleting %s", delete.Path)
	// 	if err := os.Remove(filepath.Join(c.dir, delete.Path)); err != nil {
	// 		log.Printf("Error deleting file: %#v", err)
	// 		return err
	// 	}
	// 	err := c.files.Modify(delete.Path, func(f *FileInfo) error {
	// 		f.IsDeleted = true
	// 		f.ModTime = time.Now()
	// 		return nil
	// 	})
	// 	if err != nil {
	// 		log.Printf("Error modifying file: %#v", err)
	// 		return err
	// 	}
	// }

	var fileReqs = FileStreamRequestsMessage{}

	// Add changes we need verbatim
	for _, add := range syncResp.Changes.Add {
		fileReqs.Paths = append(fileReqs.Paths, add.Path)
	}

	// Send a list of files that we need streamed
	if err := c.send(&fileReqs); err != nil {
		return err
	}

	// Then process the files
	for _ = range fileReqs.Paths {
		var header FileStreamHeaderMessage
		if err := c.receive(&header); err != nil {
			return errors.Wrap(err, "Expected FileStreamHeaderMessage")
		}

		if header.Error != "" {
			log.Printf("Error from %s: %s", header.Path, header.Error)
			continue
		}

		// TODO: security checks on this
		localPath := filepath.Join(c.dir, header.Path)

		// TODO: better handling of directories, this won't respect permissions on other side
		if err := os.MkdirAll(filepath.Dir(localPath), 0700); err != nil {
			return err
		}

		targetFile, err := os.OpenFile(localPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		log.Printf("Waiting for file stream")
		rc, err := c.receiveFileStream()
		if err != nil {
			_ = rc.Close()
			return err
		}

		log.Printf("Copying to local file")
		n, err := io.Copy(targetFile, rc)
		if err != nil {
			return err
		}

		log.Printf("Wrote %d bytes to %s", n, localPath)
		log.Printf("Closing file stream")
		if err = rc.Close(); err != nil {
			log.Printf("Error closing: %#v", err)
			return err
		}

		// log.Printf("Setting times")
		// if err = os.Chtimes(localPath, header.ModTime, header.ModTime); err != nil {
		// 	return errors.Wrapf(err, "Error setting times on %s", localPath)
		// }

		// if err = os.Chmod(localPath, os.FileMode(header.Mode)); err != nil {
		// 	return errors.Wrapf(err, "Error setting mode on %s", localPath)
		// }

		updated, err := ScanFile(c.dir, header.Path)
		if err != nil {
			return errors.Wrapf(err, "Error scanning file %s", filepath.Join(c.dir, header.Path))
		}

		if !FingerprintEqual(updated.Fingerprint, header.Fingerprint) {
			return fmt.Errorf(
				"File might be corrupt, fingerprints didn't match. Expected %#v, got %#v",
				header.Fingerprint,
				updated.Fingerprint,
			)
		}

		if err = c.files.Add(updated); err != nil {
			return errors.Wrapf(err, "Add of %s failed", updated.Path)
		}
	}

	log.Printf("Sync is finished")
	return nil
}

func (c *connection) receiveSync() error {
	for {
		var syncReq SyncMessage

		// receive the initial sync message with remote files
		if err := c.receive(&syncReq); err != nil {
			return errors.Wrap(err, "Expected SyncMessage")
		}

		log.Printf("Remote files: %d Local Files: %d", len(syncReq.Files), len(c.files))

		var changes = Diff(FileSet(syncReq.Files), FileSet(c.files), c.matcherSet)

		// generate a fileset of the diff with our filesystem and send it back
		if err := c.send(&SyncResponseMessage{Changes: changes}); err != nil {
			return err
		}

		var fileReqs FileStreamRequestsMessage

		// receive a list of files to stream
		if err := c.receive(&fileReqs); err != nil {
			return errors.Wrap(err, "Expected FileStreamRequestsMessage")
		}

		var sendFileStreamError = func(err error) error {
			log.Printf("Sending file failed: %#v", err)
			if sendErr := c.send(&FileStreamHeaderMessage{Error: err.Error()}); sendErr != nil {
				return sendErr
			}
			return nil
		}

		// send back files
		for _, path := range fileReqs.Paths {
			fileInfo, ok := c.files.Get(path)
			if !ok {
				if sendErr := sendFileStreamError(errors.New("Unknown file")); sendErr != nil {
					return sendErr
				}
				continue
			}
			file, err := os.Open(filepath.Join(fileInfo.Dir, fileInfo.Path))
			if err != nil {
				if sendErr := sendFileStreamError(err); sendErr != nil {
					return sendErr
				}
				continue
			}
			defer file.Close()

			// send a header first with the filesize
			if sendErr := c.send(&FileStreamHeaderMessage{FileInfo: fileInfo}); sendErr != nil {
				return sendErr
			}

			// Now stream those bytes directly
			errCh := c.sendFileStream(file)

			log.Printf("Waiting for stream to finish")
			if err = <-errCh; err != nil {
				return err
			}

			log.Printf("Finished streaming bytes")

		}
		log.Printf("Sync is finished")
	}
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
		log.Printf("Opened stream %d for file", stream.StreamID())
		if _, err = io.Copy(stream, f); err != nil {
			ch <- err
			return
		}
		log.Printf("Finished copying into stream %d", stream.StreamID())
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

	log.Printf("Accepted stream %d for file", stream.StreamID())
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
