// https://leetcode.com/problems/is-graph-bipartite/
// https://en.wikipedia.org/wiki/Bipartite_graph
// https://www.youtube.com/watch?v=-vu34sct1g8&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=17

// This program checks whether a given graph is bipartite or not.
// A bipartite graph is a graph in which the vertices can be divided into two disjoint sets such that every edge
// connects two vertices from different sets.
// The program uses BFS to traverse the graph and assign colors to the vertices.
// If a vertex is colored with two different colors, it means the graph is not bipartite.
package main

import "fmt"

// isBipartite checks whether a given graph is bipartite or not.
func isBipartite(graph [][]int) bool {
	// r is the number of vertices in the graph
	r := len(graph)
	// color is an array to store the color of each vertex
	// 0 means no color, -1 and 1 are two colors
	color := make([]int, r)
	// loop through all the vertices in the graph
	for k, rootNode := range graph {
		// if the vertex is not colored and has neighbors
		if len(rootNode) > 0 && color[k] == 0 {
			// assign the vertex a color and check if the graph is bipartite
			color[k] = 1
			if !setColor(k, &color, graph) {
				return false
			}
		}
	}
	return true
}

// setColor assigns colors to the vertices of the graph using BFS
func setColor(k int, c *[]int, graph [][]int) bool {
	color := (*c)
	q := []int{}
	q = append(q, k)
	for len(q) != 0 {
		// pop the first vertex from the queue
		node := q[0]
		// assign the opposite color to the neighbors of the vertex
		// if the vertex is colored with -1, the neighbors are colored with 1
		// if the vertex is colored with 1, the neighbors are colored with -1
		neighColor := -1
		if color[node] == -1 {
			neighColor = 1
		}
		for _, neighbor := range graph[node] {
			// if the neighbor is not colored, assign it a color and add it to the queue
			if color[neighbor] == 0 {
				color[neighbor] = neighColor
				q = append(q, neighbor)
			} else { // if the neighbor is already colored, check if the colors are different
				if color[neighbor] != neighColor {
					return false
				}
			}
		}
		// remove the vertex from the queue
		q = q[1:]
	}
	return true
}

func main() {
	// true
	fmt.Println(isBipartite([][]int{
		{1, 3}, {0, 2}, {1, 3}, {0, 2},
	}))
	// false
	fmt.Println(isBipartite([][]int{
		{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9},
		{2, 4, 5, 6, 7, 8},
	}))
	// false
	fmt.Println(isBipartite([][]int{
		{2, 4}, {2, 3, 4}, {0, 1}, {1}, {0, 1}, {7}, {9}, {5}, {}, {6}, {12, 14}, {}, {10},
		{}, {10}, {19}, {18}, {}, {16}, {15}, {23}, {23}, {}, {20, 21}, {}, {}, {27}, {26},
		{}, {}, {34}, {33, 34}, {}, {31}, {30, 31}, {38, 39}, {37, 38, 39}, {36}, {35, 36}, {35, 36}, {43},
		{}, {}, {40}, {}, {49}, {47, 48, 49}, {46, 48, 49}, {46, 47, 49}, {45, 46, 47, 48},
	}))
}
