package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var helper func(*TreeNode) int
	d := 0
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lh := helper(root.Left)
		rh := helper(root.Right)
		d = max(d, rh+lh)
		return 1 + max(rh, lh)
	}
	helper(root)
	return d
}
