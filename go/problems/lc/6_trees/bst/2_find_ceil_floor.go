// https://www.youtube.com/watch?v=KSsk8AhdOZA&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=42
package main

import "math"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func findCeilBst(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	ceil := math.MaxInt
	current := root
	for current != nil {
		cVal := current.Val
		if cVal == val {
			return current
		}
		if val > cVal {
			current = current.Right
		} else {
			ceil = cVal
			current = current.Left
		}
	}
	return &TreeNode{Val: ceil}
}

func findFloorBst(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	floor := math.MinInt
	current := root
	for current != nil {
		cVal := current.Val
		if cVal == val {
			return current
		}
		if val > cVal {
			floor = cVal
			current = current.Right
		} else {
			current = current.Left
		}
	}
	return &TreeNode{Val: floor}
}
