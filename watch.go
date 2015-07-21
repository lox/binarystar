package main

import (
	"log"
	"time"

	"github.com/rjeczalik/notify"
)

type Watcher struct {
	Dir     string
	Matcher *MatcherSet
	Events  chan FileEvent
}

func NewWatcher(dir string) *Watcher {
	return &Watcher{Dir: dir, Matcher: NewMatcherSet(), Events: make(chan FileEvent)}
}

func (w *Watcher) Watch() error {
	notifyCh := make(chan notify.EventInfo, 10)

	if err := notify.Watch(w.Dir+"/...", notifyCh, notifyEvents...); err != nil {
		return err
	}

	go func() {
		for ei := range notifyCh {
			if !w.Matcher.Match(ei.Path()) {
				continue
			}

			ev := FileEvent{
				Path: ei.Path(),
				Type: convertNotifyType(ei.Event()),
				Time: time.Now().UnixNano(),
			}

			log.Println(ev)
			w.Events <- ev
		}
	}()

	return nil
}
