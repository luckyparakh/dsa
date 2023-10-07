// https://leetcode.com/problems/word-search-ii/
// https://www.youtube.com/watch?v=EmvsBM7o-5k
// https://www.youtube.com/watch?v=asbcE9mZz_U

package main

// Node represents a node in the Trie
type Node struct {
	eow      bool      // Flag indicating if the node represents the end of a word
	children [26]*Node // Array of child nodes
	char     byte      // Character value of the node
	val      string    // String value formed by traversing the Trie
}

// Trie represents a Trie data structure
type Trie struct {
	root *Node // Root node of the Trie
}

// insert inserts a word into the Trie
func (t *Trie) insert(s string) {
	tNode := t.root
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if tNode.children[index] == nil {
			// Create a new child node if it doesn't exist
			tNode.children[index] = &Node{
				char: s[i],
				val:  tNode.val + string(s[i]),
			}
		}
		tNode = tNode.children[index]
	}
	tNode.eow = true
}

// findWords finds all the words from the given board that exist in the Trie
func findWords(board [][]byte, words []string) []string {
	t := Trie{root: &Node{}}
	mp := map[string]bool{} // Map to store the found words
	for _, word := range words {
		t.insert(word) // Insert each word into the Trie
	}

	/*
	This is a depth-first search (DFS) function that is used to search for words on a given board. 
	The function takes a 2D board, a node in a trie data structure, and the current position (i, j) on the board as input.
	The function first checks for base cases to terminate the search. 
	If the current position is out of bounds or the current cell on the board is marked as visited ('#'), the function returns. It also checks if the current node in the trie has a child corresponding to the letter at the current position on the board. If not, the function returns.
	If the current node represents the end of a word (eow), it adds the word to a map (mp) if it is not already present.
	The function then marks the current cell as visited by setting its value to '#' and recursively searches in all four directions: up, down, left, and right. Each recursive call is made with the child node corresponding to the letter at the current position on the board.
	After the recursive calls, the function restores the original value of the current cell.
	The purpose of this function is to perform a depth-first search on the board to find all valid words that can be formed from the given trie data structure. The words are added to a map to avoid duplicates.
	*/
	var dfs func(board [][]byte, node *Node, i, j int)
	dfs = func(board [][]byte, node *Node, i, j int) {
		// Base cases for terminating the search
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == '#' {
			return
		}

		index := board[i][j] - 'a'
		if node.children[index] == nil {
			return
		}

		if node.children[index].eow {
			// If the current node represents the end of a word, add it to the map
			if !mp[node.children[index].val] {
				mp[node.children[index].val] = true
			}
		}

		tempVal := board[i][j]
		board[i][j] = '#' // Mark the current cell as visited

		// Recursively search in all four directions
		dfs(board, node.children[index], i+1, j)
		dfs(board, node.children[index], i-1, j)
		dfs(board, node.children[index], i, j+1)
		dfs(board, node.children[index], i, j-1)

		board[i][j] = tempVal // Restore the original value of the current cell
	}

	// Perform DFS on each cell of the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, t.root, i, j)
		}
	}

	// Convert the map of found words to a slice
	ans := []string{}
	for k := range mp {
		ans = append(ans, k)
	}

	return ans
}
