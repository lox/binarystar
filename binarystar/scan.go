package binarystar

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/monmohan/xferspdy"
)

func ScanFile(dir string, key string) (FileInfo, error) {
	path := filepath.Join(dir, key)
	info, err := os.Stat(path)
	if err != nil {
		return FileInfo{}, err
	}
	return FileInfo{
		Dir:         filepath.ToSlash(dir),
		Path:        filepath.ToSlash(key),
		Size:        info.Size(),
		Mode:        uint32(info.Mode()),
		ModTime:     info.ModTime(),
		Fingerprint: encodeFingerprint(xferspdy.NewFingerprint(path, 1024)),
	}, nil
}

func Scan(dir string, matcher *MatcherSet) (FileSet, error) {
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
			Fingerprint: encodeFingerprint(xferspdy.NewFingerprint(path, 1024)),
		})
	})

	return result, err
}

func FingerprintFromString(s string) *Fingerprint {
	xfp := xferspdy.NewFingerprintFromReader(strings.NewReader(s), 1024)
	return encodeFingerprint(xfp)
}

func FingerprintEqual(f1, f2 *Fingerprint) bool {
	if f1 == nil && f2 == nil {
		return true
	}
	if f1 == nil || f2 == nil {
		return false
	}
	xfp1, err := decodeFingerprint(f1)
	if err != nil {
		log.Println(err)
		return false
	}
	xfp2, err := decodeFingerprint(f2)
	if err != nil {
		log.Println(err)
		return false
	}
	return xfp1.DeepEqual(xfp2)
}

// for our wireformat, we need to convert some of the internals of the fingerprint
func encodeFingerprint(xfp *xferspdy.Fingerprint) *Fingerprint {
	fp := &Fingerprint{
		Blocksz:  xfp.Blocksz,
		BlockMap: map[string]map[string]Block{},
	}

	// Adler-32 hash of Block --> SHA256 hash of Block -->Block
	for adlerHash, blocksMap := range xfp.BlockMap {
		newBlocksMap := map[string]Block{}

		for sha256Bytes, xblock := range blocksMap {
			newBlocksMap[hex.EncodeToString(sha256Bytes[:])] = Block{
				Start:      xblock.Start,
				End:        xblock.End,
				Checksum32: xblock.Checksum32,
				Sha256hash: xblock.Sha256hash,
				HasData:    xblock.HasData,
				RawBytes:   xblock.RawBytes,
			}
		}

		fp.BlockMap[fmt.Sprintf("%d", adlerHash)] = newBlocksMap
	}

	return fp
}

func decodeFingerprint(fp *Fingerprint) (*xferspdy.Fingerprint, error) {
	xfp := &xferspdy.Fingerprint{
		Blocksz:  fp.Blocksz,
		BlockMap: map[uint32]map[[sha256.Size]byte]xferspdy.Block{},
	}

	// Adler-32 hash of Block --> SHA256 hash of Block -->Block
	for adlerHashStr, blocksMap := range fp.BlockMap {
		xBlocksMap := map[[sha256.Size]byte]xferspdy.Block{}

		for sha256BytesStr, block := range blocksMap {
			decoded, err := hex.DecodeString(sha256BytesStr)
			if err != nil {
				return nil, err
			}

			var sha256Bytes [sha256.Size]byte
			copy(sha256Bytes[:], decoded[:sha256.Size])

			xBlocksMap[sha256Bytes] = xferspdy.Block{
				Start:      block.Start,
				End:        block.End,
				Checksum32: block.Checksum32,
				Sha256hash: block.Sha256hash,
				HasData:    block.HasData,
				RawBytes:   block.RawBytes,
			}
		}

		adlerHash, err := strconv.ParseUint(adlerHashStr, 10, 32)
		if err != nil {
			return nil, err
		}

		xfp.BlockMap[uint32(adlerHash)] = xBlocksMap
	}

	return xfp, nil
}
