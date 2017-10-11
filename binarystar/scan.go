package binarystar

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/lox/xferspdy"
	"github.com/pkg/errors"
)

// ScanFile scans a single file and returns it as FileInfo
func ScanFile(dir string, key string) (FileInfo, error) {
	path := filepath.Join(dir, key)
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, errors.Wrap(err, "ScanFile failed")
	}
	return FileInfo{
		Dir:         filepath.ToSlash(dir),
		Path:        filepath.ToSlash(key),
		Size:        info.Size(),
		Mode:        uint32(info.Mode()),
		ModTime:     info.ModTime(),
		Fingerprint: fingerprintFromFile(path),
	}, nil
}

// Scan scans a directory and returns a FileSet
func Scan(dir string, matcher *MatcherSet) (FileSet, error) {
	if matcher == nil {
		matcher = NewMatcherSet(MatchAll)
	}
	result := FileSet{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		key := strings.TrimPrefix(path, dir+string(filepath.Separator))
		if !matcher.Match(key) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		return result.Add(FileInfo{
			Dir:         filepath.ToSlash(dir),
			Path:        filepath.ToSlash(key),
			Size:        info.Size(),
			Mode:        uint32(info.Mode()),
			ModTime:     info.ModTime(),
			Fingerprint: fingerprintFromFile(path),
		})
	})

	return result, err
}

func FingerprintFromString(s string) *Fingerprint {
	xfp := xferspdy.NewFingerprintFromReader(strings.NewReader(s), 1024)
	return (*Fingerprint)(unsafe.Pointer(xfp))
}

func fingerprintFromFile(filename string) *Fingerprint {
	return (*Fingerprint)(unsafe.Pointer(xferspdy.NewFingerprint(filename, 1024)))
}

func fingerprintEqual(f1, f2 *Fingerprint) bool {
	if f1 == nil && f2 == nil {
		return true
	}
	if f1 == nil || f2 == nil {
		return false
	}
	return (*xferspdy.Fingerprint)(unsafe.Pointer(f1)).DeepEqual((*xferspdy.Fingerprint)(unsafe.Pointer(f2)))
}

func generatePatch(filename string, from, to *Fingerprint) Patch {
	xfp := (*xferspdy.Fingerprint)(unsafe.Pointer(to))
	xblocks := xferspdy.NewDiff(filename, *xfp)
	blocks := make([]Block, len(xblocks))

	for idx := range xblocks {
		block := (*Block)(unsafe.Pointer(&xblocks[idx]))
		blocks[idx] = *block
	}

	return Patch{Blocks: blocks}
}

func applyPatch(filename string, patch Patch, expected *Fingerprint) error {
	xblocks := make([]xferspdy.Block, len(patch.Blocks))

	for idx := range patch.Blocks {
		xblock := (*xferspdy.Block)(unsafe.Pointer(&patch.Blocks[idx]))
		xblocks[idx] = *xblock
	}

	tmpfile, err := ioutil.TempFile("", "patch")
	if err != nil {
		return err
	}
	if err = xferspdy.PatchFile(xblocks, filename, tmpfile); err != nil {
		return err
	}

	log.Printf("Applied patch of %d blocks to %s (%s)", len(xblocks), filename, tmpfile.Name())

	if err = tmpfile.Close(); err != nil {
		return err
	}

	newFingerprint := fingerprintFromFile(tmpfile.Name())

	if !fingerprintEqual(newFingerprint, expected) {
		return fmt.Errorf("Fingerprint didn't match after patch was applied")
	}

	if err = os.Rename(tmpfile.Name(), filename); err != nil {
		return err
	}

	return nil
}
