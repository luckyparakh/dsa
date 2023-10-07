// https://leetcode.com/problems/longest-duplicate-substring/
// https://www.youtube.com/watch?v=FQ8hcOOzQMU&list=PLEJXowNB4kPyi859E6qGUs7jlpQehJndl&index=7
package main

// Node represents a node in the Trie data structure
type Node struct {
	children [26]*Node // children of the node
	eow      bool      // end of word flag
}

// Trie represents a Trie data structure
type Trie struct {
	root *Node // root node of the Trie
}

// insert inserts a word into the Trie
func (t *Trie) insert(word string) bool {
	tempNode := t.root
	wordPresent := true
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if tempNode.children[index] == nil {
			tempNode.children[index] = &Node{}
			if i == len(word)-1 {
				tempNode.children[index].eow = true
			}
			wordPresent = false
		}
		tempNode = tempNode.children[index]
	}
	if wordPresent && tempNode.eow {
		return true
	}
	return false
}

// parseStringAndStoreTrie parses the input string and stores it in a Trie
// Returns the longest duplicate substring
// Works fine but gives out of memory error due to Trie
func parseStringAndStoreTrie(s string, l int) string {
	t := &Trie{root: &Node{}}
	for i := 0; i+l+1 <= len(s); i++ {
		slicedStr := s[i : i+l+1]
		if t.insert(slicedStr) {
			return slicedStr
		}
	}
	return ""
}

// parseString parses the input string and stores it in a map
// Returns the longest duplicate substring
func parseString(s string, l int) string {
	hm := map[string]bool{}
	for i := 0; i+l+1 <= len(s); i++ {
		slicedStr := s[i : i+l+1]
		if hm[slicedStr] {
			return slicedStr
		} else {
			hm[slicedStr] = true
		}
	}
	return ""
}

// longestDupSubstring finds the longest duplicate substring in the input string
func longestDupSubstring(s string) string {
	start := 0
	end := len(s) - 1
	ans := ""
	for start <= end {
		mid := start + (end-start)/2
		op := parseString(s, mid)
		if op == "" {
			end = mid - 1
		} else {
			ans = op
			start = mid + 1
		}
	}
	return ans
}
