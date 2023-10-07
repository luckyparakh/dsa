// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/

package main

type WordDictionary struct {
	children [26]*WordDictionary
	eow      bool // eow stands for "end of word"
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	tNode := this
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		// If the child node doesn't exist, create a new node
		if tNode.children[index] == nil {
			tNode.children[index] = &WordDictionary{}
		}
		tNode = tNode.children[index]
	}
	tNode.eow = true
}

func (this *WordDictionary) Search(word string) bool {
	return driver(this, word)
}

func driver(n *WordDictionary, s string) bool {
	if len(s) == 1 {
		if s[0] == '.' {
			// If the character is a dot, search all children nodes for a matching word
			for _, child := range n.children {
				if child != nil && child.eow {
					return true
				}
			}
			return false
		}
		index := s[0] - 'a'
		// If the character is a letter, check if the child node exists and is the end of a word
		if n.children[index] != nil && n.children[index].eow {
			return true
		}
		return false
	}
	if s[0] == '.' {
		// If the character is a dot, recursively search all children nodes for a matching word
		for _, child := range n.children {
			if child != nil && driver(child, s[1:]) {
				return true
			}
		}
		return false
	}
	index := s[0] - 'a'
	// If the character is a letter, recursively search the child node
	if n.children[index] != nil {
		return driver(n.children[index], s[1:])
	}
	return false
}