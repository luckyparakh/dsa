package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	c := 1
	return helper(root, c)
}

func helper(root *TreeNode, c int) int {
	if root.Left == nil && root.Right == nil {
		return c
	}
	cl, cr := 0, 0
	if root.Left != nil {
		cl = helper(root.Left, c+1)
	}
	if root.Right != nil {
		cr = helper(root.Right, c+1)
	}
	return max(cr, cl)
}
