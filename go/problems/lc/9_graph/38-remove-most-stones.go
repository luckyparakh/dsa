// https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/

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

// removeStones removes stones from the given 2D array based on certain conditions
func removeStones(stones [][]int) int {
	// Get the number of numberStones in the stones array
	numberStones := len(stones)

	// Create a disjoint set data structure with the given number of stones
	ds := createDS(numberStones)

	// Iterate over each stone in the stones array
	for k, stone := range stones {
		// Iterate over each other stone in the stones array
		for j, otherStone := range stones {
			// Check if the current stone is not the same as the other stone
			if k != j {
				// Check if the current stone has the same x-coordinate or y-coordinate as the other stone
				if stone[0] == otherStone[0] || stone[1] == otherStone[1] {
					// Union the current stone and the other stone in the disjoint set data structure
					ds.union(k, j)
				}
			}
		}
	}

	// Initialize a variable to keep track of the number of stones to be removed
	numberStoneRemove := 0

	// Iterate over each element in the size array of the disjoint set data structure
	for i := 0; i < len(ds.size); i++ {
		// Check if the current element is the ultimate parent of its set
		if ds.ultimateParent[i] == i {
			// Increment the number of stones to be removed by the size of the current set minus 1
			numberStoneRemove += ds.size[i] - 1
		}
	}

	// Return the total number of stones to be removed
	return numberStoneRemove
}

func main() {
	fmt.Println(removeStones([][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}})) //5
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}))         //3
	fmt.Println(removeStones([][]int{{0, 0}}))                                         //0
	fmt.Println(removeStones([][]int{{0, 0}, {0, 2}, {1, 3}, {3, 1}, {3, 2}, {4, 3}})) //5
}
