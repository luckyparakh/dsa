// https://leetcode.com/problems/delete-node-in-a-bst/
// https://www.youtube.com/watch?v=kouxiP_H5WE
package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	current, prev := search(root, key)
	if current == nil {
		return root
	}
	if prev == nil {
		// Removal of root as prev is nil
		if current.Left == nil && current.Right == nil {
			return nil
		} else if current.Left != nil && current.Right == nil {
			return current.Left
		} else if current.Left == nil && current.Right != nil {
			return current.Right
		} else {
			_, prev := search(current.Right, current.Left.Val)
			prev.Left = current.Left
			return current.Right
		}
	} else {
		if current.Left == nil && current.Right == nil {
			if current.Val > prev.Val {
				prev.Right = nil
			} else {
				prev.Left = nil
			}
		} else if current.Left != nil && current.Right == nil {
			if current.Val > prev.Val {
				prev.Right = current.Left
			} else {
				prev.Left = current.Left
			}
		} else if current.Left == nil && current.Right != nil {
			if current.Val > prev.Val {
				prev.Right = current.Right
			} else {
				prev.Left = current.Right
			}
		} else {
			if current.Val > prev.Val {
				prev.Right = current.Right
			} else {
				prev.Left = current.Right
			}
			_, prev := search(current.Right, current.Left.Val)
			prev.Left = current.Left
		}
	}
	return root
}

func search(root *TreeNode, val int) (*TreeNode, *TreeNode) {
	current := root
	var prev *TreeNode
	for current != nil {
		cVal := current.Val
		if val == cVal {
			return current, prev
		}
		prev = current
		if val > cVal {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return current, prev
}

func deleteNodeR(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key > root.Val {
		root.Right = deleteNodeR(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNodeR(root.Left, key)
	} else {
		if root.Left != nil && root.Right != nil {
			// Max node in a BST will the one with Right nil also it will be right most node
			maxNode := root.Left
			for maxNode.Right != nil {
				maxNode = maxNode.Right
			}
			maxNode.Right = root.Right
			return root.Left
		}
		if root.Left != nil {
			return root.Left
		}
		if root.Right != nil {
			return root.Right
		}
		return nil
	}
	return root
}
