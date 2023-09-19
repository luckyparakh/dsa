// https://leetcode.com/problems/search-in-a-binary-search-tree/

package main

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	current := root
	for current != nil {
		cVal := current.Val
		if cVal == val {
			return current
		}
		current = current.Left
		if val > cVal {
			current = current.Right
		}
	}
	return nil
}
