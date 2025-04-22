package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	answer := true
	var helper func(*TreeNode) int
	helper = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		lh, rh := 0, 0
		if answer {
			lh = helper(tn.Left)
			rh = helper(tn.Right)
			if abs(rh, lh) > 1 {
				answer = false
			}
		}
		return 1 + max(rh, lh)
	}
	helper(root)
	return answer
}
func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
