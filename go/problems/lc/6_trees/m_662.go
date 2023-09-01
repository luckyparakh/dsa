// https://leetcode.com/problems/maximum-width-of-binary-tree/

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}
type temp struct {
	n  *TreeNode
	cp int
}

func widthOfBinaryTree(root *TreeNode) int {
	s := []temp{}
	maxW := -1
	s = append(s, temp{root, 0})
	level, multiplier := 0, 0
	for len(s) != 0 {
		l := len(s)
		lc := s[0].cp
		rc := s[len(s)-1].cp
		if rc-lc+1 > maxW {
			maxW = rc - lc + 1
		}
		for i := 0; i < l; i++ {
			node := s[0]
			if level == 0 {
				multiplier = 2 * node.cp
			} else {
				// This needed, suppose if tree is skew and it has 3000 nodes than 
				// multiplier for would be 2^3000 (not exact) hence it would be a big number,
				// and it may overflow, hence using below formula
				// With help of this, node on next level will always start with 1 
				multiplier = 2 * (node.cp - 1)
			}
			if node.n.Left != nil {
				s = append(s, temp{node.n.Left, multiplier + 1})
			}
			if node.n.Right != nil {
				s = append(s, temp{node.n.Right, multiplier + 2})
			}
			s = s[1:]
		}
		level++
	}
	return maxW
}
func main() {

}
