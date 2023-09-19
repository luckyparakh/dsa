// https://leetcode.com/problems/construct-binary-search-tree-from-preorder-traversal/
// https://www.youtube.com/watch?v=UmJT3j26t1I

package main

import "math"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// Approach 1
// Read node one by one from slice and insert in tree

// Approach 2
// 
func bstFromPreorder(preorder []int) *TreeNode {
	start := 0
	return helper(preorder, math.MaxInt, &start)
}

func helper(preorder []int, max int, index *int) *TreeNode {
	if *index > len(preorder)-1 || preorder[*index] > max {
		return nil
	}
	root := &TreeNode{Val: preorder[*index]}
	*index = *index + 1
	root.Left = helper(preorder, root.Val, index)
	root.Right = helper(preorder, max, index)
	return root
}
