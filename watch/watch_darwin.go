// // +build darwin

// package filetree

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/fsnotify/fsevents"
// )

// func init() {
// 	newWatcher = func(tree *Tree, m *MatcherSet) (watcher, error) {
// 		return newDarwinWatcher(tree, m)
// 	}
// }

// var modEvts = []fsevents.EventFlags{
// 	fsevents.ItemCreated,
// 	fsevents.ItemRemoved,
// 	fsevents.ItemRenamed,
// 	fsevents.ItemModified,
// }

// var noteDescription = map[fsevents.EventFlags]string{
// 	fsevents.MustScanSubDirs:   "MustScanSubdirs",
// 	fsevents.UserDropped:       "UserDropped",
// 	fsevents.KernelDropped:     "KernelDropped",
// 	fsevents.EventIDsWrapped:   "EventIDsWrapped",
// 	fsevents.HistoryDone:       "HistoryDone",
// 	fsevents.RootChanged:       "RootChanged",
// 	fsevents.Mount:             "Mount",
// 	fsevents.Unmount:           "Unmount",
// 	fsevents.ItemCreated:       "Created",
// 	fsevents.ItemRemoved:       "Removed",
// 	fsevents.ItemInodeMetaMod:  "InodeMetaMod",
// 	fsevents.ItemRenamed:       "Renamed",
// 	fsevents.ItemModified:      "Modified",
// 	fsevents.ItemFinderInfoMod: "FinderInfoMod",
// 	fsevents.ItemChangeOwner:   "ChangeOwner",
// 	fsevents.ItemXattrMod:      "XAttrMod",
// 	fsevents.ItemIsFile:        "IsFile",
// 	fsevents.ItemIsDir:         "IsDir",
// 	fsevents.ItemIsSymlink:     "IsSymLink",
// }

// func logEvent(event fsevents.Event) {
// 	note := ""
// 	for bit, description := range noteDescription {
// 		if event.Flags&bit == bit {
// 			note += description + " "
// 		}
// 	}
// 	log.Printf("EventID: %d Path: %s Flags: %s", event.ID, event.Path, note)
// }

// type darwinWatcher struct {
// 	tree  *Tree
// 	match *MatcherSet
// 	es    *fsevents.EventStream
// }

// func newDarwinWatcher(tree *Tree, match *MatcherSet) (*darwinWatcher, error) {
// 	dev, err := fsevents.DeviceForPath(tree.Dir())
// 	if err != nil {
// 		return nil, err
// 	}

// 	events := make(chan []fsevents.Event)
// 	es := &fsevents.EventStream{
// 		Paths:   []string{tree.Dir()},
// 		Latency: 100 * time.Millisecond,
// 		Device:  dev,
// 		Resume:  false,
// 		Flags:   fsevents.FileEvents | fsevents.WatchRoot,
// 		Events:  events,
// 	}

// 	return &darwinWatcher{
// 		tree:  tree,
// 		match: match,
// 		es:    es,
// 	}, nil
// }

// func (w *darwinWatcher) Watch(ctx context.Context, changes chan Change) {
// 	defer close(changes)
// 	w.es.Start()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Printf("Done, stopping")
// 			w.es.Stop()
// 			return
// 		case msg := <-w.es.Events:
// 			for _, event := range msg {
// 				logEvent(event)
// 			}
// 			return
// 		}
// 	}
// }

// // checkFlag lets us know if this event is important
// func checkFlag(e fsevents.EventFlags) bool {
// 	if e&fsevents.ItemIsFile == 0 {
// 		return false
// 	}

// 	for i := range modEvts {
// 		if e&modEvts[i] == modEvts[i] {
// 			return true
// 		}
// 	}
// 	return false
// }
