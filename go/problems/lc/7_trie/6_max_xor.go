// https://www.youtube.com/watch?v=jCuZCbXnpLo&list=PLEJXowNB4kPyi859E6qGUs7jlpQehJndl&index=8&t=569s
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/solutions/3055362/concise-golang-solution/
package main

import "fmt"

// Node represents a node in the Trie
type Node struct {
	// Don't Use map as it will give error of out of memory for large input
	// children map[int]*Node
	
	children [2]*Node
}

// Trie represents the Trie data structure
type Trie struct {
	root *Node
}

/*
This function is used to insert a number into a Trie data structure. The parameter  `n`  represents the number to be inserted. The function starts by initializing a variable  `tNode`  to the root of the Trie.
The function then iterates from the most significant bit (31) to the least significant bit (0) of the number  `n` . It extracts each bit of  `n`  using the right shift operator and bitwise AND operation.
If the child node corresponding to the extracted bit doesn't exist, a new child node is created and assigned to  `tNode.children[bit]` .
Finally, the function moves to the child node by updating  `tNode`  to  `tNode.children[bit]` . This process is repeated for each bit of  `n` , effectively inserting the number into the Trie.
*/
func (t *Trie) insert(n int) {
	tNode := t.root
	for i := 31; i >= 0; i-- {
		bit := n >> i & 1
		// If the child node doesn't exist, create a new one
		if tNode.children[bit] == nil {
			tNode.children[bit] = &Node{}
		}
		// Move to the child node
		tNode = tNode.children[bit]
	}
}

/*
The function  `find`  is a method of the  `Trie`  struct. It takes an integer  `n`  as input and returns an integer.
The function starts by initializing a variable  `tNode`  with the root node of the trie and a variable  `val`  with 0.
It then iterates from the 31st bit to the 0th bit of the input integer  `n` .
Inside the loop, it extracts the  `i` th bit of  `n`  using the right shift and bitwise AND operations.
It then calculates the XOR of the extracted bit and 1, and assigns it to the variable  `Xor` .
If the XOR child node exists in the trie for the current bit, it updates the  `val`  by setting the  `i` th bit to  `Xor`  and moves the  `tNode`  to the XOR child node.
If the XOR child node does not exist, it sets the  `val`  by setting the  `i` th bit to the extracted bit and moves the  `tNode`  to the bit child node.
After the loop, it returns the final value  `val` .
This function is used to find the maximum XOR value for a given integer  `n`  in the trie.
*/
func (t *Trie) find(n int) int {
	tNode := t.root
	val := 0
	for i := 31; i >= 0; i-- {
		bit := n >> i & 1
		Xor := bit ^ 1
		// If the XOR child node exists, update the value and move to that node
		if tNode.children[Xor] != nil {
			val |= Xor << i
			tNode = tNode.children[Xor]
		} else {
			// Otherwise, move to the bit child node
			val |= bit << i
			tNode = tNode.children[bit]
		}
	}
	return val
}

// findMaximumXOR finds the maximum XOR value among all pairs of numbers in the given array
func findMaximumXOR(nums []int) int {
	ans := 0
	t := &Trie{root: &Node{}}
	for _, num := range nums {
		t.insert(num)
	}
	for _, num := range nums {
		val := t.find(num) ^ num
		// Update the answer if the XOR value is greater
		if val > ans {
			ans = val
		}
	}
	return ans
}

func main() {
	// Test the findMaximumXOR function
	fmt.Println(findMaximumXOR([]int{8, 5, 4, 9, 1}))
}
