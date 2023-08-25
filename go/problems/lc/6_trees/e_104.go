package main

// https://leetcode.com/problems/maximum-depth-of-binary-tree/

import "math"

type TreeNode struct {
	VAl   int
	Right *TreeNode
	Left  *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl := maxDepth(root.Left)
	rl := maxDepth(root.Right)
	return 1 + int(math.Max(float64(hl), float64(rl)))
}
