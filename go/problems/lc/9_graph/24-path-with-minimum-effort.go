// https://leetcode.com/problems/path-with-minimum-effort/
// https://www.youtube.com/watch?v=0ytpZyiZFhA&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=37

package main

import (
	"fmt"
	"math"
	"sort"
)

type Node struct {
	x      int
	y      int
	effort int
}

func minimumEffortPath(heights [][]int) int {
	// Initialize distance as maximum integer value
	// effort := 0

	// Get the number of rows and columns in the grid
	row, col := len(heights), len(heights[0])

	// Create a 2D slice to keep track of visited nodes
	efforts := make([][]int, row)
	for k := range efforts {
		efforts[k] = make([]int, col)
		for j := 0; j < col; j++ {
			efforts[k][j] = math.MaxInt
		}

	}

	// Create a queue to store nodes
	q := []Node{}

	q = append(q, Node{0, 0, 0})
	efforts[0][0] = 0
	for len(q) != 0 {
		sort.SliceStable(q, func(i, j int) bool {
			return q[i].effort < q[j].effort
		})
		node := q[0]
		// Define the directions to explore
		directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		// Explore each direction
		for _, dir := range directions {
			x, y := node.x-dir[0], node.y-dir[1]
			// Check if the next node is valid and not visited
			if x >= 0 && y >= 0 && x < row && y < col {

				newEffort := int(math.Max(float64(node.effort), math.Abs(float64(heights[x][y]-heights[node.x][node.y]))))

				if efforts[x][y] > newEffort {
					efforts[x][y] = newEffort
					q = append(q, Node{x, y, newEffort})
				}
			}

		}
		// Dequeue the current node
		q = q[1:]
	}
	return efforts[row-1][col-1]
}

func main() {
	fmt.Println(minimumEffortPath([][]int{
		{1, 2, 2}, {3, 8, 2}, {5, 3, 5},
	}))
}
