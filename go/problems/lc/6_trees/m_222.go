// https://www.youtube.com/watch?v=CvrPf1-flAA
// https://leetcode.com/problems/count-complete-tree-nodes/
package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := findHgL(root.Left)
	rh := findHgR(root.Right)
	fmt.Println(lh, rh, root.Val)
	if lh == rh {
		return int(math.Pow(2, float64(lh+1))) - 1
	}
	return 1 + countNodes(root.Right) + countNodes(root.Left)
}
func findHgL(root *TreeNode) int {
	if root != nil {
		return 1 + findHgL(root.Left)
	}
	return 0
}
func findHgR(root *TreeNode) int {
	if root != nil {
		return 1 + findHgR(root.Right)
	}
	return 0
}
