package main

import (
	"fmt"

	"github.com/docker/libchan"
)

const (
	SyncEvent   = 1
	CreateEvent = 2
	RemoveEvent = 3
	WriteEvent  = 4
	RenameEvent = 5
	ModifyEvent = 6
)

var eventNames = map[int]string{
	SyncEvent:   "Sync",
	CreateEvent: "Create",
	RemoveEvent: "Remove",
	WriteEvent:  "Write",
	RenameEvent: "Rename",
	ModifyEvent: "Modify",
}

type FileEvent struct {
	Type int
	Path string
	Time int64
	Mode string
}

func (e *FileEvent) String() string {
	return fmt.Sprintf("%s %q", eventNames[e.Type], e.Path)
}

func ReceiveEvents(ch chan FileEvent, receiver libchan.Receiver) error {
	for {
		event := &FileEvent{}
		err := receiver.Receive(event)
		if err != nil {
			return err
		} else {
			ch <- *event
		}
	}
	return nil
}

func SendEvents(ch chan FileEvent, sender libchan.Sender) error {
	for ev := range ch {
		if err := sender.Send(ev); err != nil {
			return err
		}
	}
	return nil
}
