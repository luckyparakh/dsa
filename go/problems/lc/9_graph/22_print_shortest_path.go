// https://www.geeksforgeeks.org/problems/shortest-path-in-weighted-undirected-graph/1

package main

import (
	"fmt"
	"math"
	"sort"
)

func shortestPath(n int, m int, edges [][]int) []int {
	// Create adjacency list
	adj := make([][][]int, n)
	for k := range adj {
		adj[k] = make([][]int, 0)
	}
	// Iterate over each edge in the edges slice
	for _, edge := range edges {
		// Decrement the first element of the edge by 1 to match the index of the adj slice
		node1 := edge[0] - 1
		// Create a new slice with the second element of the edge decremented by 1 and the third element which is weight
		node2 := []int{edge[1] - 1, edge[2]}
		// Append the new node2 slice to the adj slice at index node1
		adj[node1] = append(adj[node1], node2)

		// Add reverse Entry
		node1r := edge[1] - 1
		node2r := []int{edge[0] - 1, edge[2]}
		adj[node1r] = append(adj[node1r], node2r)
	}
	fmt.Println(adj)
	// Initialize distances array with maximum values
	distances := make([]int, n)
	for k := range distances {
		distances[k] = math.MaxInt
	}

	// Initialize visitedNodes array
	visitedNodes := make([]bool, n)

	// Set source node
	S := 0

	// Initialize priority queue
	q := [][]int{}
	q = append(q, []int{S, 0}) // 2D slice containing the starting node S and a distance value of 0

	// Initialize path array
	path := make([][]int, n)
	path[S] = []int{}

	// Set distance of source node to 0
	distances[S] = 0

	for len(q) != 0 {
		// Sort the queue based on distance
		sort.SliceStable(q, func(i, j int) bool {
			// If distances are equal, compare node number
			if q[i][1] == q[j][1] {
				return q[i][0] < q[j][0]
			}
			// Compare distance
			return q[i][1] < q[j][1]
		})

		node := q[0]

		if node[0] == n-1 {
			// If we reach the destination node, update the path and return it
			path[node[0]] = append(path[node[0]], node[0])
			for k := range path[node[0]] {
				path[node[0]][k] = path[node[0]][k] + 1
			}
			return path[node[0]]
		}

		visitedNodes[node[0]] = true

		for _, neighbor := range adj[node[0]] {
			if !visitedNodes[neighbor[0]] {
				distance := node[1] + neighbor[1]

				// Update distance and path if a shorter path is found
				if distances[neighbor[0]] > distance {
					distances[neighbor[0]] = distance
					path_node := path[node[0]]
					path_node = append(path_node, node[0])
					path[neighbor[0]] = path_node
				}

				// Add the neighbor node to the queue with the updated distance
				q = append(q, []int{neighbor[0], distance})
			}
		}

		// Remove the processed node from the queue
		q = q[1:]
	}

	// If no path is found, return -1
	return []int{-1}
}
func main() {
	fmt.Println(shortestPath(5, 7,
		[][]int{{4, 3, 3},
			{2, 1, 4},
			{3, 1, 3},
			{2, 5, 5},
			{5, 1, 5},
			{2, 4, 2},
			{4, 5, 1}})) // {1,5}
	fmt.Println(shortestPath(5, 6,
		[][]int{{1, 2, 2}, {2, 5, 5}, {2, 3, 4}, {1, 4, 1}, {4, 3, 3}, {3, 5, 1}})) // {1,4,3,5}
}
