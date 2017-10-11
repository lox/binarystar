package binarystar_test

import (
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lox/binarystar/binarystar"
)

func TestFileDiffWithDifferentPaths(t *testing.T) {
	if binarystar.FileInfoEqual(binarystar.FileInfo{Path: "llamas"},
		binarystar.FileInfo{Path: "alpacas"}) {
		t.Fatalf("Different paths shouldn't be equal")
	}
}

func TestFileDiffWithDifferentModes(t *testing.T) {
	if binarystar.FileInfoEqual(binarystar.FileInfo{Mode: 0700},
		binarystar.FileInfo{Mode: 0660}) {
		t.Fatalf("Different modes shouldn't be equal")
	}
}

func TestFileDiffWithDifferentTimes(t *testing.T) {
	if binarystar.FileInfoEqual(binarystar.FileInfo{ModTime: testMtime},
		binarystar.FileInfo{ModTime: testMtime.Add(time.Hour)}) {
		t.Fatalf("Different times shouldn't be equal")
	}
}

func TestFileDiffWithDifferentSizes(t *testing.T) {
	if binarystar.FileInfoEqual(binarystar.FileInfo{Size: 1000},
		binarystar.FileInfo{Size: 1001}) {
		t.Fatalf("Different sizes shouldn't be equal")
	}
}

func TestFileDiffWithDifferentFingerprints(t *testing.T) {
	if binarystar.FileInfoEqual(binarystar.FileInfo{Fingerprint: binarystar.FingerprintFromString("llamas")},
		binarystar.FileInfo{Fingerprint: binarystar.FingerprintFromString("alpacas")}) {
		t.Fatalf("Different fingerprints shouldn't be equal")
	}
}

func TestDiff(t *testing.T) {
	var testCases = map[string]struct {
		s1, s2 binarystar.FileSet
		m      *binarystar.MatcherSet
		// expectations
		add int
		mod int
		del int
	}{
		"Added Files": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
				{Path: "subdir/alpacas.txt", Mode: 0644, ModTime: testMtime},
			},
			add: 1,
		},
		"Removed Files": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
				{Path: "subdir/alpacas.txt", Mode: 0644, ModTime: testMtime},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
			},
			del: 1,
		},
		"Order of Files Ignored": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
				{Path: "subdir/alpacas.txt", Mode: 0644, ModTime: testMtime},
			},
			s2: binarystar.FileSet{
				{Path: "subdir/alpacas.txt", Mode: 0644, ModTime: testMtime},
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
			},
		},
		"Modified Mod Time": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime.Add(time.Hour)},
			},
			mod: 1,
		},
		"Modified Mode": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0660, ModTime: testMtime},
			},
			mod: 1,
		},
		"Modified Size": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0600, ModTime: testMtime, Size: 1000},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Mode: 0660, ModTime: testMtime, Size: 1100},
			},
			mod: 1,
		},
		"Modified Fingerprints": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Fingerprint: binarystar.FingerprintFromString("llamas")},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Fingerprint: binarystar.FingerprintFromString("alpacas")},
			},
			mod: 1,
		},
		"Same Fingerprints": {
			s1: binarystar.FileSet{
				{Path: "llamas.txt", Fingerprint: binarystar.FingerprintFromString("llamas")},
			},
			s2: binarystar.FileSet{
				{Path: "llamas.txt", Fingerprint: binarystar.FingerprintFromString("llamas")},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if tc.m == nil {
				tc.m = binarystar.NewMatcherSet(binarystar.MatchAll)
			}
			changes := binarystar.Diff(tc.s1, tc.s2, tc.m)
			if actual := len(changes.Add); actual != tc.add {
				t.Errorf("Expected %d ADD changes, got %d", tc.add, actual)
				t.Log(spew.Sdump(changes))
			}
			if actual := len(changes.Modify); actual != tc.mod {
				t.Errorf("Expected %d MODIFY changes, got %d", tc.mod, actual)
				t.Log(spew.Sdump(changes))
			}
			if actual := len(changes.Delete); actual != tc.del {
				t.Errorf("Expected %d DELETE changes, got %d", tc.del, actual)
				t.Log(spew.Sdump(changes))
			}

			for _, del := range changes.Delete {
				if del.IsDeleted == false {
					t.Error("Expected delete changes to be all IsDeleted=true")
					t.Log(spew.Sdump(del))
				}
			}
		})
	}
}

// m := filetree.NewMatcherSet()
// subdirTree, err := filetree.ScanWithMatcher(tree1.Dir(), m)
// if err != nil {
// 	t.Fatal(err)
// }

// if subdirTree.Len() != 2 {
// 	t.Fatalf("Expected subdir/ tree to have 2 items, got %d", subdirTree.Len())
// }

// changes := tree1.DiffSubset(subdirTree, m)

// if len(changes) != 1 {
// 	t.Fatalf("Expected only 1 change under subdir: %v", changes)
// }

// func TestTreeDiffShowsFilesAdded(t *testing.T) {
// 	tree1 := createTestTree(t, []testFile{
// 		{File: filetree.File{Path: "llamas.txt", Mode: 0400}, Data: []byte("llamas rock")},
// 		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
// 	})
// 	defer os.RemoveAll(tree1.Dir())

// 	writeFile(t, tree1, "another_file1.txt", []byte("some content1"), 0700)
// 	writeFile(t, tree1, "another_file2.txt", []byte("some content2"), 0600)
// 	writeFile(t, tree1, "another_file3.txt", []byte("some content3"), 0644)

// 	tree2, err := filetree.Scan(tree1.Dir())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	actualChanges := tree1.Diff(tree2)
// 	expectedChanges := filetree.Changes{
// 		filetree.CreateFileChange{Path: "another_file1.txt", Size: 13, Mode: 0700, ModTime: testMtime},
// 		filetree.CreateFileChange{Path: "another_file2.txt", Size: 13, Mode: 0600, ModTime: testMtime},
// 		filetree.CreateFileChange{Path: "another_file3.txt", Size: 13, Mode: 0644, ModTime: testMtime},
// 	}

// 	if err = changesEqual(expectedChanges, actualChanges); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestTreeDiffShowsFilesDeleted(t *testing.T) {
// 	tree1 := createTestTree(t, []testFile{
// 		{File: filetree.File{Path: "llamas.txt", Mode: 0400}, Data: []byte("llamas rock")},
// 		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
// 	})
// 	defer os.RemoveAll(tree1.Dir())

// 	if err := os.RemoveAll(filepath.Join(tree1.Dir(), "subdir")); err != nil {
// 		t.Fatal(err)
// 	}

// 	tree2, err := filetree.Scan(tree1.Dir())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	actualChanges := tree1.Diff(tree2)
// 	expectedChanges := filetree.Changes{
// 		filetree.DeleteFileChange{Path: "subdir/alpacas.txt"},
// 	}

// 	if err = changesEqual(expectedChanges, actualChanges); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestTreeDiffShowsFilesModified(t *testing.T) {
// 	tree1 := createTestTree(t, []testFile{
// 		{File: filetree.File{Path: "llamas.txt", Mode: 0700}, Data: []byte("llamas rock")},
// 		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
// 	})
// 	defer os.RemoveAll(tree1.Dir())

// 	writeFile(t, tree1, "llamas.txt", []byte("some content1"), 0700)

// 	tree2, err := filetree.Scan(tree1.Dir())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	changes := tree1.Diff(tree2)

// 	if len(changes) != 1 {
// 		t.Fatalf("Expected 1 change, got %d", len(changes))
// 	}

// 	ch, ok := changes[0].(filetree.FileContentChange)
// 	if !ok {
// 		t.Fatalf("Change wasn't a FileContentChange: %T", changes[0])
// 	}

// 	if ch.Path != "llamas.txt" {
// 		t.Fatalf("Expected llamas.txt to change, got %s", ch.Path)
// 	}
// }

// func TestApplyingChangesBetweenTrees(t *testing.T) {
// 	tree1 := createTestTree(t, []testFile{
// 		{File: filetree.File{Path: "llamas.txt", Mode: 0700}, Data: []byte("llamas rock")},
// 		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas are ok")},
// 	})
// 	defer os.RemoveAll(tree1.Dir())

// 	tree2 := createTestTree(t, []testFile{
// 		{File: filetree.File{Path: "subdir/alpacas.txt", Mode: 0700}, Data: []byte("alpacas rock")},
// 	})
// 	defer os.RemoveAll(tree1.Dir())

// 	changes := tree1.Diff(tree2)
// 	if err := tree1.Apply(changes); err != nil {
// 		t.Fatal(err)
// 	}

// 	after := tree1.Diff(tree2)
// 	if len(after) > 0 {
// 		t.Fatalf("Expected no changes after, got:\n%s", after.String())
// 	}
// }

// func changesEqual(expected, actual filetree.Changes) error {
// 	c1s := expected.String()
// 	c2s := actual.String()

// 	if !reflect.DeepEqual(c1s, c2s) {
// 		return fmt.Errorf("Unexpected changes:\nExpected:\n%s\nActual:\n%s", c1s, c2s)
// 	}

// 	return nil
// }

// var (
// 	testAtime = time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
// 	testMtime = time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
// )

// func writeFile(t *testing.T, tree *filetree.Tree, key string, data []byte, mode os.FileMode) {
// 	path := filepath.Join(tree.Dir(), key)
// 	if err := ioutil.WriteFile(path, data, mode); err != nil {
// 		t.Fatal(err)
// 	}
// 	if err := os.Chtimes(path, testAtime, testMtime); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func createTestTree(t *testing.T, files []testFile) *filetree.Tree {
// 	dir := createTestDir(t, files)
// 	tree, err := filetree.Scan(dir)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	return tree
// }

// func createTestDir(t *testing.T, files []testFile) string {
// 	dir, err := ioutil.TempDir("", "treetest")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, f := range files {
// 		path := filepath.Join(dir, f.Path)
// 		if f.Mode == os.FileMode(0) {
// 			f.Mode = os.FileMode(0700)
// 		}
// 		if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err = ioutil.WriteFile(path, f.Data, f.Mode); err != nil {
// 			t.Fatal(err)
// 		}
// 		if err := os.Chtimes(path, testAtime, testMtime); err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// 	return dir
// }

// type testFile struct {
// 	filetree.File
// 	Data []byte
// }
