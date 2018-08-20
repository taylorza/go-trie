package trie

import "testing"

func TestTrieInsert(t *testing.T) {
	d := New()

	d.Insert("abc", true)
	d.Insert("abg", true)
	d.Insert("abcd", true)
	d.Insert("xyz", true)
}

func TestTrieDelete(t *testing.T) {
	d := New()

	d.Insert("abc", true)
	d.Insert("abg", true)

	d.Delete("abc")
	if _, ok := d.Find("abc"); ok {
		t.Fail()
	}

	if _, ok := d.Find("abg"); !ok {
		t.Fail()
	}

	d.Delete("abg")
	if _, ok := d.Find("abg"); ok {
		t.Fail()
	}
}

func TestTrieFind(t *testing.T) {
	d := New()

	d.Insert("abc", "abc")
	d.Insert("abg", "abg")

	if v, ok := d.Find("abc"); v != "abc" || !ok {
		t.Fail()
	}

	if _, ok := d.Find("ab"); ok {
		t.Fail()
	}

	if v, ok := d.Find("abg"); v != "abg" || !ok {
		t.Fail()
	}

	d.Insert("ab", "ab")
	if v, ok := d.Find("ab"); v != "ab" || !ok {
		t.Fail()
	}
}

func TestSortedDictInsert(t *testing.T) {
	var s sortedDict
	s.insert('g', nil)
	if s.set[0].ch != 'g' {
		t.Fail()
	}

	s.insert('c', nil)
	if s.set[0].ch != 'c' {
		t.Fail()
	}

	s.insert('e', nil)
	if s.set[1].ch != 'e' {
		t.Fail()
	}

	s.insert('j', nil)
	if s.set[3].ch != 'j' {
		t.Fail()
	}
}

func TestSortedDictDelete(t *testing.T) {
	var s sortedDict
	a := &trieNode{}
	b := &trieNode{}
	c := &trieNode{}

	s.insert('a', a)
	s.insert('b', b)
	s.insert('c', c)

	i, err := s.delete('b')
	if i != 1 || err != nil {
		t.Fail()
	}

	if s.indexof('a') != 0 || s.indexof('c') != 1 {
		t.Fail()
	}

	i, err = s.delete('c')
	if s.indexof('a') != 0 || s.indexof('c') >= 0 {
		t.Fail()
	}

	i, err = s.delete('a')
	if len(s.set) != 0 {
		t.Fail()
	}
}

func TestSortedDictFind(t *testing.T) {
	var s sortedDict
	a := &trieNode{}
	b := &trieNode{}
	c := &trieNode{}

	s.insert('a', a)
	s.insert('c', c)
	s.insert('b', b)

	if n, _ := s.find('a'); n != a {
		t.Fail()
	}

	if n, _ := s.find('b'); n != b {
		t.Fail()
	}

	if n, _ := s.find('c'); n != c {
		t.Fail()
	}

	if _, err := s.find('d'); err == nil {
		t.Fail()
	}
}

func TestSortedDictIndexOf(t *testing.T) {
	var s sortedDict
	s.set = append(s.set,
		dictItem{'a', nil},
		dictItem{'b', nil},
		dictItem{'c', nil},
		dictItem{'d', nil},
		dictItem{'f', nil},
		dictItem{'i', nil},
		dictItem{'j', nil},
		dictItem{'k', nil},
		dictItem{'l', nil},
		dictItem{'m', nil},
		dictItem{'o', nil})

	if s.indexof('a') != 0 {
		t.Fail()
	}

	if s.indexof('b') != 1 {
		t.Fail()
	}

	if s.indexof('k') != 7 {
		t.Fail()
	}

	if s.indexof('o') != 10 {
		t.Fail()
	}

	if s.indexof('e') != ^4 {
		t.Fail()
	}

	if s.indexof('p') != ^11 {
		t.Fail()
	}
}
