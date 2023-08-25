// https://leetcode.com/problems/balanced-binary-tree/

package main

type TreeNode struct {
	VAl   int
	Right *TreeNode
	Left  *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if height(root) == -1 {
		return false
	}
	return true
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	if lh == -1 {
		return -1
	}
	rh := height(root.Right)
	if rh == -1 {
		return -1
	}
	if abs(rh-lh) > 1 {
		return -1
	}
	return 1 + max(lh, rh)
}
func abs(z int) int {
	if z < 0 {
		return -z
	}
	return z
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
