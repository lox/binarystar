package binarystar

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Tree struct {
	sync.RWMutex
	FileSet
	Dir        string
	MatcherSet *MatcherSet
	Changes    chan ChangeSet
}

func NewTree(dir string, m *MatcherSet) (*Tree, error) {
	files, err := Scan(dir, m)
	if err != nil {
		return nil, err
	}

	return &Tree{
		FileSet:    files,
		Dir:        dir,
		MatcherSet: m,
		Changes:    make(chan ChangeSet),
	}, nil
}

func (t *Tree) Watch(ctx context.Context) error {
	return watch(ctx, t)
}

func (t *Tree) Add(fi FileInfo) error {
	fi.Dir = t.Dir
	return t.FileSet.Add(fi)
}

func (t *Tree) Replace(fi FileInfo) error {
	if err := t.Delete(fi.Path); err != nil {
		return err
	}
	return t.FileSet.Add(fi)
}

func (t *Tree) RelativePath(fullPath string) (string, error) {
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(absPath, t.Dir) {
		return "", fmt.Errorf("%s is not a subdir of %s", absPath, t.Dir)
	}

	return strings.TrimPrefix(absPath, t.Dir+string(filepath.Separator)), nil
}

var ErrTreeAlreadyUpToDate = errors.New("Tree already up to date")

func canApplyAddChange(c AddChange, t *Tree) error {
	current, ok := t.Get(c.Path)
	if !ok {
		return nil
	}
	if eq, _ := FileInfoEqual(c.FileInfo, current); eq {
		return ErrTreeAlreadyUpToDate
	}
	return nil
}

func applyAddChangeToFilesystem(c AddChange) error {
	log.Printf("ADD_FILESYSTEM %s", c.Path)
	if err := os.MkdirAll(filepath.Dir(filepath.Join(c.Dir, c.Path)), 0700); err != nil {
		return err
	}
	if err := os.Chtimes(c.File.Name(), c.ModTime, c.ModTime); err != nil {
		return err
	}
	if err := os.Chmod(c.File.Name(), os.FileMode(c.Mode)); err != nil {
		return err
	}
	if err := os.Rename(c.File.Name(), filepath.Join(c.Dir, c.Path)); err != nil {
		return err
	}
	after, err := ScanFile(c.Dir, c.Path)
	if err != nil {
		return err
	}
	if eq, _ := FileInfoEqual(c.FileInfo, after); !eq {
		return fmt.Errorf("ADD(%s) didn't apply correctly", c.Path)
	}
	return nil
}

func applyAddChange(c AddChange, t *Tree) error {
	log.Printf("ADD %s", c.Path)
	if err := t.Add(c.FileInfo); err != nil {
		return err
	}
	t.Changes <- ChangeSet{Add: []AddChange{c}}
	return nil
}

func canApplyModifyChange(c ModifyChange, t *Tree) error {
	current, ok := t.Get(c.To.Path)
	if !ok {
		return ErrNotExists
	}
	if eq, _ := FileInfoEqual(c.To, current); eq {
		return ErrTreeAlreadyUpToDate
	}
	return nil
}

func applyModifyChangeToFilesystem(c ModifyChange) error {
	log.Printf("MODIFY_FILESYSTEM %s", c.To.Path)
	fileName := filepath.Join(c.To.Dir, c.To.Path)
	if err := os.Chtimes(fileName, c.To.ModTime, c.To.ModTime); err != nil {
		return err
	}
	if err := os.Chmod(fileName, os.FileMode(c.To.Mode)); err != nil {
		return err
	}
	after, err := ScanFile(c.To.Dir, c.To.Path)
	if err != nil {
		return err
	}
	if err = applyPatch(filepath.Join(c.To.Dir, c.To.Path), c.Patch, c.To.Fingerprint); err != nil {
		return err
	}
	if eq, _ := FileInfoEqual(c.To, after); !eq {
		return fmt.Errorf("MODIFY_FILESYSTEM(%s) didn't apply correctly", c.To.Path)
	}
	return nil
}

func applyModifyChange(c ModifyChange, t *Tree) error {
	log.Printf("MODIFY %s", c.To.Path)
	if err := t.Replace(c.To); err != nil {
		return err
	}
	t.Changes <- ChangeSet{Modify: []ModifyChange{c}}
	return nil
}

func canApplyDeleteChange(c DeleteChange, t *Tree) error {
	current, ok := t.Get(c.To.Path)
	if !ok {
		return ErrNotExists
	}
	if eq, _ := FileInfoEqual(c.To, current); eq {
		return ErrTreeAlreadyUpToDate
	}
	return nil
}

func applyDeleteChangeToFilesystem(c DeleteChange) error {
	log.Printf("DELETE_FILESYSTEM %s", c.To.Path)
	if err := os.Remove(filepath.Join(c.To.Dir, c.To.Path)); err != nil {
		log.Printf("Error deleting file: %#v", err)
		return err
	}
	return nil
}

func applyDeleteChange(c DeleteChange, t *Tree) error {
	log.Printf("DELETE %s", c.To.Path)
	err := t.Modify(c.To.Path, func(f *FileInfo) error {
		f.IsDeleted = true
		f.ModTime = time.Now()
		return nil
	})
	if err != nil {
		return err
	}
	t.Changes <- ChangeSet{Delete: []DeleteChange{c}}
	return nil
}
