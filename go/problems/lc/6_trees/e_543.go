package main

// https://leetcode.com/problems/diameter-of-binary-tree/
// https://www.youtube.com/watch?v=Rezetez59Nk&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=18
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, d := height(root)
	return d
}
func height(root *TreeNode) (int, int) {
	m := -1
	if root == nil {
		return 0, m
	}
	lh, m1 := height(root.Left)
	rh, m2 := height(root.Right)
	if m1 > m2 {
		m = m1
	} else {
		m = m2
	}
	if lh+rh > m {
		m = lh + rh
	}
	return 1 + max(lh, rh), m
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}
