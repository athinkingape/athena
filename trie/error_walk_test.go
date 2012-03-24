package trie

import (
	"testing"
)

func TestWalk(t *testing.T) {
	// t.Parallel()

	s := newStateSet()
	s.Add(activeState{0, 1, 1})
	s.Add(activeState{2, 0, 0})
	s.Add(activeState{1, 1, 1})

	s.Finalize()
	t.Logf("%v", s)
}
