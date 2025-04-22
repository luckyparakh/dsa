package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	ans := [][]int{}
	for len(q) > 0 {
		level := []int{}
		l := len(q)
		for range l {
			node := q[0]
			level = append(level, node.Val)
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}
