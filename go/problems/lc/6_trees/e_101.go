// https://leetcode.com/problems/symmetric-tree/
// https://www.baeldung.com/cs/binary-tree-is-symmetric
// https://www.youtube.com/watch?v=nKggNAiEpBE&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=26

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	s := []*TreeNode{}
	s = append(s, root.Left, root.Right)
	for len(s) != 0 {
		root1 := s[0]
		root2 := s[1]
		if root1 == nil && root2 == nil {
			s = s[2:]
			continue
		} else if root1 == nil || root2 == nil {
			return false
		} else if root1.Val != root2.Val {
			return false
		} else {
			s = append(s, root1.Left, root2.Right)
			s = append(s, root1.Right, root2.Left)
		}
		s = s[2:]
	}
	return true
}
func isSymmetricR(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	} else if root1.Val != root2.Val {
		return false
	}
	return root1.Val == root2.Val && check(root1.Left, root2.Right) && check(root1.Right, root2.Left)
}

func main() {

}
