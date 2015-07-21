package main

import (
	"log"
	"net"
	"time"

	"github.com/docker/libchan"
	"github.com/docker/libchan/spdy"
)

type Client struct {
	Watcher *Watcher
	Tree    *FileTree
}

func NewClient(tree *FileTree, watcher *Watcher) *Client {
	return &Client{Tree: tree, Watcher: watcher}
}

func (c *Client) Connect(address string) {
	var client net.Conn
	var err error

	log.Printf("connecting to tcp://%s", address)
	client, err = net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	transport, err := spdy.NewClientTransport(client)
	if err != nil {
		log.Fatal(err)
	}

	sender, err := transport.NewSendChannel()
	if err != nil {
		log.Fatal(err)
	}

	remote, err := c.connectRemote(sender)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("running initial sync")

	t := time.Now()
	if err = remote.SyncFiles(c.Tree); err != nil {
		log.Fatal(err)
	}

	log.Printf("finished sync in %s", time.Now().Sub(t))

	if err = remote.SendEvents(c.Watcher.Events); err != nil {
		log.Fatal(err)
	}

	for ev := range remote.ReceiveEvents() {
		log.Printf("got remote event %s", ev.String())
	}
}

func (c *Client) connectRemote(sender libchan.Sender) (*Remote, error) {
	responseReceiver, responseSender := libchan.Pipe()
	eventReceiver, eventSender := libchan.Pipe()
	fileRequestReceiver, _ := libchan.Pipe()

	connect := &ClientConnect{
		ResponseChan:    responseSender,
		EventChan:       eventReceiver,
		FileRequestChan: fileRequestReceiver,
	}

	if err := sender.Send(connect); err != nil {
		return nil, err
	}

	response := &ClientConnectResponse{}
	if err := responseReceiver.Receive(response); err != nil {
		return nil, err
	}

	remote := &Remote{
		LocalEvents:  eventSender,
		RemoteEvents: response.EventChan,
		FileRequests: response.FileRequestChan,
	}

	return remote, nil
}
