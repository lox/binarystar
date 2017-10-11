package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/lox/binarystar/binarystar"
)

func main() {
	dir1 := os.Args[1]
	dir2 := os.Args[2]

	log.Printf("Verifying that %s and %s are identical", dir1, dir2)
	fs1, err := binarystar.Scan(dir1, binarystar.NewMatcherSet(binarystar.MatchAll))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d files in %s", len(fs1), dir1)
	fs2, err := binarystar.Scan(dir2, binarystar.NewMatcherSet(binarystar.MatchAll))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d files in %s", len(fs2), dir2)
	changes := binarystar.Diff(fs1, fs2, binarystar.NewMatcherSet(binarystar.MatchAll))

	log.Printf("Found %d difference(s)", changes.Len())
	if changes.Len() > 0 {
		spew.Dump(changes)
	}
}
