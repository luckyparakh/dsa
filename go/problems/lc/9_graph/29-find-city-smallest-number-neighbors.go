// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
// https://www.youtube.com/watch?v=PwMVNSJ5SLI&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=43

package main

import (
	"fmt"
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Create adjacency matrix with initial values set to infinity
	matrix := make([][]int, n)
	for i := range matrix {
		for j := 0; j < n; j++ {
			matrix[i] = append(matrix[i], math.MaxInt)
		}
	}

	// Populate adjacency matrix with edge weights
	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
		matrix[edge[1]][edge[0]] = edge[2]
	}

	// Set diagonal elements to 0
	for i := 0; i < n; i++ {
		matrix[i][i] = 0
	}

	// Use Floyd-Warshall algorithm to find shortest paths between all pairs of vertices
	for via := 0; via < n; via++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if matrix[i][via] != math.MaxInt && matrix[via][j] != math.MaxInt {
					matrix[i][j] = int(math.Min(float64(matrix[i][j]), float64(matrix[i][via]+matrix[via][j])))
				}
			}
		}
	}

	// Find the city with the minimum number of reachable cities within the distance threshold
	minCount := math.MaxInt
	ans := -1

	// Iterate over each row
	for i := 0; i < n; i++ {
		c := 0

		// Count the number of elements in the row that are less than or equal to distanceThreshold
		for j := 0; j < n; j++ {
			if matrix[i][j] <= distanceThreshold {
				c++
			}
		}

		// Update minCount and ans if the current row has a smaller or equal count
		if c <= minCount {
			minCount = c
			ans = i
		}
	}

	return ans
}

func main() {
	fmt.Println(findTheCity(4, [][]int{
		{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1},
	}, 4)) //3
}
