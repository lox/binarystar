package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"

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
		debug   = flag.Bool("debug", false, "enable debugging")
	)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}

	if (*listen == "" && *connect == "") || (*listen != "" && *connect != "") {
		log.Fatal("Either -listen or -connect must be used")
	}

	log.SetFlags(log.Lmicroseconds)

	if *debug {
		go startProfiler()
	}

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

	tree, err := binarystar.NewTree(dir, matcher)
	if err != nil {
		log.Fatal(err)
	}

	daemon, err := binarystar.NewDaemon(context.Background(), tree)
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

func startProfiler() {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Running pprof on http://%s/debug/pprof", listener.Addr().String())
	runtime.SetMutexProfileFraction(5)

	log.Fatal(http.Serve(listener, nil))
}
