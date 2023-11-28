package main

import (
	"fmt"
	"math"
	"sort"
)

// dijkstra implements Dijkstra's algorithm to find the shortest path from a source node to all other nodes in a graph.
// It takes the number of vertices (V), the adjacency list (adj), and the source node (S) as input and returns an array of distances.
func dijkstra(V int, adj [][][]int, S int) []int {
	// Initialize distances array with maximum values
	distances := make([]int, V)
	for k := range distances {
		distances[k] = math.MaxInt
	}

	// Initialize visitedNodes array
	visitedNodes := make([]bool, V)

	// Initialize priority queue
	q := [][]int{}
	// 2D slice containing the starting node  S  and a distance value of 0
	q = append(q, []int{S, 0})

	// Set distance of source node to 0
	distances[S] = 0

	// Run Dijkstra's algorithm
	for len(q) != 0 {
		// Sort the priority queue based on distance
		sort.SliceStable(q, func(i, j int) bool {
			// If distances are equal, Compares Node Number
			if q[i][1] == q[j][1] {
				return q[i][0] < q[j][0]
			}
			// Compares distance
			return q[i][1] < q[j][1]
		})

		// Get the node with the minimum distance
		node := q[0]

		// Mark the node as visited
		visitedNodes[node[0]] = true

		// Update distances of neighboring nodes
		for _, neighbor := range adj[node[0]] {
			// Check if the neighbor node has not been visited
			if !visitedNodes[neighbor[0]] {
				// Calculate the distance from the current node to the neighbor node
				distance := node[1] + neighbor[1]

				// Update the distance if it is smaller than the current distance
				if distances[neighbor[0]] > distance {
					distances[neighbor[0]] = distance
				}

				// Add the neighbor node to the queue with the updated distance
				q = append(q, []int{neighbor[0], distance})
			}
		}

		// Remove the processed node from the queue
		q = q[1:]
	}

	return distances
}

func main() {
	adj := [][][]int{
		{{1, 10}, {2, 5}},
		{{3, 1}},
		{{1, 3}, {3, 9}},
		{{2, 2}},
	}

	S := 0

	distances := dijkstra(len(adj), adj, S)
	fmt.Println(distances) // {0, 8, 5, 9}
}
