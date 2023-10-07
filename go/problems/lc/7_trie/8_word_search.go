// https://leetcode.com/problems/word-search/description/
package main

// Function to check if a word exists in the given board
func exist(board [][]byte, word string) bool {
	// Iterate through each cell in the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// If the word is found, return true
			if dfs(board, word, i, j) {
				return true
			}
		}
	}
	// If the word is not found, return false
	return false
}

// Depth-first search function to check if the word can be formed from the current cell
func dfs(board [][]byte, word string, i, j int) bool {
	// Base cases:
	// 1. If the current cell is out of bounds or has already been visited, return false
	// 2. If the current cell does not match the first character of the word, return false
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '#' || board[i][j] != word[0] {
		return false
	}
	// If the word has only one character left, return true
	if len(word) == 1 {
		return true
	}
	// Store the value of the current cell
	tempVal := board[i][j]
	// Mark the current cell as visited by replacing its value with '#'
	board[i][j] = '#'
	// Recursively check if the word can be formed by moving in any of the four directions
	if dfs(board, word[1:], i+1, j) || dfs(board, word[1:], i-1, j) ||
		dfs(board, word[1:], i, j+1) || dfs(board, word[1:], i, j-1) {
		// If the word can be formed, restore the original value of the current cell and return true
		board[i][j] = tempVal
		return true
	}
	// If the word cannot be formed, restore the original value of the current cell and return false
	board[i][j] = tempVal
	return false
}