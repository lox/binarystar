//go:generate msgp -o=wiretypes_msgp.go -tests=false
package binarystar

import (
	"crypto/sha256"
	"time"
)

type SyncMessage struct {
	Files []FileInfo `msg:"files"`
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
	Delete []DeleteChange `msg:"delete"`
	Modify []ModifyChange `msg:"modify"`
}

func (c ChangeSet) Len() int {
	return len(c.Add) + len(c.Delete) + len(c.Modify)
}

type AddChange struct {
	FileInfo
}

type DeleteChange struct {
	FileInfo
}

type ModifyChange struct {
	FileInfo
}

type Block struct {
	Start      int64             `msg:"start"`
	End        int64             `msg:"end"`
	Checksum32 uint32            `msg:"checksum32"`
	Sha256hash [sha256.Size]byte `msg:"sha256hash"`
	HasData    bool              `msg:"hasdata"`
	RawBytes   []byte            `msg:"rawbytes"`
}

type Fingerprint struct {
	Blocksz  uint32                      `msg:"blocksz"`
	BlockMap map[string]map[string]Block `msg:"blockmap"`
}

type FileStreamRequestsMessage struct {
	Paths []string `msg:"path"`
}

type FileStreamHeaderMessage struct {
	FileInfo
	Error string `msg:"error"`
}
