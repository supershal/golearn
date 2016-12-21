package main

type node struct {
	children map[byte]*node
	isWord   bool
}

func newNode() *node {
	return &node{
		children: make(map[byte]*node),
		isWord:   false,
	}
}

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{
		newNode(),
	}
}

func (t *trie) add(w string) {
	curr := t.root
	chars := []byte(w)
	for _, b := range chars {
		if _, ok := curr.children[b]; !ok {
			curr.children[b] = newNode()
		}
		curr = curr.children[b]
	}

	curr.isWord = true
}

func (t *trie) delete(w string) bool {

}

func deleteUtil(root *node, w string, index int) bool {
	if index == len(w) {
		if !root.isWord {
			return false
		}
		root.isWord = false
		return len(root.children) == 0
	}
	chars := []byte(w)
	if _, ok = root.children[chars[index]]; !ok {
		return false
	}
	wordshouldbedeleted := deleteUtil(root.children[chars[index]], w, index+1)
	if wordshouldbedeleted {
		delete(root.children, chars[index])
		return len(root.children) == 0
	}
	return false
}
