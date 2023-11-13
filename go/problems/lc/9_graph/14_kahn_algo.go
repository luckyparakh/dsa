/*
1. Start by selecting a node that has no incoming edges (in-degree of 0).
If multiple nodes have no incoming edges, you can choose any of them.
2. Visit the selected node and add it to the sorted list.
3. Remove the selected node from the graph along with its outgoing edges.
4. Repeat steps 1-3 for the remaining nodes until all nodes have been visited.
5. If there are still nodes remaining in the graph after step 4, it means that the graph contains cycles and topological sorting is not possible.
*/
package main

import "fmt"

// topoSort performs a topological sort on a directed acyclic graph
// V is the number of vertices and adj is the adjacency list
func topoSort(V int, adj [][]int) []int {
	// inDegree stores the number of incoming edges for each vertex
	inDegree := make([]int, V)
	// ans stores the sorted order of vertices
	ans := []int{}
	// q is a queue used to perform the sort
	q := []int{}
	// calculate the inDegree for each vertex
	for i := 0; i < V; i++ {
		for _, node := range adj[i] {
			inDegree[node]++
		}
	}
	// add all vertices with inDegree 0 to the queue
	for node, degree := range inDegree {
		if degree == 0 {
			q = append(q, node)
		}
	}
	// perform the sort
	for len(q) != 0 {
		// get the next vertex with inDegree 0
		n := q[0]
		// add it to the sorted order
		ans = append(ans, n)
		// update the inDegree of its neighbors and add them to the queue if their inDegree becomes 0
		for _, neighbor := range adj[n] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
		// remove the processed vertex from the queue
		q = q[1:]
	}
	// return the sorted order of vertices
	return ans
}

func main() {
	fmt.Println(topoSort(6, [][]int{
		{}, {}, {3}, {1}, {0, 1}, {0, 2},
	}))


	fmt.Println(topoSort(6, [][]int{
		{5}, {}, {3}, {1}, {0, 1}, {0, 2},
	}))
}
