package filetree_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/lox/binarystar/filetree"
)



func TestDiffSubset(t *testing.T) {
	tree1 := createTestTree(t, []testFile{
		{File: filetree.File{Path: "llamas.txt", Mode: 0400}, Data: []byte("llamas rock")},
		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
	})
	defer os.RemoveAll(tree1.Dir())

	writeFile(t, tree1, "another_file1.txt", []byte("some content1"), 0700)
	writeFile(t, tree1, "subdir/blah.txt", []byte("some content1"), 0700)

	m := filetree.NewMatcherSet(filetree.MatchPrefix("subdir/"))
	subdirTree, err := filetree.ScanWithMatcher(tree1.Dir(), m)
	if err != nil {
		t.Fatal(err)
	}

	if subdirTree.Len() != 2 {
		t.Fatalf("Expected subdir/ tree to have 2 items, got %d", subdirTree.Len())
	}

	changes := tree1.DiffSubset(subdirTree, m)

	if len(changes) != 1 {
		t.Fatalf("Expected only 1 change under subdir: %v", changes)
	}
}

func TestTreeDiffShowsFilesAdded(t *testing.T) {
	tree1 := createTestTree(t, []testFile{
		{File: filetree.File{Path: "llamas.txt", Mode: 0400}, Data: []byte("llamas rock")},
		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
	})
	defer os.RemoveAll(tree1.Dir())

	writeFile(t, tree1, "another_file1.txt", []byte("some content1"), 0700)
	writeFile(t, tree1, "another_file2.txt", []byte("some content2"), 0600)
	writeFile(t, tree1, "another_file3.txt", []byte("some content3"), 0644)

	tree2, err := filetree.Scan(tree1.Dir())
	if err != nil {
		t.Fatal(err)
	}

	actualChanges := tree1.Diff(tree2)
	expectedChanges := filetree.Changes{
		filetree.CreateFileChange{Path: "another_file1.txt", Size: 13, Mode: 0700, ModTime: testMtime},
		filetree.CreateFileChange{Path: "another_file2.txt", Size: 13, Mode: 0600, ModTime: testMtime},
		filetree.CreateFileChange{Path: "another_file3.txt", Size: 13, Mode: 0644, ModTime: testMtime},
	}

	if err = changesEqual(expectedChanges, actualChanges); err != nil {
		t.Fatal(err)
	}
}

func TestTreeDiffShowsFilesDeleted(t *testing.T) {
	tree1 := createTestTree(t, []testFile{
		{File: filetree.File{Path: "llamas.txt", Mode: 0400}, Data: []byte("llamas rock")},
		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
	})
	defer os.RemoveAll(tree1.Dir())

	if err := os.RemoveAll(filepath.Join(tree1.Dir(), "subdir")); err != nil {
		t.Fatal(err)
	}

	tree2, err := filetree.Scan(tree1.Dir())
	if err != nil {
		t.Fatal(err)
	}

	actualChanges := tree1.Diff(tree2)
	expectedChanges := filetree.Changes{
		filetree.DeleteFileChange{Path: "subdir/alpacas.txt"},
	}

	if err = changesEqual(expectedChanges, actualChanges); err != nil {
		t.Fatal(err)
	}
}

func TestTreeDiffShowsFilesModified(t *testing.T) {
	tree1 := createTestTree(t, []testFile{
		{File: filetree.File{Path: "llamas.txt", Mode: 0700}, Data: []byte("llamas rock")},
		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
	})
	defer os.RemoveAll(tree1.Dir())

	writeFile(t, tree1, "llamas.txt", []byte("some content1"), 0700)

	tree2, err := filetree.Scan(tree1.Dir())
	if err != nil {
		t.Fatal(err)
	}

	changes := tree1.Diff(tree2)

	if len(changes) != 1 {
		t.Fatalf("Expected 1 change, got %d", len(changes))
	}

	ch, ok := changes[0].(filetree.FileContentChange)
	if !ok {
		t.Fatalf("Change wasn't a FileContentChange: %T", changes[0])
	}

	if ch.Path != "llamas.txt" {
		t.Fatalf("Expected llamas.txt to change, got %s", ch.Path)
	}
}

func TestApplyingChangesBetweenTrees(t *testing.T) {
	tree1 := createTestTree(t, []testFile{
		{File: filetree.File{Path: "llamas.txt", Mode: 0700}, Data: []byte("llamas rock")},
		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
	})
	defer os.RemoveAll(tree1.Dir())

	tree2 := createTestTree(t, []testFile{
		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas rock")},
	})
	defer os.RemoveAll(tree1.Dir())

	changes := tree1.Diff(tree2)
	if err := tree1.Apply(changes); err != nil {
		t.Fatal(err)
	}

	after := tree1.Diff(tree2)
	if len(after) > 0 {
		t.Fatalf("Expected no changes after, got:\n%s", after.String())
	}
}

func changesEqual(expected, actual filetree.Changes) error {
	c1s := expected.String()
	c2s := actual.String()

	if !reflect.DeepEqual(c1s, c2s) {
		return fmt.Errorf("Unexpected changes:\nExpected:\n%s\nActual:\n%s", c1s, c2s)
	}

	return nil
}

var (
	testAtime = time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	testMtime = time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
)

func writeFile(t *testing.T, tree *filetree.Tree, key string, data []byte, mode os.FileMode) {
	path := filepath.Join(tree.Dir(), key)
	if err := ioutil.WriteFile(path, data, mode); err != nil {
		t.Fatal(err)
	}
	if err := os.Chtimes(path, testAtime, testMtime); err != nil {
		t.Fatal(err)
	}
}

func createTestTree(t *testing.T, files []testFile) *filetree.Tree {
	dir := createTestDir(t, files)
	tree, err := filetree.Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	return tree
}

func createTestDir(t *testing.T, files []testFile) string {
	dir, err := ioutil.TempDir("", "treetest")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		path := filepath.Join(dir, f.Path)
		if f.Mode == os.FileMode(0) {
			f.Mode = os.FileMode(0700)
		}
		if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		if err = ioutil.WriteFile(path, f.Data, f.Mode); err != nil {
			t.Fatal(err)
		}
		if err := os.Chtimes(path, testAtime, testMtime); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

type testFile struct {
	filetree.File
	Data []byte
}
