// https://leetcode.com/problems/rotting-oranges/

package main

type node struct {
	x int
	y int
	t int
}

func orangesRotting(grid [][]int) int {
	s := []node{}
	ans := 0
	vn := [][]bool{}
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
	if i > 0 && j > 0 && i < row && j < col && grid[i+1][j] == 1 {
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
