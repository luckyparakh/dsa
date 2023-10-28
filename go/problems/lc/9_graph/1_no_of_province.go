// https://leetcode.com/problems/number-of-provinces/
// https://www.youtube.com/watch?v=ACzkVtewUYA&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=7

package main
/*
This code defines a function  `findCircleNum`  that takes a 2D matrix  `isConnected`  as input. It uses depth-first search (DFS) to find the number of connected circles in the matrix. Each circle represents a group of nodes that are connected to each other. The function initializes a counter  `c`  to 0 and a visited nodes map  `vn` . It then defines a recursive DFS function  `dfs`  that explores the neighbors of a given node and marks them as visited. The main function iterates through all nodes in the matrix, calling the DFS function for unvisited nodes and incrementing the counter. Finally, it returns the total number of circles found.
*/
func findCircleNum(isConnected [][]int) int {
	c := 0
	vn := map[int]bool{}
	var dfs func(i int)
	dfs = func(node int) {
		for neighbor, v := range isConnected[node] {
			if v == 1 && neighbor != node && !vn[neighbor] {
				vn[neighbor] = true
				dfs(neighbor)
			}
		}
	}
	for node := range isConnected {
		if !vn[node] {
			vn[node] = true
			c++
			dfs(node)
		}
	}
	return c
}
