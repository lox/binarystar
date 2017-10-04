package binarystar_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/lox/binarystar/binarystar"
)

func TestAddingToSet(t *testing.T) {
	set := &binarystar.FileSet{}
	file := binarystar.FileInfo{Path: "blah", Mode: 0700}

	err := set.Add(file)
	if err != nil {
		t.Fatal(err)
	}

	ff, ok := set.Get("blah")
	if !ok {
		t.Fatalf("Didn't find blah")
	}

	if !reflect.DeepEqual(file, ff) {
		t.Fatalf("Not equal %v and %v", file, ff)
	}
}

func TestModifyingSet(t *testing.T) {
	set := binarystar.FileSet{}

	_ = set.Add(binarystar.FileInfo{Path: "blah1", Size: 100})

	err := set.Modify("blah1", func(f *binarystar.FileInfo) error {
		f.Size = 600
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	f, ok := set.Get("blah1")
	if !ok {
		t.Fatalf("Didn't find blah1")
	}

	if f.Size != 600 {
		t.Fatalf("Expected size to have been updated, got %v", f.Size)
	}
}

func TestDeletingFromSet(t *testing.T) {
	set := binarystar.FileSet{}

	_ = set.Add(binarystar.FileInfo{Path: "blah", Mode: 0700})

	_, ok := set.Get("blah")
	if !ok {
		t.Fatalf("Didn't find blah")
	}

	err := set.Delete("not-exists")
	if err != os.ErrNotExist {
		t.Fatalf("Missing key should have failed")
	}

	err = set.Delete("blah")
	if err != nil {
		t.Fatal(err)
	}

	_, ok = set.Get("blah")
	if ok {
		t.Fatalf("Blah still there")
	}
}
