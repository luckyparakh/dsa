// https://www.geeksforgeeks.org/problems/number-of-islands/1

package main

import "fmt"

type DS struct {
	size           []int
	ultimateParent []int
}

func (d *DS) union(u, v int) bool {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
	if pu == pv {
		return false
	}
	if d.size[pu] >= d.size[pv] {
		d.ultimateParent[pv] = pu
		d.size[u] += d.size[v]
	} else {
		d.ultimateParent[pu] = pv
		d.size[v] += d.size[u]
	}
	return true
}
func (d *DS) findUltimateParent(u int) int {
	if u == d.ultimateParent[u] {
		return u
	}
	d.ultimateParent[u] = d.findUltimateParent(d.ultimateParent[u])
	return d.ultimateParent[u]
}
func createDS(nodes int) *DS {
	size := make([]int, nodes)
	up := make([]int, nodes)
	for i := 0; i < nodes; i++ {
		size[i] = 1
		up[i] = i
	}
	ds := DS{
		size:           size,
		ultimateParent: up,
	}
	return &ds
}

// numOfIslands calculates the number of islands based on the given operators and returns the count for each operator.
func numOfIslands(rows int, cols int, operators [][]int) []int {
	// Create a disjoint set data structure
	ds := createDS(rows * cols)

	// Initialize count and currentCount variables
	count := []int{}
	currentCount := 0

	// Create a matrix with the given number of rows and columns
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	// Iterate through each operator
	for _, operator := range operators {
		x, y := operator[0], operator[1]

		// Check if the cell at (x, y) is not already visited
		if matrix[x][y] == 0 {
			// Define the possible directions to explore
			directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

			// Increment the currentCount assuming a new island will be added
			currentCount++

			// Explore each direction
			for _, direction := range directions {
				nx, ny := x+direction[0], y+direction[1]

				// Check if the new coordinates are within the matrix boundaries and the cell is already part of an island
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && matrix[nx][ny] == 1 {
					// Check if the two cells are in different sets, if so, merge them
					// Calculate the indices of the nodes
					current_node_index := nx*cols + ny
					parent_node_index := x*cols + y
					if ds.union(current_node_index, parent_node_index) {
						// Decrement the currentCount because this cell has connected two islands
						currentCount--
					}
				}
			}

			// Mark the current cell as visited
			matrix[x][y] = 1
		}

		// Append the currentCount to the count slice
		count = append(count, currentCount)
	}

	// Return the count for each operator
	return count
}

func main() {
	fmt.Println(numOfIslands(4, 5, [][]int{{1, 1}, {0, 1}, {3, 3}, {3, 4}})) // [1,1,2,2]
	fmt.Println(numOfIslands(4, 5, [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}})) // [1,2,3,4]
	fmt.Println(numOfIslands(2, 4, [][]int{{1, 3}, {0, 3},
		{0, 1}, {1, 1}, {1, 0}, {1, 2}, {0, 3}, {1, 2}})) // [1,1,2,2,2,1,1,1]
}
