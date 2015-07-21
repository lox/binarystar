package main

import (
	"log"
	"net"

	"github.com/docker/libchan"
	"github.com/docker/libchan/spdy"
)

type Server struct {
	Watcher *Watcher
	Tree    *FileTree
}

func NewServer(tree *FileTree, watcher *Watcher) *Server {
	return &Server{Tree: tree, Watcher: watcher}
}

func (s *Server) Listen(address string) {
	var listener net.Listener
	var err error

	listener, err = net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("listening on %s", address)
	tl, err := spdy.NewTransportListener(listener, spdy.NoAuthenticator)
	if err != nil {
		log.Fatal(err)
	}

	for {
		t, err := tl.AcceptTransport()
		if err != nil {
			log.Print(err)
			break
		}

		go func() {
			for {
				receiver, err := t.WaitReceiveChannel()
				if err != nil {
					log.Print(err)
					break
				}

				go func() {
					for {
						remote, err := s.acceptRemote(receiver)
						if err != nil {
							log.Println(err)
							break
						}

						log.Printf("client connected")

						go func() {
							if err = remote.SendEvents(s.Watcher.Events); err != nil {
								panic(err)
							}
						}()

						go func() {
							for ev := range remote.ReceiveEvents() {
								log.Printf("got remote event %s", ev.String())
							}
						}()
					}
				}()
			}
		}()
	}
}

func (s *Server) acceptRemote(receiver libchan.Receiver) (*Remote, error) {
	connect := &ClientConnect{}
	err := receiver.Receive(connect)
	if err != nil {
		return nil, err
	}

	eventReceiver, eventSender := libchan.Pipe()
	fileRequestReceiver, _ := libchan.Pipe()

	response := &ClientConnectResponse{
		EventChan:       eventReceiver,
		FileRequestChan: fileRequestReceiver,
	}

	err = connect.ResponseChan.Send(response)
	if err != nil {
		return nil, err
	}

	remote := &Remote{
		LocalEvents:  eventSender,
		RemoteEvents: response.EventChan,
		FileRequests: response.FileRequestChan,
	}

	return remote, nil
}
