package main

import "github.com/rjeczalik/notify"

var notifyEvents = []notify.Event{
	notify.All,
}

func convertNotifyType(ei notify.Event) int {
	switch ei {
	case notify.Create:
		return CreateEvent
	case notify.Remove:
		return RemoveEvent
	case notify.Write:
		return WriteEvent
	case notify.Rename:
		return RenameEvent
	default:
		return ModifyEvent
	}
}
