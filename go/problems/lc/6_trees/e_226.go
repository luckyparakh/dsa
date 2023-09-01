// https://leetcode.com/problems/invert-binary-tree/

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	helper(root)
	return root
}

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return
}

func main() {

}
