// https://leetcode.com/problems/binary-tree-right-side-view/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	// helper(root, &ans)
	level(root, &ans)
	return ans
}
// Do level traverse and append val in a temp slice
// And at end of level add last value of that temp slice in ans slice
func level(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	s := []*TreeNode{}
	s = append(s, root)
	for len(s) != 0 {
		t := []int{}
		for i := 0; i < len(s); i++ {
			node := s[0]
			t = append(t, node.Val)
			if node.Left != nil {
				s = append(s, node.Left)
			}
			if node.Right != nil {
				s = append(s, node.Right)
			}
			s = s[1:]
		}
		*ans = append(*ans, t[len(t)-1])
	}
}

// Below will fail if tree is [1,2,3,4]
func helper(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	*ans = append(*ans, root.Val)
	if root.Right != nil {
		helper(root.Right, ans)
	} else if root.Left != nil {
		helper(root.Left, ans)
	}
}
func main() {

}
