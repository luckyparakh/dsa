// https://leetcode.com/problems/course-schedule/

package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj = append(adj, []int{})
	}
	for _, v := range prerequisites {
		adj[v[0]] = append(adj[v[0]], v[1])
	}
	return hasCycle(numCourses, adj)
}

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

// hasCycle checks if a graph represented by the adjacency matrix has a cycle
func hasCycle(V int, adj [][]int) bool {
	// Check if the number of vertices is equal to the length of the topological sort
	if len(topoSort(V, adj)) == V {
		return false // No cycle found
	}
	return true // Cycle found
}
