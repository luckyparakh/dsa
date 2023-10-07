// https://github.com/doocs/leetcode/blob/main/solution/1800-1899/1858.Longest%20Word%20With%20All%20Prefixes/README_EN.md
// https://www.youtube.com/watch?v=m9zawMC6QAI
// This program finds the longest word that can be made by combining prefixes of other words in a given list
package main

import "fmt"

type Node struct {
	// eow (end of word) indicates if a given node represents the end of a word
	eow bool
	// children is an array of pointers to child nodes, one for each letter of the alphabet
	children [26]*Node
	// val is the value of the node (i.e. the letter it represents)
	val byte
}

type Trie struct {
	// root is a pointer to the root node of the trie
	root *Node
}

// insert inserts a string into the trie
func (t *Trie) insert(s string) {
	tNode := t.root
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		// If the child node doesn't exist, create a new node
		if tNode.children[index] == nil {
			tNode.children[index] = &Node{val: s[i]}
		}
		tNode = tNode.children[index]
	}
	tNode.eow = true
}

// search recursively searches the trie for the longest word that can be made by combining prefixes of other words
/*
The given code is a recursive function named "search" that takes a pointer to a Node as input and returns a string. 
The function iterates over the children of the input node using a for loop. 
For each child that is not nil and has the "eow" (end of word) flag set to true, the function recursively calls itself with that child as the input. 
The result of the recursive call is stored in a temporary variable "tAns".
If the length of "tAns" is greater than the length of the current answer "ans", the value of "ans" is updated to "tAns".
After the loop, the function checks if "ans" is still an empty string. 
If it is, the function sets "ans" to be the string representation of the value of the input node. 
Otherwise, it concatenates the string representation of the value of the input node with "ans".
Finally, the function returns the value of "ans".

In summary, the function recursively searches for the longest string formed by concatenating the values of nodes in a trie data structure, 
where each node represents a character and has a boolean flag indicating the end of a word.
*/
func search(n *Node) string {
	ans := ""
	for i := 0; i < len(n.children); i++ {
		if n.children[i] != nil && n.children[i].eow {
			tAns := search(n.children[i])
			if len(tAns) > len(ans) {
				ans = tAns
			}
		}
	}
	if ans == "" {
		ans = string(n.val)
	} else {
		ans = string(n.val) + ans
	}
	return ans
}

// longestWord finds the longest word that can be made by combining prefixes of other words in a given list
func longestWord(words []string) string {
	tr := Trie{root: &Node{}}
	for _, word := range words {
		tr.insert(word)
	}
	return search(tr.root)
}

func main() {
	// Test cases
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Println(longestWord([]string{"k","ki","kir","kira", "kiran"}))
	fmt.Println(longestWord([]string{"abc", "bc", "ab", "qwe"}))
}