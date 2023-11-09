// Problem: https://leetcode.com/problems/surrounded-regions/

package main

import "fmt"

type Node struct {
	x int
	y int
}

func solve(board [][]byte) {
	row := len(board)
	col := len(board[0])

	// Create a 2D slice to keep track of visited nodes
	vn := make([][]bool, row)
	for i := 0; i < row; i++ {
		vn[i] = make([]bool, col)
	}

	// Create a stack to perform DFS
	stack := make([]Node, 0, row*col)

	// Add boundary nodes to the stack
	addBoundaryNode(board, &stack, &vn)

	// Perform DFS on the stack
	for len(stack) != 0 {
		tempNode := stack[0]
		i := tempNode.x
		j := tempNode.y

		// Add adjacent nodes to the stack
		addNode(i+1, j, board, &stack, &vn)
		addNode(i-1, j, board, &stack, &vn)
		addNode(i, j+1, board, &stack, &vn)
		addNode(i, j-1, board, &stack, &vn)

		// Remove the visited node from the stack
		stack = stack[1:]
	}

	// Print the visited nodes
	fmt.Println(vn)

	// Replace 'O' with 'X' for the unvisited nodes
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' && !vn[i][j] {
				board[i][j] = 'X'
			}
		}
	}

	// Print the modified board
	fmt.Println(board)
}

// Function to add an adjacent node to the stack
func addNode(i, j int, grid [][]byte, n *[]Node, vn *[][]bool) bool {
	row := len(grid)
	col := len(grid[0])

	// Check if the node is within the grid, is unvisited, and is 'O'
	if i >= 0 && j >= 0 && i < row && j < col && grid[i][j] == 'O' && !(*vn)[i][j] {
		// Add the node to the stack
		*n = append(*n, Node{
			x: i,
			y: j,
		})

		// Mark the node as visited
		(*vn)[i][j] = true
	}

	return false
}

// Function to add boundary nodes to the stack
func addBoundaryNode(board [][]byte, node *[]Node, vn *[][]bool) {
	row := len(board)
	col := len(board[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// Check if the node is on the boundary and is 'O'
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				if board[i][j] == 'O' {
					// Add the node to the stack
					*node = append(*node, Node{
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

func main() {
	// Test cases
	solve([][]byte{
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
	})
	solve([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'X', 'X', 'O'},
	})
}