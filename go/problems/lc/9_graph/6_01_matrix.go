// https://leetcode.com/problems/01-matrix/
// https://www.youtube.com/watch?v=edXdVwkYHF8&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=13
package main

import "fmt"

type node struct {
	x int
	y int
	t int
}

/* 
The function takes a matrix as input and returns a new matrix with the same dimensions, 
where each element represents the shortest distance to the nearest zero element in the input matrix.
The code uses a stack to perform a breadth-first search (BFS) on the matrix. 
It starts by initializing the stack with all the zero elements in the matrix and marking them as visited.
Then, it iteratively processes each element in the stack, updating the distance value in the output matrix 
and adding neighboring elements to the stack if they haven't been visited yet.
Finally, the updated matrix is returned as the result.
Note: The code assumes that the input matrix is a 2D array of integers.
*/

func updateMatrix(mat [][]int) [][]int {
	row := len(mat)
	col := len(mat[0])
	stack := make([]node, 0, row*col)
	vn := make([][]bool, row)
	for i := 0; i < row; i++ {
		vn[i] = make([]bool, col)
	}
	ans := make([][]int, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] == 0 {
				stack = append(stack, node{
					x: i,
					y: j,
				})
				vn[i][j] = true
			}
		}
	}
	// fmt.Println(len(stack))
	// fmt.Println(vn)

	for len(stack) != 0 {
		tempNode := stack[0]
		i := tempNode.x
		j := tempNode.y
		if tempNode.t > 0 {
			ans[i][j] = tempNode.t
		}
		addNode(i+1, j, tempNode.t+1, mat, &stack, &vn)
		addNode(i-1, j, tempNode.t+1, mat, &stack, &vn)
		addNode(i, j+1, tempNode.t+1, mat, &stack, &vn)
		addNode(i, j-1, tempNode.t+1, mat, &stack, &vn)
		stack = stack[1:]
	}
	return ans
}

// This function adds a node to the list of nodes if the given coordinates are within the grid boundaries 
// and the node has not been visited before.
// Parameters:
// - i: The row index of the node.
// - j: The column index of the node.
// - time: The time value of the node.
// - grid: The 2D grid.
// - n: A pointer to the list of nodes.
// - vn: A pointer to the 2D boolean matrix representing visited nodes.
// Returns:
// - bool: Always returns false.

func addNode(i, j, time int, grid [][]int, n *[]node, vn *[][]bool) bool {
	row := len(grid)
	col := len(grid[0])
	if i >= 0 && j >= 0 && i < row && j < col && !(*vn)[i][j] {
		*n = append(*n, node{
			x: i,
			y: j,
			t: time,
		})
		(*vn)[i][j] = true
	}
	return false
}

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0}, {0, 1, 0}, {1, 1, 1},
	}))
}
