// Problem: https://leetcode.com/problems/sum-of-prefix-scores-of-strings/

package main

type Node struct {
	count    int
	children [26]*Node
}

type Trie struct {
	root *Node
}

// insert inserts a string into the trie
func (t *Trie) insert(s string) {
	tNode := t.root
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		// If the child node doesn't exist, create a new node
		if tNode.children[index] == nil {
			tNode.children[index] = &Node{}
		}
		tNode.children[index].count++
		tNode = tNode.children[index]
	}
}

// PrefixCount returns the count of all words with the given prefix
func (t *Trie) PrefixCount(s string) int {
	tNode := t.root
	count := 0
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		count += tNode.children[index].count
		tNode = tNode.children[index]
	}
	return count
}

// sumPrefixScores calculates the sum of prefix scores for each word in the given list of words
func sumPrefixScores(words []string) []int {
	t := Trie{
		root: &Node{},
	}
	ans := make([]int, len(words))
	for _, word := range words {
		t.insert(word)
	}
	for _, word := range words {
		ans = append(ans, t.PrefixCount(word))
	}
	return ans
}
