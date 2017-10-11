// +build darwin

package binarystar

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsevents"
)

func watch(ctx context.Context, tree *Tree) error {
	dev, err := fsevents.DeviceForPath(tree.Dir)
	if err != nil {
		return err
	}

	events := make(chan []fsevents.Event)
	es := &fsevents.EventStream{
		Paths:   []string{tree.Dir},
		Latency: 100 * time.Millisecond,
		Device:  dev,
		Resume:  false,
		Flags:   fsevents.FileEvents | fsevents.WatchRoot,
		Events:  events,
	}

	es.Start()
	log.Printf("Starting watching %s", tree.Dir)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Done, stopping: %#v", ctx.Err())
				es.Stop()
				return
			case msg := <-es.Events:
				for _, event := range msg {
					switch {
					// Not a file event
					case event.Flags&fsevents.ItemIsFile != fsevents.ItemIsFile:
						continue
					// Add events
					case event.Flags&fsevents.ItemCreated == fsevents.ItemCreated:
						if err := handleItemCreated(tree, event); err != nil {
							log.Printf("Error handling create: %#v", err)
						}
					case event.Flags&fsevents.ItemModified == fsevents.ItemModified:
						if err := handleItemModified(tree, event); err != nil {
							log.Printf("Error handling modify: %#v", err)
						}
					case event.Flags&fsevents.ItemRemoved == fsevents.ItemRemoved:
						if err := handleItemDeleted(tree, event); err != nil {
							log.Printf("Error handling delete: %#v", err)
						}
					case event.Flags&fsevents.ItemRenamed == fsevents.ItemRenamed:
						// TODO
					case event.Flags&fsevents.ItemChangeOwner == fsevents.ItemChangeOwner:
						// TODO
					default:
						logEvent(event)
					}
				}
			}
		}
	}()

	return nil
}

func handleItemCreated(tree *Tree, event fsevents.Event) error {
	path, err := filepath.EvalSymlinks("/" + event.Path)
	if err != nil {
		return err
	}

	rel, err := tree.RelativePath(path)
	if err != nil {
		return err
	}

	fi, err := ScanFile(tree.Dir, rel)
	if err != nil {
		return err
	}

	ac := AddChange{FileInfo: fi}

	if err := canApplyAddChange(ac, tree); err == nil {
		return applyAddChange(ac, tree)
	}

	return nil
}

func handleItemDeleted(tree *Tree, event fsevents.Event) error {
	rel, err := tree.RelativePath("/" + event.Path)
	if err != nil {
		return err
	}
	fi, ok := tree.Get(rel)
	if !ok {
		return fmt.Errorf("Can't find %s to delete", rel)
	}

	dc := DeleteChange{
		From: fi,
		To:   fi.AsDeleted(),
	}

	if err := canApplyDeleteChange(dc, tree); err == nil {
		return applyDeleteChange(dc, tree)
	}

	return nil
}

func handleItemModified(tree *Tree, event fsevents.Event) error {
	path, err := filepath.EvalSymlinks("/" + event.Path)
	if err != nil {
		return err
	}

	rel, err := tree.RelativePath(path)
	if err != nil {
		return err
	}

	current, ok := tree.Get(rel)
	if !ok {
		return ErrNotExists
	}

	fi, err := ScanFile(tree.Dir, rel)
	if err != nil {
		return err
	}

	mc := ModifyChange{
		From:  current,
		To:    fi,
		Patch: generatePatch(path, current.Fingerprint, fi.Fingerprint),
	}

	if err := canApplyModifyChange(mc, tree); err == nil {
		return applyModifyChange(mc, tree)
	}

	return nil
}

var noteDescription = map[fsevents.EventFlags]string{
	fsevents.MustScanSubDirs:   "MustScanSubdirs",
	fsevents.UserDropped:       "UserDropped",
	fsevents.KernelDropped:     "KernelDropped",
	fsevents.EventIDsWrapped:   "EventIDsWrapped",
	fsevents.HistoryDone:       "HistoryDone",
	fsevents.RootChanged:       "RootChanged",
	fsevents.Mount:             "Mount",
	fsevents.Unmount:           "Unmount",
	fsevents.ItemCreated:       "Created",
	fsevents.ItemRemoved:       "Removed",
	fsevents.ItemInodeMetaMod:  "InodeMetaMod",
	fsevents.ItemRenamed:       "Renamed",
	fsevents.ItemModified:      "Modified",
	fsevents.ItemFinderInfoMod: "FinderInfoMod",
	fsevents.ItemChangeOwner:   "ChangeOwner",
	fsevents.ItemXattrMod:      "XAttrMod",
	fsevents.ItemIsFile:        "IsFile",
	fsevents.ItemIsDir:         "IsDir",
	fsevents.ItemIsSymlink:     "IsSymLink",
}

func logEvent(event fsevents.Event) {
	note := ""
	for bit, description := range noteDescription {
		if event.Flags&bit == bit {
			note += description + " "
		}
	}
	log.Printf("FSEvent Path: %s Flags: %s", event.Path, note)
}
