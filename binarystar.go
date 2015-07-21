package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: binarystar (-listen bind | -connect host) [..args] [dir]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	var (
		listen  = flag.String("listen", "", "the ip and port to bind to")
		connect = flag.String("connect", "", "the ip and port to connect to")
		ignore  = flag.String("ignore", "", "ignores files that match the provided pattern")
	)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}

	dir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	matcher := NewMatcherSet()
	if *ignore != "" {
		ignoreMatcher, err := NewRegexMatcher(*ignore)
		if err != nil {
			log.Fatal(err)
		}
		matcher.exclude = append(matcher.exclude, ignoreMatcher)
	}

	watcher := NewWatcher(dir)
	watcher.Matcher = matcher
	if err = watcher.Watch(); err != nil {
		log.Fatal(err)
	}

	if (*listen == "" && *connect == "") || (*listen != "" && *connect != "") {
		log.Fatal("Either -listen or -connect must be used")
	}

	tree, err := NewFileTree(dir, matcher)
	if err != nil {
		log.Fatal(err)
	}

	if *listen != "" {
		NewServer(tree, watcher).Listen(*listen)
	} else {
		NewClient(tree, watcher).Connect(*connect)
	}
}
