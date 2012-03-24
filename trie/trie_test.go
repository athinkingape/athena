package trie_test

import (
	"github.com/athinkingape/athena/trie"
	"testing"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	add := []string{
		"abc",
		"ab",
		"a",
		"abcd",
		"aed",
		"aeed",
		"Abc",
		"Ab",
	}

	dne := []string{
		"doesnotexist",
		"doesNotExist",
	}

	tr := &trie.Tree{}

	mustContain := func(s string) {
		if !tr.Contains(s) {
			t.Logf("tree does not contain %q, when it should", s)
			t.Fail()
		}
	}

	mustNotContain := func(s string) {
		if tr.Contains(s) {
			t.Logf("tree contains %q, when it should not", s)
			t.Fail()
		}
	}

	for _, v := range add {
		for _, d := range dne {
			mustNotContain(d)
		}

		mustNotContain(v)
		tr.Add(v)
		mustContain(v)
	}

	for _, d := range dne {
		mustNotContain(d)
	}

	for _, v := range add {
		mustContain(v)
	}
}
