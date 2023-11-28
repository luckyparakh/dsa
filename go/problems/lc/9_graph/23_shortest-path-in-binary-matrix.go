// https://leetcode.com/problems/shortest-path-in-binary-matrix/
// https://www.youtube.com/watch?v=U5Mw4eyUmw4&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=36
package main

import (
	"fmt"
	"math"
)

func shortestPathBinaryMatrix(grid [][]int) int {
	// Initialize distance as maximum integer value
	distance := math.MaxInt

	// Get the number of rows and columns in the grid
	row, col := len(grid), len(grid[0])

	// Create a 2D slice to keep track of visited nodes
	visitedNodes := make([][]bool, row)
	for k := range visitedNodes {
		visitedNodes[k] = make([]bool, col)
	}

	// Define the start and end nodes
	startNode := Node{
		x: 0,
		y: 0,
	}
	endNode := Node{
		x: row - 1,
		y: col - 1,
	}

	// Create a queue to store nodes
	q := []Node{}

	// Check if the start node is valid and add it to the queue
	if grid[startNode.x][startNode.y] == 0 {
		visitedNodes[startNode.x][startNode.y] = true
		q = append(q, startNode)
	}

	// Perform breadth-first search
	for len(q) != 0 {
		// Dequeue the first node from the queue
		node := q[0]

		// Check if the current node is the end node
		if node.x == endNode.x && node.y == endNode.y {
			// Update the minimum distance if necessary
			if node.distance < distance {
				distance = node.distance
			}
		}

		// Define the directions to explore
		dirX := []int{0, 0, -1, -1, -1, 1, 1, 1}
		dirY := []int{1, -1, 0, -1, 1, 0, -1, 1}

		// Explore each direction
		for k := range dirX {
			x, y := node.x-dirX[k], node.y-dirY[k]

			// Check if the next node is valid and not visited
			if x >= 0 && y >= 0 && x < row && y < col && grid[x][y] == 0 && !visitedNodes[x][y] {
				// Enqueue the next node with updated distance
				q = append(q, Node{
					x:        x,
					y:        y,
					distance: node.distance + 1,
				})
				visitedNodes[x][y] = true
			}
		}

		// Dequeue the current node
		q = q[1:]
	}

	// Check if a valid path is found
	if distance == math.MaxInt {
		return -1
	}

	// Return the number of nodes visited while traversing shortest distance
	return distance + 1
}

type Node struct {
	x        int
	y        int
	distance int
}

func shortestPathBinaryMatrixWithStartEndNode(grid [][]int, start, end Node) int {
	// Initialize distance as maximum integer value
	distance := math.MaxInt

	// Get the number of rows and columns in the grid
	row, col := len(grid), len(grid[0])

	// Create a 2D slice to keep track of visited nodes
	visitedNodes := make([][]bool, row)
	for k := range visitedNodes {
		visitedNodes[k] = make([]bool, col)
	}

	// Define the start and end nodes
	startNode := Node{
		x: start.x,
		y: start.y,
	}
	endNode := Node{
		x: end.x,
		y: end.y,
	}

	// Create a queue to store nodes
	q := []Node{}

	// Check if the start node is valid and add it to the queue
	if grid[startNode.x][startNode.y] == 0 {
		visitedNodes[startNode.x][startNode.y] = true
		q = append(q, startNode)
	}

	// Perform breadth-first search
	for len(q) != 0 {
		// Dequeue the first node from the queue
		node := q[0]

		// Check if the current node is the end node
		if node.x == endNode.x && node.y == endNode.y {
			// Update the minimum distance if necessary
			if node.distance < distance {
				distance = node.distance
			}
		}

		// Define the directions to explore
		dirX := []int{0, 0, -1, -1, -1, 1, 1, 1}
		dirY := []int{1, -1, 0, -1, 1, 0, -1, 1}

		// Explore each direction
		for k := range dirX {
			x, y := node.x-dirX[k], node.y-dirY[k]

			// Check if the next node is valid and not visited
			if x >= 0 && y >= 0 && x < row && y < col && grid[x][y] == 0 && !visitedNodes[x][y] {
				// Enqueue the next node with updated distance
				q = append(q, Node{
					x:        x,
					y:        y,
					distance: node.distance + 1,
				})
				visitedNodes[x][y] = true
			}
		}

		// Dequeue the current node
		q = q[1:]
	}

	// Check if a valid path is found
	if distance == math.MaxInt {
		return -1
	}

	// Return the number of nodes visited while traversing shortest distance
	return distance + 1
}

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0}, {1, 1, 0}, {1, 1, 0},
	})) //4
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 1}, {1, 0},
	})) //2
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{1, 0, 0}, {1, 1, 0}, {1, 1, 0},
	})) //-1
	fmt.Println(shortestPathBinaryMatrixWithStartEndNode([][]int{
		{0, 0, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}, {0, 0, 1, 1}, {0, 1, 1, 1},
	}, Node{0, 1, 0}, Node{2, 2, 0})) //3
}
