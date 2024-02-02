// https://www.youtube.com/watch?v=Rhxs4k6DyMM

package main

import (
	"fmt"
	"math"
)

func findBridge(V int, adj [][]int) [][]int {
	bridgeEdge := [][]int{}
	edges := make([][]int, V) // Initialize a slice to store the edges of the graph

	// Build the adjacency list representation of the graph
	for i := 0; i < len(adj); i++ {
		edges[adj[i][0]] = append(edges[adj[i][0]], adj[i][1])
		edges[adj[i][1]] = append(edges[adj[i][1]], adj[i][0])
	}

	parent := make([]int, V)
	low := make([]int, V) // Initialize a slice to store the low values of each node
	dis := make([]int, V) // Initialize a slice to store the discovery time of each node

	// Initialize the low, parent and discovery time arrays
	for i := 0; i < V; i++ {
		low[i] = -1
		dis[i] = -1
		parent[i] = -1
	}
	timer := -1
	var dfs func(int, int) int
	dfs = func(node, p int) int {
		if dis[node] == -1 {
			timer += 1
			low[node] = timer
			dis[node] = timer
			parent[node] = p
			for _, neighbor := range edges[node] {
				if neighbor != parent[node] {
					neighborLow := dfs(neighbor, node)
					if neighborLow < low[node] {
						low[node] = neighborLow
					}
					if neighborLow > dis[node] {
						fmt.Println(neighborLow, dis[node], node, neighbor)
						bridgeEdge = append(bridgeEdge, []int{node, neighbor})
					}
				}
			}
			return low[node]
		}
		//back edge
		return int(math.Min(float64(dis[node]), float64(low[node])))
	}
	dfs(0, -1)
	return bridgeEdge
}

func main() {
	fmt.Println(findBridge(5, [][]int{{0, 1}, {0, 2}, {1, 2}, {0, 3}, {3, 4}}))
}
