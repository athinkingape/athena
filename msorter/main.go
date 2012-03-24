package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func rec(s []string, w *bufio.Writer) {
	l := len(s)
	if l == 0 {
		return
	}

	fmt.Fprintf(w, "%s\n", s[l/2])
	rec(s[:l/2], w)
	rec(s[(l/2)+1:], w)
}

func main() {
	strs := []string{}
	reader := bufio.NewReader(os.Stdin)

	line, isPrefix, err := reader.ReadLine()
	for err == nil {
		if isPrefix {
			log.Fatalf("line bigger than buffer")
		}
		if len(line) >= 2 && line[0] == '"' && line[len(line)-1] == '"' {
			line = line[1 : len(line)-1]
		}
		strs = append(strs, string(line))

		line, isPrefix, err = reader.ReadLine()
	}

	sort.Strings(strs)

	b := bufio.NewWriter(os.Stdout)
	defer b.Flush()

	rec(strs, b)
}
