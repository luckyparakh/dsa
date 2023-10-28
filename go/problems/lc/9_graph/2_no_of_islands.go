// https://leetcode.com/problems/number-of-islands/
// https://www.youtube.com/watch?v=muncqlKJrH0&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=8

package main

// This function calculates the number of islands in a grid.
func numIslands(grid [][]byte) int {
	// Initialize a counter variable to keep track of the number of islands.
	c := 0
	
	// Get the number of rows and columns in the grid.
	row, col := len(grid), len(grid[0])
	
	// Create a visited array to keep track of visited cells.
	// Even vn is not needed, we can mark visited in Grid itself.
	vn := make([][]int, row)
	for i := range vn {
		vn[i] = make([]int, col)
	}
	
	// Define a DFS function to traverse the grid.
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// Check if the current cell is within the bounds of the grid,
		// is a land cell, and has not been visited before.
		if i >= 0 && j >= 0 && i < row && j < col && grid[i][j] == '1' && vn[i][j] == 0 {
			// Mark the current cell as visited.
			vn[i][j] = 1
			
			// Recursively call the DFS function on the adjacent cells.
			dfs(i, j-1)
			dfs(i, j+1)
			dfs(i-1, j)
			dfs(i+1, j)
		}
	}
	
	// Iterate through each cell in the grid.
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// If the current cell is a land cell and has not been visited before,
			// increment the counter and call the DFS function.
			if grid[i][j] == '1' && vn[i][j] == 0 {
				c++
				dfs(i, j)
			}
		}
	}
	
	// Return the number of islands.
	return c
}