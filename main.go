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
	line, isPrefix, err := reader.ReadLine()

	for err == nil {
		if isPrefix {
			log.Fatalf("line bigger than buffer")
		}

		t.Add(string(line))

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
	fmt.Fprintf(os.Stdout, "Done building ternary search tree\n")
	<-time.After(time.Hour * 24)
}
