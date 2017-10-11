package binarystar

import (
	"errors"
	"os"
)

var (
	ErrAlreadyInSet = errors.New("file already in set")
	ErrNotExists    = errors.New("file not in set")
)

type FileSet []FileInfo

func (set *FileSet) Add(f FileInfo) error {
	if _, exists := set.Get(f.Path); exists {
		return ErrAlreadyInSet
	}
	*set = append(*set, f)
	return nil
}

func (set *FileSet) Get(key string) (FileInfo, bool) {
	for _, f := range *set {
		if f.Path == key {
			return f, true
		}
	}
	return FileInfo{}, false
}

func (set *FileSet) Len() int {
	return len(*set)
}

func (set *FileSet) Modify(key string, modFunc func(f *FileInfo) error) error {
	for idx := range *set {
		f := &(*set)[idx]
		if f.Path == key {
			return modFunc(f)
		}
	}
	return os.ErrNotExist
}

func (set *FileSet) Delete(key string) error {
	for idx, f := range *set {
		if f.Path == key {
			*set = append((*set)[:idx], (*set)[idx+1:]...)
			return nil
		}
	}
	return os.ErrNotExist
}
