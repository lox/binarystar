package main

import "github.com/docker/libchan"

type Node struct {
	EventSender   libchan.Sender
	EventReceiver libchan.Receiver
}
