// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// https://www.youtube.com/watch?v=9TJYWh0adfk&t=49s

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	values := []int{}
	var solve func(root *TreeNode)
	solve = func(root *TreeNode) {
		if root == nil {
			return
		}
		solve(root.Left)
		values = append(values, root.Val)
		solve(root.Right)
	}
	solve(root)
	return values[k-1]
}

func kthSmallestMT(root *TreeNode, k int) int {
	values := []int{}
	current := root
	for current != nil {
		if current.Left != nil {
			node := current.Left
			for node.Right != nil && node.Right != current {
				node = node.Right
			}
			if node.Right == nil {
				node.Right = current
				current = current.Left
			}else{
				values = append(values, current.Val)
				node.Right = nil
				current = current.Right
			}
		} else {
			values = append(values, current.Val)
			current = current.Right
		}
	}
	return values[k-1]
}
