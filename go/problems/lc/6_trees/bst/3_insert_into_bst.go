// https://leetcode.com/problems/insert-into-a-binary-search-tree/
// https://www.youtube.com/watch?v=FiFiNvM29ps&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=44

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}
	current := root
	var prev *TreeNode
	for current != nil {
		cVal := current.Val
		prev = current
		if val > cVal {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	if val > prev.Val {
		prev.Right = node
	} else {
		prev.Left = node
	}
	return root
}

func insertIntoBSTR(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBSTR(root.Right, val)
	} else {
		root.Left = insertIntoBSTR(root.Left, val)
	}
	return root
}
