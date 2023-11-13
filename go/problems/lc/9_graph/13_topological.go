// https://practice.geeksforgeeks.org/problems/topological-sort/1
// https://www.youtube.com/watch?v=5lZ0iJMrUMk&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=21

/*
Topological sorting is a technique used to order the nodes in a directed acyclic graph (DAG) based on their dependencies. 
It provides a linear ordering of the nodes such that for every directed edge (u, v), 
node u comes before node v in the ordering.

In a directed acyclic graph, the edges have a specific direction and there are no cycles, 
meaning it is not possible to traverse the graph and return to a previously visited node. 
This acyclic property is crucial for topological sorting to be meaningful and well-defined.
 
Topological sorting is commonly used in scenarios where there are dependencies between tasks or processes. 
It helps determine the order in which these tasks should be executed to satisfy the dependencies.
*/
package main

import "fmt"


// topoSort performs a topological sort on a directed acyclic graph
func topoSort(V int, adj [][]int) []int {
	// nodes is the number of nodes in the graph
	nodes := len(adj)
	// visitedNodes keeps track of which nodes have been visited during the DFS
	visitedNodes := make([]bool, nodes)
	// temp holds the nodes in the order they were visited during the DFS
	temp := make([]int, 0, nodes)
	// dfs performs a depth-first search on the graph starting at node
	// and adds each visited node to temp
	var dfs func(node int)
	dfs = func(node int) {
		visitedNodes[node] = true
		for _, neighbor := range adj[node] {
			if !visitedNodes[neighbor] {
				dfs(neighbor)
			}
		}
		temp = append(temp, node)
	}
	// Perform DFS on each unvisited node
	for node := range adj {
		if !visitedNodes[node] {
			dfs(node)
		}
	}
	// Reverse the order of temp to get the topological sort
	ans := []int{}
	for len(temp) != 0 {
		lastNodeIndex := len(temp) - 1
		ans = append(ans, temp[lastNodeIndex])
		temp = temp[:lastNodeIndex]
	}
	return ans
}

func main() {
	fmt.Println(topoSort(6, [][]int{
		{}, {}, {3}, {1}, {0, 1}, {0, 2},
	}))
}
