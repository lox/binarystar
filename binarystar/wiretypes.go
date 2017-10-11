//go:generate msgp -o=wiretypes_msgp.go -tests=false
package binarystar

import (
	"os"
	"time"
	"unsafe"

	"github.com/lox/xferspdy"
)

type SyncMessage struct {
	Files   []FileInfo `msg:"files"`
	Changes ChangeSet  `msg:"changes"`
}

type SyncResponseMessage struct {
	Changes ChangeSet `msg:"changes"`
	Error   *string   `msg:"error"`
}

type FileInfo struct {
	Dir         string       `msg:"-"`
	Path        string       `msg:"path"`
	Size        int64        `msg:"size"`
	Mode        uint32       `msg:"mode"`
	IsDeleted   bool         `msg:"isdeleted"`
	ModTime     time.Time    `msg:"modtime"`
	Fingerprint *Fingerprint `msg:"fingerprint"`
}

func (f FileInfo) AsDeleted() FileInfo {
	f.IsDeleted = true
	f.ModTime = time.Now()
	return f
}

type ChangeSet struct {
	Add    []AddChange    `msg:"add"`
	Modify []ModifyChange `msg:"modify"`
	Delete []DeleteChange `msg:"delete"`
}

func (c ChangeSet) Len() int {
	return len(c.Add) + len(c.Modify) + len(c.Delete)
}

func (c ChangeSet) Merge(c2 ChangeSet) (c3 ChangeSet) {
	c3.Add = append(c.Add, c2.Add...)
	c3.Modify = append(c.Modify, c2.Modify...)
	c3.Delete = append(c.Delete, c2.Delete...)
	return
}

type AddChange struct {
	FileInfo
	File *os.File `msg:"-"`
}

func (ac AddChange) Apply(t *Tree) error {
	if err := canApplyAddChange(ac, t); err != nil {
		return err
	}
	if err := applyAddChangeToFilesystem(ac); err != nil {
		return err
	}
	return applyAddChange(ac, t)
}

type DeleteChange struct {
	From FileInfo `msg:"from"`
	To   FileInfo `msg:"to"`
}

func (dc DeleteChange) Apply(t *Tree) error {
	if err := canApplyDeleteChange(dc, t); err != nil {
		return err
	}
	if err := applyDeleteChangeToFilesystem(dc); err != nil {
		return err
	}
	return applyDeleteChange(dc, t)
}

type ModifyChange struct {
	From  FileInfo `msg:"from"`
	To    FileInfo `msg:"to"`
	Patch Patch    `msg:"-"`
}

func (mc ModifyChange) Apply(t *Tree) error {
	if err := canApplyModifyChange(mc, t); err != nil {
		return err
	}
	if err := applyModifyChangeToFilesystem(mc); err != nil {
		return err
	}
	return applyModifyChange(mc, t)
}

type Patch struct {
	Blocks []Block `msgp:"blocks"`
}

// Block represent a byte slice from the file. For each block, following are computed.
//
// * Adler-32 and SHA256 checksum,
//
// * Start and End byte pos of the block,
//
// * Whether or not its a data block -If this is a data block, RawBytes will capture the byte data represented by this block
type Block struct {
	Start      int64    `msgp:"start"`
	End        int64    `msgp:"end"`
	Checksum32 uint32   `msgp:"checksum32"`
	Sha256hash [32]byte `msgp:"sha256hash"`
	HasData    bool     `msgp:"hasdata"`
	RawBytes   []byte   `msgp:"rawbytes"`
}

// Fingerprint of a given File, encapsulates the following mapping -
//   Adler-32 hash of Block --> SHA256 hash of Block -->Block
// Also stores the block size and the source
type Fingerprint struct {
	Blocksz  uint32                      `msgp:"blocksz"`
	BlockMap map[string]map[string]Block `msgp:"blockmap"`
	Source   string                      `msgp:"source"`
}

func (f Fingerprint) String() string {
	return (*xferspdy.Fingerprint)(unsafe.Pointer(&f)).String()
}
