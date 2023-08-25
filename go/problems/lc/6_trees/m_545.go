// https://leetcode.com/problems/boundary-of-binary-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func boundaryTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	fmt.Println("Left")
	findLeft(root.Left)
	fmt.Println("Leaves")
	findLeaves(root)
	fmt.Println("Right")
	findRight(root.Right)
	return
}
func findLeft(r *TreeNode) {
	if r == nil {
		return
	}
	if r.Left != nil || r.Right != nil {
		fmt.Println(r.Val)
	}

	if r.Left != nil {
		findLeft(r.Left)
	} else if r.Right != nil {
		findLeft(r.Right)
	}
}

func findLeaves(r *TreeNode) {
	if r == nil {
		return
	}
	findLeaves(r.Left)
	if r.Left == nil && r.Right == nil {
		fmt.Println(r.Val)
	}
	findLeaves(r.Right)
}
func findRight(r *TreeNode) {
	// Print while returning
	if r == nil {
		return
	}
	if r.Right != nil {
		findRight(r.Right)
	} else if r.Left != nil {
		findRight(r.Left)
	}
	if r.Left != nil || r.Right != nil {
		fmt.Println(r.Val)
	}
}
func main() {

}
