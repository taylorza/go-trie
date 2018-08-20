package trie

import "errors"

type trieNode struct {
	*sortedDict
	isWord bool
	value  interface{}
}

func newTrieNode() *trieNode {
	return &trieNode{
		sortedDict: &sortedDict{},
		isWord:     false,
		value:      nil}
}

// Trie stores strings and associated values in a form
// that is optimized for lookup by full word or prefix
type Trie interface {
	Insert(key string, value interface{}) error
	Delete(key string)
	Find(key string) (value interface{}, ok bool)
}

type trieImpl struct {
	root *trieNode
}

// New creates a new instance of a trie
func New() Trie {
	return &trieImpl{root: newTrieNode()}
}

// Insert a key value pair into the trie
func (t *trieImpl) Insert(key string, value interface{}) error {
	curr := t.root
	for _, c := range key {
		n, err := curr.find(c)
		if err != nil {
			n = newTrieNode()
			curr.insert(c, n)
		}
		curr = n
	}
	if curr.isWord {
		return errors.New("key already in trie")
	}
	curr.value = value
	curr.isWord = true
	return nil
}

// Delete removes a key and the associated value from the trie
func (t trieImpl) Delete(key string) {
	deleteNode(t.root, []rune(key), 0)
}

func deleteNode(curr *trieNode, key []rune, index int) bool {
	if index == len(key) {
		if !curr.isWord {
			return false
		}
		// If there are still children the node needs to remain but the
		// isWord is set to false and the value is set to nil for GC
		curr.isWord = false
		curr.value = nil

		// if there are no remaining children in the set, the node can be removed entirely
		return len(curr.set) == 0
	}

	ch := key[index]
	n, err := curr.find(ch)
	if err != nil {
		return false
	}
	canDelete := deleteNode(n, key, index+1)
	if canDelete {
		curr.delete(ch)
		// if there are no more children in the set then we can remove the node
		return len(curr.set) == 0
	}
	return false
}

// Find will find the key in the trie and return the associated value and true
// otherwise it returns false in the second return value
func (t trieImpl) Find(key string) (interface{}, bool) {
	curr := t.root
	for _, c := range key {
		n, err := curr.find(c)
		if err != nil {
			return nil, false
		}
		curr = n
	}
	return curr.value, curr.isWord
}
