package binarystar_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/lox/binarystar/binarystar"
)

func TestScanningDirectory(t *testing.T) {
	dir := newTestFileDir(t,
		testFile{Path: "llamas.txt", Mode: 0600, Data: []byte("llamas rock")},
		testFile{Path: "subdir/alpacas.txt", Mode: 0644, Data: []byte("alpacas are ok")},
	)
	defer dir.RemoveAll()

	scanned, err := binarystar.Scan(dir.Dir, binarystar.NewMatcherSet(binarystar.MatchAll))
	if err != nil {
		t.Fatal(err)
	}

	if actual := scanned.Len(); actual != 2 {
		t.Fatalf("Expected length to be 4, got %d", actual)
	}

	for _, file := range dir.Files {
		tf, ok := scanned.Get(file.Path)
		if !ok {
			t.Fatalf("Didn't find %s", file.Path)
		}
		if tf.Mode != uint32(file.Mode) {
			t.Fatalf("Expected Mode to be %v, got %v", tf.Mode, file.Mode)
		}
	}
}

var (
	testAtime = time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	testMtime = time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
)

type testFile struct {
	Path string
	Mode os.FileMode
	Data []byte
}

type testFileDir struct {
	Dir   string
	Files []testFile
	t     *testing.T
}

func newTestFileDir(t *testing.T, file ...testFile) *testFileDir {
	dirPath, err := ioutil.TempDir("", "scantest")
	if err != nil {
		t.Fatal(err)
	}
	resolvedDirPath, err := filepath.EvalSymlinks(dirPath)
	if err != nil {
		t.Fatal(err)
	}
	dir := &testFileDir{
		Dir:   resolvedDirPath,
		Files: []testFile{},
		t:     t,
	}
	for _, file := range file {
		dir.WriteFile(file)
	}
	return dir
}

func (d *testFileDir) RemoveAll() {
	if err := os.RemoveAll(d.Dir); err != nil {
		d.t.Fatal(err)
	}
}

func (d *testFileDir) WriteFile(f testFile) {
	path := filepath.Join(d.Dir, f.Path)
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		d.t.Fatal(err)
	}
	if err := ioutil.WriteFile(path, f.Data, f.Mode); err != nil {
		d.t.Fatal(err)
	}
	if err := os.Chtimes(path, testAtime, testMtime); err != nil {
		d.t.Fatal(err)
	}
}
