package binarystar

import "log"

func FileInfoEqual(f1, f2 FileInfo) bool {
	if f1.Path != f2.Path || f1.Size != f2.Size || f1.Mode != f2.Mode {
		log.Printf("Path, size or mode wasn't equal")
		return false
	}
	if f1.IsDeleted != f2.IsDeleted {
		log.Printf("IsDeleted weren't equal")
		return false
	}
	if !f1.ModTime.Equal(f2.ModTime) {
		log.Printf("Modtimes weren't equal")
		return false
	}
	if !FingerprintEqual(f1.Fingerprint, f2.Fingerprint) {
		log.Printf("Fingerprints weren't equal")
		return false
	}
	return true
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
			changes.Delete = append(changes.Delete, DeleteChange{f1.AsDeleted()})
		} else if !FileInfoEqual(f1, f2) {
			changes.Modify = append(changes.Modify, ModifyChange{f2})
		}
	}

	// now look for keys in the new tree that aren't in the old set and add them
	for _, f2 := range s2 {
		if !matcher.Match(f2.Path) {
			continue
		}
		if _, ok := s1.Get(f2.Path); !ok {
			changes.Add = append(changes.Add, AddChange{f2})
		}
	}
	return changes
}
