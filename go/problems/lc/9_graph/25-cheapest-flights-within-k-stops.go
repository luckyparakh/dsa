// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// https://www.youtube.com/watch?v=9XybHVqTHcQ&list=PLgUwDviBIf0oE3gA41TKO2H5bHpPd7fzn&index=38

package main

import (
	"fmt"
	"math"
	"sort"
)

type Node struct {
	number int
	cost   int
	stop   int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := -1
	// Create adjacency matrix
	adj := make([][][]int, n)
	for i := range adj {
		adj[i] = make([][]int, 0)
	}
	for _, flight := range flights {
		// Add edge to adjacency matrix
		adj[flight[0]] = append(adj[flight[0]], []int{flight[1], flight[2]})
	}

	fmt.Println(adj)
	costs := make([]int, n)
	for k := range costs {
		costs[k] = math.MaxInt32
	}

	// Define the start and end nodes
	startNode := Node{
		number: src,
	}

	// Create a queue to store nodes
	q := []Node{}

	q = append(q, startNode)
	costs[startNode.number] = 0
	for len(q) != 0 {
		sort.SliceStable(q, func(i, j int) bool {
			return q[i].cost < q[j].cost
		})

		// Pop a node from the queue
		node := q[0]
		if node.number == dst && node.stop-1 <= k {
			return node.cost
		}
		for _, neighbor := range adj[node.number] {
			cost := node.cost + neighbor[1]
			if costs[neighbor[0]] > cost {
				costs[neighbor[0]] = cost
				node:=Node{
					number: neighbor[0],
					cost:   cost,
					stop:   node.stop + 1,
				}
				fmt.Println(node)
				q = append(q, node)
			}
		}
		q = q[1:]

	}
	fmt.Println(costs)
	return cost
}

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 1}, {0, 2, 5}, {1, 2, 1}, {2, 3, 1}}, 0, 3, 1))
}
