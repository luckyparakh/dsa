// https://leetcode.com/problems/number-of-provinces/
package main

// DS represents a disjoint set data structure
type DS struct {
	size           []int
	ultimateParent []int
}

// Union merges two sets by updating the ultimate parent and size
func (d DS) Union(u, v int) {
	pu, pv := d.findUltimateParent(u), d.findUltimateParent(v)
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
}

// findUltimateParent finds the ultimate parent of a node and performs path compression
func (d DS) findUltimateParent(u int) int {
	if u == d.ultimateParent[u] {
		return u
	}
	d.ultimateParent[u] = d.findUltimateParent(d.ultimateParent[u])
	return d.ultimateParent[u]
}

// isConnected checks if two nodes are connected
func (d DS) isConnected(u, v int) bool {
	return d.ultimateParent[u] == d.ultimateParent[v]
}

// createDS creates a disjoint set with the given node count
func createDS(nodeCount int) DS {
	ds := DS{
		size:           make([]int, nodeCount+1),
		ultimateParent: make([]int, nodeCount+1),
	}
	for i := 0; i < nodeCount+1; i++ {
		ds.size[i] = 1
		ds.ultimateParent[i] = i
	}
	return ds
}

// findCircleNum finds the number of connected components in the given graph
func findCircleNum(isConnected [][]int) int {
	// Create a slice of edges to store connected nodes
	edges := make([][]int, 0)

	// Iterate over the isConnected matrix
	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected[i]); j++ {
			// Check if nodes i and j are connected and not the same node
			if i != j && isConnected[i][j] == 1 {
				// Add the connected nodes to the edges slice
				edges = append(edges, []int{i, j})
			}
		}
	}

	// Initialize a counter variable
	count := 0

	// Create a disjoint set data structure with the length of isConnected matrix
	ds := createDS(len(isConnected))

	// Iterate over the edges slice
	for _, edge := range edges {
		// Check if the two nodes of the edge are not already connected
		if !ds.isConnected(edge[0], edge[1]) {
			// Union the two nodes
			ds.Union(edge[0], edge[1])
		}
	}

	// Count the number of ultimate parents in the disjoint set data structure
	for k, v := range ds.ultimateParent {
		if k == v {
			count++
		}
	}

	// Return the count of ultimate parents
	return count

}
func main() {

}
