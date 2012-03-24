// Copyright 2012 A Thinking Ape.  All rights reserved.

/* 
	package trie implements a ternary-search-tree variation of the trie datastructure

	This package is not intended to implement a general purpose trie. It's 
	implementation is specifically intended to support athena's usage 
	patterns
*/
package trie

import (
	"fmt"
	"sync"
	"unicode/utf8"
)

type Node struct {
	Left, Center, Right *Node
	Rune                rune
	Word                bool
}

type Tree struct {
	*Node
	mu sync.RWMutex
}

func (t *Tree) Add(s string) {
	if len(s) < 1 {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	node := &(t.Node)
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if *node == nil {
			*node = &Node{Rune: r}
		}

		if r == (*node).Rune {
			s = s[size:]
			if len(s) > 0 {
				node = &((*node).Center)
			}
		} else if r < (*node).Rune {
			node = &((*node).Left)
		} else {
			node = &((*node).Right)
		}
	}

	(*node).Word = true
}

func (t *Tree) Contains(s string) bool {
	return false
}

var runeBuf = []byte{0, 0, 0, 0}

func (n *Node) printStrings(prefix []byte) {
	if n == nil {
		return
	}

	n.Left.printStrings(prefix)
	utf8.EncodeRune(runeBuf, n.Rune)
	lenBef := len(prefix)
	prefix = append(prefix, runeBuf...)
	if n.Word {
		fmt.Printf("%s\n", prefix)
	}
	n.Center.printStrings(prefix)
	prefix = prefix[:lenBef]
	n.Right.printStrings(prefix)
}

func (n *Node) nodeCount(i *int) {
	if n == nil {
		return
	}
	*i = (*i) + 1
	n.Left.nodeCount(i)
	n.Center.nodeCount(i)
	n.Right.nodeCount(i)
}
