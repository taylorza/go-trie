package trie

import "errors"

type dictItem struct {
	ch   rune
	node *trieNode
}
type sortedDict struct {
	set []dictItem
}

func (s *sortedDict) insert(c rune, n *trieNode) (int, error) {
	i := s.indexof(c)
	if i >= 0 {
		return i, errors.New("item already in set")
	}
	i = ^i
	if len(s.set) == 0 || i == len(s.set) {
		s.set = append(s.set, dictItem{c, n})
		return i, nil
	}

	s.set = append(s.set, dictItem{})
	copy(s.set[i+1:], s.set[i:])
	s.set[i] = dictItem{c, n}

	return i, nil
}

func (s *sortedDict) delete(c rune) (int, error) {
	i := s.indexof(c)
	if i < 0 {
		return i, errors.New("item not in set")
	}

	copy(s.set[i:], s.set[i+1:])
	s.set[len(s.set)-1].node = nil
	s.set = s.set[:len(s.set)-1]
	return i, nil
}

func (s *sortedDict) find(c rune) (*trieNode, error) {
	i := s.indexof(c)
	if i < 0 {
		return nil, errors.New("item does not exist")
	}
	return s.set[i].node, nil
}

func (s *sortedDict) indexof(c rune) int {
	if len(s.set) == 0 {
		return -1
	}

	l := 0
	h := len(s.set) - 1
	var m int
	for l <= h {
		m = l + (h-l)/2

		if c == s.set[m].ch {
			return m
		} else if c < s.set[m].ch {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return ^l
}
