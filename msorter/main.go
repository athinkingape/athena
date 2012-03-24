package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func rec(s []string) {
	l := len(s)
	if l == 0 {
		return
	}

	fmt.Println(s[l/2])
	rec(s[:l/2])
	rec(s[(l/2) + 1:])
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
	rec(strs)
}
