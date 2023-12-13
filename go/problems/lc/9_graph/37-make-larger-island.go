// https://leetcode.com/problems/making-a-large-island/

package main

import "fmt"

type DS struct {
	size           []int
	ultimateParent []int
}

func (d *DS) union(u, v int) {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
	// fmt.Println("Parents", pu, pv)
	if pu == pv {
		return
	}
	if d.size[pu] >= d.size[pv] {
		d.ultimateParent[pv] = pu
		d.size[pu] += d.size[pv]
	} else {
		d.ultimateParent[pu] = pv
		d.size[pv] += d.size[pu]
	}
	// fmt.Println("Size", d.size)
	// fmt.Println("UP", d.ultimateParent)
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

// largestIsland finds the size of the largest island in a grid.
func largestIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	// Create a disjoint set data structure
	ds := createDS(rows * cols)

	// Iterate over each cell in the grid
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			// Check if the cell is part of an island
			if grid[x][y] == 1 {
				directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

				// Check the neighboring cells
				for _, direction := range directions {
					nx, ny := x+direction[0], y+direction[1]

					// Check if the new coordinates are within the matrix boundaries and the cell is already part of an island
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
						// Calculate the indices of the nodes
						neighNodeIndex := nx*cols + ny
						currentNodeIndex := x*cols + y

						// Union the current cell with its neighboring cell
						ds.union(neighNodeIndex, currentNodeIndex)
					}
				}
			}
		}
	}

	maxCount := -1
	zeroPresent := false

	// Iterate over each cell in the grid again
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			// Check if the current cell is a zero and it a candidate to swap
			if grid[x][y] == 0 {
				zeroPresent = true
				count := 1
				// To keep track of visited parent
				parents := map[int]bool{}
				directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

				// Check the neighboring cells
				for _, direction := range directions {
					nx, ny := x+direction[0], y+direction[1]

					// Check if the new coordinates are within the matrix boundaries and the cell is already part of an island
					if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
						up := ds.findUltimateParent(nx*cols + ny)

						// Check if the parent of the neighboring cell has already been counted
						if !parents[up] {
							size := ds.size[up]
							count += size
							parents[up] = true
						}
					}
				}

				// Update the maximum count if necessary
				if count > maxCount {
					maxCount = count
				}
			}
		}
	}

	// Return the size of the largest island or the total number of cells if no zeros are present
	if zeroPresent {
		return maxCount
	}
	return rows * cols
}

func main() {
	fmt.Println(largestIsland([][]int{
		{1, 1},
		{1, 1},
	})) //4
	fmt.Println(largestIsland([][]int{
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 1, 1},
		{0, 0, 1, 1, 1},
	})) //20
	fmt.Println(largestIsland([][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1},
	})) //14
}
