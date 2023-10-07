// https://www.youtube.com/watch?v=m9zawMC6QAI
// https://leetcode.com/problems/word-break/
package main

func wordBreak(s string, wordDict []string) bool {
	t := Construct()
	for _, word := range wordDict {
		t.Insert(word)
	}
	var wordBreaker func(s string, t Trie) bool
	// wordBreaker checks if a given string can be broken down into words using a Trie data structure
	wordBreaker = func(s string, t Trie) bool {
		// If the string is empty, it can be broken down into words
		if len(s) == 0 {
			return true
		}
		// Iterate through the string
		for i := 1; i < len(s); i++ {
			// Check if the substring from the beginning to the current index is present in the Trie
			if t.Search(s[:i]) && wordBreaker(s[i:], t) {
				// If the substring is present in the Trie and the remaining string can be broken down into words,
				// return true
				return true
			}
		}
		// If none of the substrings can be broken down into words, return false
		return false
	}
	return wordBreaker(s, t)
}

type Node struct {
	children  [26]*Node
	endOfWord bool
}

type Trie struct {
	root *Node
}

func Construct() Trie {
	return Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string) {
	tNode := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if tNode.children[index] == nil {
			tNode.children[index] = &Node{}
		}
		if i == len(word)-1 {
			tNode.children[index].endOfWord = true
		}
		tNode = tNode.children[index]
	}
}

func (t *Trie) Search(word string) bool {
	tNode := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if tNode.children[index] == nil {
			return false
		}
		if i == len(word)-1 {
			if !tNode.children[index].endOfWord {
				return false
			}
		}
		tNode = tNode.children[index]
	}
	return true
}
