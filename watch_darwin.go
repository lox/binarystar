package main

import "github.com/rjeczalik/notify"

var notifyEvents = []notify.Event{
	notify.All,
	notify.FSEventsInodeMetaMod,
	notify.FSEventsRenamed,
	notify.FSEventsModified,
	notify.FSEventsChangeOwner,
	notify.FSEventsXattrMod,
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
	case notify.FSEventsInodeMetaMod:
		fallthrough
	case notify.FSEventsChangeOwner:
		return ModifyEvent
	default:
		return -1
	}
}
