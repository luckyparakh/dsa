// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// https://www.youtube.com/watch?v=WszrfSwMz58&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=18
package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func maxPathSum(root *TreeNode) int {
	m := -10000
	// Use pointer of m as m will passed across the functions
	findSum(root, &m)
	return m
}

func findSum(root *TreeNode, m *int) int {
	if root == nil {
		return 0
	}
	ls := findSum(root.Left, m)
	// Max is needed, hence ignore the -ve value node
	ls = max(0, ls)
	rs := findSum(root.Right, m)
	// Max is needed, hence ignore the -ve value node
	rs = max(0, rs)
	// Check for max path by adding values of lh,rh and val
	if ls+rs+root.Val > *m {
		*m = ls + rs + root.Val
	}
	// Return max
	return root.Val + max(ls, rs)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {

}
