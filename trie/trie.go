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

	// TODO(sanjay): optimize the following three fields into 2
	Rune rune
	Word bool
	Id   int
}

type Tree struct {
	*Node
	sync.RWMutex
	nodeList []*Node
}

func (t *Tree) Add(s string) {
	if len(s) < 1 {
		return
	}

	t.Lock()
	defer t.Unlock()

	node := &(t.Node)
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		if *node == nil {
			n := &Node{Rune: r, Id: len(t.nodeList)}
			t.nodeList = append(t.nodeList, n)
			*node = n
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
	t.RLock()
	defer t.RUnlock()

	node := t.Node
	for len(s) > 0 {
		if node == nil {
			return false
		}

		r, size := utf8.DecodeRuneInString(s)

		if r == node.Rune {
			s = s[size:]
			if len(s) > 0 {
				node = node.Center
			}
		} else if r < (*node).Rune {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return node.Word
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
