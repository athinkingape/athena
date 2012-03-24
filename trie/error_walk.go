package trie

import (
	"fmt"
	"sort"
)

// TODO(sanjay) optimize the space of this struct

type activeState struct {
	errValue uint8
	nodeId   int
	parentId int
}

func (a activeState) String() string {
	return fmt.Sprintf("[err:%d node:%d parent:%d]", a.errValue, a.nodeId, a.parentId)
}

type stateSet []activeState

func (a stateSet) Clear() {
	a = a[:0]
}

func (s *stateSet) Add(a activeState) {
	*s = append(*s, a)
}

func (s stateSet) Len() int {
	return len(s)
}

func (s stateSet) Less(i, j int) bool {
	l, m := s[i], s[j]
	if l.errValue < m.errValue {
		return true
	} else if m.errValue < l.errValue {
		return false
	}

	return l.nodeId < m.nodeId
}

func (s stateSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stateSet) Finalize() {
	sort.Sort(s)
}

func newStateSet() stateSet {
	return make([]activeState, 0)
}

func ErrorWalk(r rune, a stateSet, e uint8) stateSet {
	panic("")
}
