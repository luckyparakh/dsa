// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

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
	c := 0
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
		c++
		if c%2 == 1 {
			op = append(op, temp)
		} else {
			// Insert in reverse order
			abc := []int{}
			for i := len(temp) - 1; i >= 0; i-- {
				abc = append(abc, temp[i])
			}
			op = append(op, abc)
		}
	}
	return op
}
func main() {

}
