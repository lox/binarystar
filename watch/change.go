package filetree

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/monmohan/xferspdy"
)

type Changes []Change

func (c Changes) String() string {
	changes := make([]string, len(c))
	for idx := range c {
		changes[idx] = c[idx].String()
	}
	return strings.Join(changes, "\n")
}

type Change interface {
	Apply(tree *Tree) error
	String() string
}

type CreateFileChange struct {
	Path    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
}

func (c CreateFileChange) Apply(tree *Tree) error {
	return tree.add(File{
		Path:    c.Path,
		Size:    c.Size,
		Mode:    c.Mode,
		ModTime: c.ModTime,
	})
}

func (c CreateFileChange) String() string {
	return fmt.Sprintf("CREATE [path=%s, size=%v, mode=%v, modtime=%v]",
		c.Path, c.Size, c.Mode, c.ModTime.UTC().Format(time.RFC3339))
}

type DeleteFileChange struct {
	Path     string
	Previous File
}

func (c DeleteFileChange) Apply(tree *Tree) error {
	return tree.delete(c.Path)
}

func (c DeleteFileChange) String() string {
	return fmt.Sprintf("DELETE [path=%s]", c.Path)
}

type FileModeChange struct {
	Path string
	From os.FileMode
	To   os.FileMode
}

func (c FileModeChange) Apply(tree *Tree) error {
	return tree.change(c.Path, func(f *File) error {
		f.Mode = c.To
		return nil
	})
}

func (c FileModeChange) String() string {
	return fmt.Sprintf("CHANGE MODE [path=%s, from=%v, to=%v]", c.Path, c.From, c.To)
}

type FileModTimeChange struct {
	Path string
	From time.Time
	To   time.Time
}

func (c FileModTimeChange) Apply(tree *Tree) error {
	return tree.change(c.Path, func(f *File) error {
		f.ModTime = c.To
		return nil
	})
}

func (c FileModTimeChange) String() string {
	return fmt.Sprintf("CHANGE MODTIME [path=%s, from=%v, to=%v]",
		c.Path,
		c.From.UTC().Format(time.RFC3339),
		c.To.UTC().Format(time.RFC3339),
	)
}

type FileContentChange struct {
	Path string
	From *xferspdy.Fingerprint
	To   *xferspdy.Fingerprint
}

func (c FileContentChange) Apply(tree *Tree) error {
	log.Printf("%#v", c)
	err := tree.change(c.Path, func(f *File) error {
		f.Fingerprint = c.To
		return nil
	})
	log.Printf("%s", tree.String())
	return err
}

func (c FileContentChange) String() string {
	return fmt.Sprintf("MODIFY CONTENT [path=%s, from=%s, to=%s]",
		c.Path,
		formatFingerprint(c.From),
		formatFingerprint(c.To),
	)
}

func formatFingerprint(fp *xferspdy.Fingerprint) string {
	if fp == nil {
		return "<nil>"
	}
	h := sha1.New()
	for _, hashMap := range fp.BlockMap {
		for hash, _ := range hashMap {
			h.Write(hash[:])
		}
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
