// package filetree_test

// import (
// 	"context"
// 	"os"
// 	"sync"
// 	"testing"
// 	"time"

// 	"github.com/lox/binarystar/filetree"
// )

// func TestWatch(t *testing.T) {
// 	tree1 := createTestTree(t, []testFile{
// 		{File: filetree.File{Path: "llamas.txt", Mode: 0700}, Data: []byte("llamas rock")},
// 		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
// 	})
// 	defer os.RemoveAll(tree1.Dir())

// 	// some watchers need a short pause otherwise they catch the above events
// 	time.Sleep(time.Millisecond * 50)

// 	// cancel our watcher if it takes too long
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	changes, err := tree1.Watch(ctx, filetree.NewMatcherSet(filetree.MatchAll))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	writeFile(t, tree1, "llamas.txt", []byte("some content1"), 0700)
// 	writeFile(t, tree1, "another_file.txt", []byte("some content1"), 0644)

// 	var once sync.Once

// 	allChanges := []filetree.Change{}
// 	for change := range changes {
// 		allChanges = append(allChanges, change)
// 		once.Do(cancel)
// 	}

// 	if len(allChanges) != 2 {
// 		t.Fatalf("Expected 2 changes, got:\n%s", filetree.Changes(allChanges))
// 	}
// }
