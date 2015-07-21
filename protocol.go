package main

import "github.com/docker/libchan"

// ClientConnect is sent when a client connects to a server
type ClientConnect struct {
	ResponseChan    libchan.Sender
	EventChan       libchan.Receiver
	FileRequestChan libchan.Receiver
}

// ClientConnectResponse is sent from a server to a client in response to a ClientConnect
type ClientConnectResponse struct {
	EventChan       libchan.Receiver
	FileRequestChan libchan.Receiver
}
