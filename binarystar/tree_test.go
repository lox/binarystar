package binarystar_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/lox/binarystar/binarystar"
)

func TestTreeRelativePath(t *testing.T) {

}

func TestTreeWatch(t *testing.T) {
	dir := newTestFileDir(t,
		testFile{Path: "llamas.txt", Mode: 0600, Data: []byte("llamas rock")},
		testFile{Path: "subdir/alpacas.txt", Mode: 0644, Data: []byte("alpacas are ok")},
	)
	defer dir.RemoveAll()

	// scan file system
	tree, err := binarystar.NewTree(dir.Dir, nil)
	if err != nil {
		t.Fatal(err)
	}

	// some watchers need a short pause otherwise they catch the above events
	time.Sleep(time.Millisecond * 1000)

	// cancel our watcher if it takes too long
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := tree.Watch(ctx); err != nil {
		t.Fatal(err)
	}

	dir.WriteFile(testFile{Path: "llamas.txt", Mode: 0700, Data: []byte("some content1")})
	dir.WriteFile(testFile{Path: "another_file.txt", Mode: 0644, Data: []byte("some content2")})

	var once sync.Once

	var changes binarystar.ChangeSet
	for cs := range tree.Changes {
		changes = changes.Merge(cs)
		once.Do(cancel)
	}

	if changes.Len() != 2 {
		t.Fatalf("Expected 2 changes, got:\n%#v", changes)
	}

	if len(changes.Add) != 1 {
		t.Fatalf("Expected 1 ADD changes, got:\n%#v", len(changes.Add))
	}

	if len(changes.Modify) != 1 {
		t.Fatalf("Expected 1 ADD changes, got:\n%#v", len(changes.Modify))
	}

}
