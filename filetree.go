package main

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	Path    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

type FileTree struct {
	Files map[string]*File
}

func NewFileTree(dir string, matcher *MatcherSet) (*FileTree, error) {
	tree := &FileTree{Files: map[string]*File{}}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if !matcher.Match(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		tree.Files[strings.TrimPrefix(path, dir+"/")] = &File{
			Path:    info.Name(),
			Size:    info.Size(),
			Mode:    info.Mode(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
		}
		return nil
	})

	return tree, err
}

func (t *FileTree) Len() int {
	return len(t.Files)
}
