// https://leetcode.com/problems/validate-binary-search-tree/

package main

import "math"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isValidBSTInOrder(root *TreeNode) bool {
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
	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	return true
}
func isValidBSTRange(root *TreeNode) bool {
	// https://www.youtube.com/watch?v=f-sj7I5oXEI
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, min, max int) bool {
	if root.Val > min && root.Val < max {
		if root.Left != nil {
			op := helper(root, min, root.Val)
			if !op {
				return false
			}
		}
		if root.Right != nil {
			op := helper(root, root.Val, max)
			if !op {
				return false
			}
		}
		return true
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	} else if root.Left != nil && root.Right != nil {
		rl := checkNodesS(root.Left, root.Val)
		rr := checkNodesL(root.Right, root.Val)
		if rl && rr {
			return isValidBST(root.Left) && isValidBST(root.Right)
		} else {
			return false
		}
	} else if root.Left != nil {
		rl := checkNodesS(root.Left, root.Val)
		if rl {
			return isValidBST(root.Left)
		}
		return false
	} else {
		rr := checkNodesL(root.Right, root.Val)
		if rr {
			return isValidBST(root.Right)
		}
		return false
	}
}

func checkNodesS(root *TreeNode, val int) bool {
	if root.Val < val {
		if root.Left != nil {
			op := checkNodesS(root.Left, val)
			if !op {
				return op
			}
		}
		if root.Right != nil {
			op := checkNodesS(root.Right, val)
			if !op {
				return op
			}
		}
		return true
	}
	return false
}

func checkNodesL(root *TreeNode, val int) bool {
	if root.Val > val {
		if root.Left != nil {
			op := checkNodesL(root.Left, val)
			if !op {
				return op
			}
		}
		if root.Right != nil {
			op := checkNodesL(root.Right, val)
			if !op {
				return op
			}
		}
		return true
	}
	return false
}
