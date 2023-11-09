// https://leetcode.com/problems/number-of-enclaves/
package main

import "fmt"

type Node struct {
	x int
	y int
}

func numEnclaves(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	// Create a 2D slice to keep track of visited nodes
	vn := make([][]bool, row)
	for i := 0; i < row; i++ {
		vn[i] = make([]bool, col)
	}

	// Create a queue to perform BFS
	queue := make([]Node, 0, row*col)

	// Add boundary nodes to the stack
	addBoundaryNode(grid, &queue, &vn)

	// Perform BFS on the stack
	for len(queue) != 0 {
		tempNode := queue[0]
		i := tempNode.x
		j := tempNode.y

		// Add adjacent nodes to the stack
		addNode(i+1, j, grid, &queue, &vn)
		addNode(i-1, j, grid, &queue, &vn)
		addNode(i, j+1, grid, &queue, &vn)
		addNode(i, j-1, grid, &queue, &vn)

		// Remove the visited node from the stack
		queue = queue[1:]
	}

	// Print the visited nodes
	fmt.Println(vn)

	// Count the unvisited nodes with values as 1
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 && !vn[i][j] {
				count++
			}
		}
	}

	// Print the modified board
	return count
}

// Function to add an adjacent node to the stack
func addNode(i, j int, grid [][]int, n *[]Node, vn *[][]bool) {
	row := len(grid)
	col := len(grid[0])

	// Check if the node is within the grid, is unvisited, and is 1
	if i >= 0 && j >= 0 && i < row && j < col && grid[i][j] == 1 && !(*vn)[i][j] {
		// Add the node to the stack
		*n = append(*n, Node{
			x: i,
			y: j,
		})

		// Mark the node as visited
		(*vn)[i][j] = true
	}

}

// Function to add boundary nodes to the stack
func addBoundaryNode(board [][]int, q *[]Node, vn *[][]bool) {
	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// Check if the node is on the boundary and is 1
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				if board[i][j] == 1 {
					// Add the node to the stack
					*q = append(*q, Node{
						x: i,
						y: j,
					})

					// Mark the node as visited
					(*vn)[i][j] = true
				}
			}
		}
	}
}
