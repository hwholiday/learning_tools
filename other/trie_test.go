package main

import "testing"

func TestTrie(t *testing.T) {

}

type Trie struct {
	next  map[string]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.isEnd = false
	trie.next = make(map[string]*Trie)
	return trie
}
