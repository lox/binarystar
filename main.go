package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lox/binarystar/binarystar"
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

	log.SetFlags(log.Lmicroseconds)

	dir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if err = os.MkdirAll(dir, 0777); err != nil {
		log.Fatal(err)
	}

	matcher := binarystar.NewMatcherSet()
	if *ignore != "" {
		matcher.Exclude(binarystar.MatchPattern(*ignore))
	}

	if (*listen == "" && *connect == "") || (*listen != "" && *connect != "") {
		log.Fatal("Either -listen or -connect must be used")
	}

	daemon, err := binarystar.NewDaemon(dir, matcher)
	if err != nil {
		log.Fatal(err)
	}

	if *listen != "" {
		if err = daemon.Listen(*listen); err != nil {
			log.Fatal(err)
		}
	} else {
		if err = daemon.Connect(*connect); err != nil {
			log.Fatal(err)
		}
	}
}
