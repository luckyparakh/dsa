// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// https://www.youtube.com/watch?v=NOKVBiJwkD0
// https://www.youtube.com/watch?v=sWf7k1x9XR4

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			node := cur.Left
			for node.Right != nil {
				node = node.Right
			}
			node.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}
