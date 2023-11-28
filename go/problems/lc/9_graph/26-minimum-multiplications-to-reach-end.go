// https://practice.geeksforgeeks.org/problems/minimum-multiplications-to-reach-end/1
package main

import (
	"fmt"
	"math"
)

type Node struct {
	step  int
	value int
}

// minimumMultiplications calculates the minimum number of multiplications required to reach the end value
// from the start value using the given array of integers.
func minimumMultiplications(arr []int, start int, end int) int {
	// Define the maximum possible node value
	maxNode := 100000

	// Create an array to store the minimum number of steps required to reach each node
	steps := make([]int, maxNode)
	for i := range steps {
		// Set the initial number of steps to a very large value
		steps[i] = math.MaxInt
	}

	// Set the number of steps for the start node to 0
	steps[start] = 0

	// Create a queue to store the nodes to be processed
	q := make([]Node, 0)
	q = append(q, Node{0, start})

	// Process the nodes in the queue until it is empty
	for len(q) != 0 {
		// Get the first node from the queue
		node := q[0]

		// Check if the current node is equal to the end value
		if node.value == end {
			// Return the number of steps required to reach the end value
			return steps[node.value]
		}

		// Iterate over the array of integers
		for _, val := range arr {
			// Calculate the number of steps required to reach the new node
			step := steps[node.value] + 1

			// Calculate the new node value
			newVal := (node.value * val) % maxNode

			// Check if the new node has fewer steps than the current number of steps
			if steps[newVal] > step {
				// Update the number of steps for the new node
				steps[newVal] = step

				// Add the new node to the queue
				q = append(q, Node{step, newVal})
			}
		}

		// Remove the first node from the queue
		q = q[1:]
	}

	// Return -1 if the end value cannot be reached
	return -1
}

func main() {
	fmt.Println(minimumMultiplications([]int{2, 5, 7}, 3, 30))
	fmt.Println(minimumMultiplications([]int{3, 4, 65}, 7, 66175))
	fmt.Println(minimumMultiplications([]int{3, 19, 16, 17}, 7, 41711))
}
