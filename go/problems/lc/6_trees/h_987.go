// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	ans := [][]int{}
	m := map[int][]int{}
	driver(root, m)

	// In map key are not ordered, hence to order Map do the following
	keys := []int{}
	// Take all keys from map
	for k := range m {
		keys = append(keys, k)
	}
	// Sort key in map
	sort.Ints(keys)
	// Again add in another slice
	for _, key := range keys {
		ans = append(ans, m[key])
	}
	fmt.Println(ans)
	return ans
}

type temp struct {
	node     *TreeNode
	colPosition int
}

func driver(root *TreeNode, ans map[int][]int) {
	s := []temp{}
	s = append(s, temp{root, 0})
	for len(s) != 0 {
		l := len(s)
		// Create temp map for each level
		nodesByColumn := map[int][]int{}
		for i := 0; i < l; i++ {
			t := s[0]
			if _, found := nodesByColumn[t.colPosition]; !found {
				nodesByColumn[t.colPosition] = []int{t.node.Val}
			} else {
				arr := nodesByColumn[t.colPosition]
				arr = append(arr, t.node.Val)
				nodesByColumn[t.colPosition] = arr
			}
			if t.node.Left != nil {
				s = append(s, temp{t.node.Left, t.colPosition - 1})
			}
			if t.node.Right != nil {
				s = append(s, temp{t.node.Right, t.colPosition + 1})
			}
			s = s[1:]
		}
		// fmt.Println(m1)
		// Sort nodes in a column and add in final answer
		for verticalPosition, nodes := range nodesByColumn {
			sort.Ints(nodes)
			if _, found := ans[verticalPosition]; !found {
				ans[verticalPosition] = nodes
			} else {
				// arr := ans[verticalPosition]
				// arr = append(arr, nodes...)
				ans[verticalPosition] = append(ans[verticalPosition], nodes...)
			}
		}
	}
}
func main() {

}
