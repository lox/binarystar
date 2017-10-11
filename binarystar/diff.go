package binarystar

import (
	"fmt"
	"os"
)

func FileInfoEqual(f1, f2 FileInfo) (bool, string) {
	if f1.Path != f2.Path {
		return false, fmt.Sprintf("Path wasn't equal: %q vs %q", f1.Path, f2.Path)
	}
	if f1.Size != f2.Size {
		return false, fmt.Sprintf("Size wasn't equal on %s: %v vs %v", f1.Path, f1.Size, f2.Size)
	}
	if f1.Mode != f2.Mode {
		return false, fmt.Sprintf("Mode wasn't equal on %s: %v vs %v", f1.Path, os.FileMode(f1.Mode), os.FileMode(f2.Mode))
	}
	if f1.IsDeleted != f2.IsDeleted {
		return false, fmt.Sprintf("IsDeleted weren't equal")
	}
	if !f1.ModTime.Equal(f2.ModTime) {
		return false, fmt.Sprintf("Modtimes weren't equal")
	}
	if !fingerprintEqual(f1.Fingerprint, f2.Fingerprint) {
		return false, fmt.Sprintf("Fingerprints weren't equal")
	}
	return true, ""
}

// Diff compares two sets and returns the difference as a ChangeSet
func Diff(s1, s2 FileSet, matcher *MatcherSet) ChangeSet {
	changes := ChangeSet{}

	// check for keys that exist in s1 that don't exist in s2 and mark them deleted or modified
	for _, f1 := range s1 {
		if !matcher.Match(f1.Path) {
			continue
		}
		f2, ok := s2.Get(f1.Path)
		if !ok {
			changes.Delete = append(changes.Delete, DeleteChange{From: f1, To: f1.AsDeleted()})
		} else if equal, _ := FileInfoEqual(f1, f2); equal {
			changes.Modify = append(changes.Modify, ModifyChange{From: f1, To: f2})
		}
	}

	// now look for keys in the new tree that aren't in the old set and add them
	for _, f2 := range s2 {
		if !matcher.Match(f2.Path) {
			continue
		}
		if _, ok := s1.Get(f2.Path); !ok {
			changes.Add = append(changes.Add, AddChange{FileInfo: f2})
		}
	}
	return changes
}
