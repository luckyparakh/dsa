// https://leetcode.com/problems/binary-tree-level-order-traversal/description/

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	op := [][]int{}
	s := []*TreeNode{}
	s = append(s, root)
	for len(s) != 0 {
		temp := []int{}
		l := len(s)
		for i := 0; i < l; i++ {
			node := s[0]
			temp = append(temp, node.Val)
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
			s = s[1:]
		}
		op = append(op, temp)
	}
	return op
}
func main() {

}
