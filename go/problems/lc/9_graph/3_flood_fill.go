// Problem: https://leetcode.com/problems/flood-fill/
// Solution: https://www.youtube.com/watch?v=C-2_uSRli8o&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=9

package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// Store the previous color of the starting pixel
	previousColor := image[sr][sc]
	row, col := len(image), len(image[0])

	// If the starting pixel already has the target color, no need to do anything
	if previousColor == color {
		return image
	}

	// Define a DFS function to perform flood fill
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// Check if the current cell is within the boundaries and has the previous color
		if i >= 0 && j >= 0 && i < row && j < col && image[i][j] == previousColor {
			// Change the color of the current cell to the target color
			image[i][j] = color

			// Recursively call the DFS function on the adjacent cells
			dfs(i, j-1) // Left
			dfs(i, j+1) // Right
			dfs(i-1, j) // Up
			dfs(i+1, j) // Down
		}
	}

	// Start the flood fill from the starting pixel
	dfs(sr, sc)

	return image
}