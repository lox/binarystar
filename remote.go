package main

import "github.com/docker/libchan"

type Remote struct {
	LocalEvents  libchan.Sender
	RemoteEvents libchan.Receiver
	FileRequests libchan.Receiver
}

func (r *Remote) SyncFiles(tree *FileTree) error {
	for path, f := range tree.Files {
		ev := &FileEvent{
			Type: SyncEvent,
			Path: path,
			Time: f.ModTime.Unix(),
			Mode: f.Mode.String(),
		}
		if err := r.LocalEvents.Send(ev); err != nil {
			return err
		}
	}
	return nil
}

func (r *Remote) SendEvents(events chan FileEvent) error {
	for ev := range events {
		if err := r.LocalEvents.Send(ev); err != nil {
			return err
		}
	}
	return nil
}

func (r *Remote) ReceiveEvents() (events chan FileEvent) {
	go func() {
		for {
			event := &FileEvent{}
			err := r.RemoteEvents.Receive(event)
			if err != nil {
				panic(err)
			}
			events <- *event
		}
		close(events)
	}()
	return
}
