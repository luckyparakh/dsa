// https://leetcode.com/problems/rotting-oranges/
// https://www.youtube.com/watch?v=yf3oUhkvqA0&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=13

// Given a grid of oranges, where 0 represents an empty cell, 1 represents a fresh orange, and 2 represents a rotten orange,
// the code finds the minimum time required for all oranges to become rotten.
// The code uses a breadth-first search (BFS) approach to traverse the grid and update the time for each rotten orange.
// It keeps track of visited nodes to avoid revisiting them.
// If there is at least one fresh orange left after all the rotten oranges have been processed, the code returns -1.
// Otherwise, it returns the maximum time it took for an orange to become rotten.

package main

import "fmt"

type node struct {
	x int
	y int
	t int
}

func orangesRotting(grid [][]int) int {
	s := []node{}
	ans := 0
	vn := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		vn[i] = make([]bool, len(grid[0]))
	}
	findRotten(grid, &s)
	for len(s) != 0 {
		tempNode := s[0]
		i := tempNode.x
		j := tempNode.y
		if !vn[i][j] {
			if tempNode.t > ans {
				ans = tempNode.t
			}
			addNode(i+1, j, tempNode.t+1, grid, &s)
			addNode(i-1, j, tempNode.t+1, grid, &s)
			addNode(i, j+1, tempNode.t+1, grid, &s)
			addNode(i, j-1, tempNode.t+1, grid, &s)
			vn[i][j] = true
		}
		s = s[1:]
	}
	if ifOnePresent(grid, vn) {
		return -1
	}
	return ans
}

func ifOnePresent(grid [][]int, vn [][]bool) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && !vn[i][j] {
				return true
			}
		}
	}
	return false
}
func addNode(i, j, time int, grid [][]int, n *[]node) bool {
	row := len(grid)
	col := len(grid[0])
	if i >= 0 && j >= 0 && i < row && j < col && grid[i][j] == 1 {
		*n = append(*n, node{
			x: i,
			y: j,
			t: time,
		})
	}
	return false
}
func findRotten(grid [][]int, n *[]node) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				*n = append(*n, node{
					x: i,
					y: j,
				})
			}
		}
	}
}

func main() {
	fmt.Println(orangesRotting([][]int{
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}))
}
