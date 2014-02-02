package trie

import "fmt"

type Trie struct {
	root *node
}

type node struct {
	children map[string]*node
}

func (t *Trie) String() string {
	return fmt.Sprintf("root: %+v", t.root)
}

func (n *node) String() string {
	return fmt.Sprintf("%+v", n.children)
}

func New() *Trie {
	return &Trie{root: &node{}}
}

func (t *Trie) Add(term string) error {
	return t.root.add(term)
}

func (t *Trie) Find(term string) bool {
	return t.root.find(term, 0, len(term))
}

func (n *node) add(term string) error {
	if len(term) == 0 {
		return nil
	}
	key := term[0:1]
	remaining := term[1:]

	if n.children == nil {
		n.children = map[string]*node{}
	}
	var next *node
	if n.children[key] == nil {
		next = &node{}
		n.children[key] = next
	}
	if next == nil {
		next = n.children[key]
	}
	return next.add(remaining)
}

func (n *node) find(term string, index, length int) bool {
	key := term[0 : index+1]
	if length == (index + 1) {
		return key == term
	}
	token := term[index : index+1]

	if n.children == nil {
		return key == term
	}
	if n.children[token] == nil {
		return key == term
	}
	next := n.children[token]
	index += 1
	return next.find(term, index, length)
}
