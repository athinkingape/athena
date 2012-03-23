package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/athinkingape/athena/trie"
)

var file = flag.String("file", "", "the file to read trie data")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	t := &trie.Tree{}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatalf("os.Open error: %v", err)
	}

	reader := bufio.NewReader(f)
	var (
		line     []byte
		isPrefix bool
	)

	line, isPrefix, err = reader.ReadLine()

	for err == nil {
		if isPrefix {
			log.Fatalf("line bigger than buffer")
		}
		if len(line) < 2 {
			fmt.Fprintf(os.Stdout, "skipping \"%s\"\n", line)
		} else {
			t.Add(string(line[1 : len(line)-1]))
		}

		line, isPrefix, err = reader.ReadLine()
	}

	runtime.GC()

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
	//t.Root.PrintStrings()
	//fmt.Fprintf(os.Stdout, "Done building ternary search tree: %d nodes\n", t.Root.NodeCount())
	<-time.After(time.Hour * 24)
}
