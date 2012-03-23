package trie_test

import (
	"github.com/athinkingape/athena/trie"
	"testing"
)

func TestSimple(t *testing.T) {
	tr := &trie.Tree{}

	tr.Add("sanjay")
	tr.Add("santa")
}
