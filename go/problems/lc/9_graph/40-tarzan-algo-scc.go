// https://www.youtube.com/watch?v=ZeDNSeilf-Y&t=1622s

package main

import (
	"fmt"
	"math"
)

/*
## Summary
This code implements Tarjan's algorithm to find the strongly connected components (SCCs) in a directed graph.

## Example Usage
```go
fmt.Println(tarzanAlgo(5, [][]int{{1, 0}, {0, 2}, {2, 1}, {0, 3}, {3, 4}}))         // [[4] [3] [1 2 0]] , 3
fmt.Println(tarzanAlgo(6, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {4, 0}, {4, 5}})) // [[2] [5] [4 3 1 0]] , 3
```

## Code Analysis
### Inputs
- `V` (integer): The number of vertices in the graph.
- `adj` (2D integer slice): The adjacency list representation of the graph, where each element `adj[i]` is a pair `[u, v]` representing a directed edge from vertex `u` to vertex `v`.
___
### Flow
1. Create an empty slice `scc` to store the strongly connected components.
2. Create an empty slice `edges` to store the edges of the graph.
3. Build the adjacency list representation of the graph by iterating over the input `adj` and adding the edges to the `edges` slice.
4. Create two empty slices `low` and `dis` to store the low values and discovery time of each node.
5. Initialize the `low` and `dis` slices with -1.
6. Create an empty stack to perform depth-first search (DFS) and an empty boolean array `inStack` to keep track of nodes in the stack.
7. Define a recursive function `dfs` that takes a node and a timer as input.
8. In the `dfs` function, if the node has not been visited yet, set its low and discovery time to the current timer value, mark it as in the stack, and push it to the stack.
9. Iterate over the neighbors of the node and recursively call `dfs` on each neighbor, updating the low value of the node if necessary.
10. If the node is the head of a strongly connected component, pop nodes from the stack until the current node is reached, adding the popped nodes to the SCC.
11. Return the low value of the node.
12. If the node is in the stack, it is a back edge, so return its discovery time.
13. If the node is not in the stack, it is a cross edge, so return infinity.
14. Call the `dfs` function starting from node 0.
15. Print the strongly connected components.
16. Return the number of SCCs.
___
### Outputs
- The strongly connected components of the graph, represented as a 2D integer slice.
- The number of strongly connected components in the graph.
___

*/

func tarzanAlgo(V int, adj [][]int) int {
	scc := make([][]int, 0)   // Initialize a slice to store strongly connected components
	edges := make([][]int, V) // Initialize a slice to store the edges of the graph

	// Build the adjacency list representation of the graph
	for i := 0; i < len(adj); i++ {
		edges[adj[i][0]] = append(edges[adj[i][0]], adj[i][1])
	}

	low := make([]int, V) // Initialize a slice to store the low values of each node
	dis := make([]int, V) // Initialize a slice to store the discovery time of each node

	// Initialize the low and discovery time arrays
	for i := 0; i < V; i++ {
		low[i] = -1
		dis[i] = -1
	}

	stack := make([]int, 0)    // Initialize a stack to perform depth-first search
	inStack := make([]bool, V) // Initialize a boolean array to keep track of nodes in the stack

	var dfs func(int, int) int
	dfs = func(node int, timer int) int {
		if dis[node] == -1 { // If the node has not been visited yet
			low[node] = timer // Set the low and discovery time of the node
			dis[node] = timer
			inStack[node] = true        // Mark the node as in the stack
			stack = append(stack, node) // Push the node to the stack

			// Perform depth-first search on the neighbors of the node
			for k, neighbor := range edges[node] {
				newLow := dfs(neighbor, timer+1+k) // Recursively call dfs on the neighbor
				if newLow < low[node] {            // Update the low value of the node if necessary
					low[node] = newLow
				}
			}

			// If the node is the head of a strongly connected component
			if dis[node] == low[node] {
				sccNodes := []int{}        // Initialize a slice to store the nodes in the SCC
				pop := stack[len(stack)-1] // Pop nodes from the stack until the current node is reached
				for pop != node {
					sccNodes = append(sccNodes, pop) // Add the popped node to the SCC
					inStack[pop] = false             // Mark the popped node as not in the stack
					stack = stack[:len(stack)-1]     // Remove the popped node from the stack
					pop = stack[len(stack)-1]        // Get the next node to be popped
				}
				sccNodes = append(sccNodes, pop) // Add the current node to the SCC
				inStack[pop] = false             // Mark the current node as not in the stack
				stack = stack[:len(stack)-1]     // Remove the current node from the stack
				scc = append(scc, sccNodes)      // Add the SCC to the list of SCCs
			}
			return low[node]
		}

		// If the node is in the stack, it is a back edge
		if inStack[node] {
			return dis[node]
		}

		// If the node is not in the stack, it is a cross edge
		return math.MaxInt
	}

	dfs(0, 0)        // Perform depth-first search starting from node 0
	fmt.Println(scc) // Print the strongly connected components
	return len(scc)
}
func main() {
	fmt.Println(tarzanAlgo(5, [][]int{{1, 0}, {0, 2}, {2, 1}, {0, 3}, {3, 4}}))         // [[4] [3] [1 2 0]] , 3
	fmt.Println(tarzanAlgo(6, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {4, 0}, {4, 5}})) // [[2] [5] [4 3 1 0]] , 3
}
