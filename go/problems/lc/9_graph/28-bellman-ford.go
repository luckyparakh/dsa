// https://practice.geeksforgeeks.org/problems/distance-from-the-source-bellman-ford-algorithm/0
package main

import (
	"fmt"
)

// bellmanFord calculates the shortest distances from a source vertex to all other vertices in a graph using the Bellman-Ford algorithm.
// V is the number of vertices in the graph.
// edges is a 2D slice representing the edges in the graph. Each edge is represented as [u, v, w], where u and v are the vertices connected by the edge and w is the weight of the edge.
// S is the source vertex.
// The function returns a slice of distances from the source vertex to all other vertices.
func bellman_ford(V int, edges [][]int, S int) []int {
	inf := 100000000
	distances := make([]int, V)
	
	// Initialize all distances to infinity except for the source vertex
	for i := 0; i < V; i++ {
		distances[i] = inf
	}
	distances[S] = 0
	
	// Relax edges V-1 times
	for i := 0; i < V-1; i++ {
		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]
			
			// Skip if the distance to the source vertex is infinity
			if distances[u] == inf {
				continue
			}
			
			// Update the distance if a shorter path is found
			if distances[u]+w < distances[v] {
				distances[v] = distances[u] + w
			}
		}
	}
	
	// Check for negative cycles
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		
		// If a shorter path is found, there is a negative cycle
		if distances[u] != inf && distances[u]+w < distances[v] {
			return []int{-1}
		}
	}
	
	return distances
}

func main() {
	edges := [][]int{
		{0, 1, 5},
		{1, 0, 3},
		{1, 2, -1},
		{2, 0, 1},
	}
	fmt.Println(bellman_ford(3, edges, 2)) // [1 6 0]

	fmt.Println(bellman_ford(2, [][]int{
		{0, 1, 9},
	}, 0)) // [0 9]

	fmt.Println(bellman_ford(3, [][]int{
		{0, 1, 9},
		{1, 2, 3},
		{2, 0, -1},
	}, 0))
}
