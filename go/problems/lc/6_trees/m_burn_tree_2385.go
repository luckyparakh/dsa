// https://www.interviewbit.com/problems/burn-a-tree/
// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/description/

// Same as m_863
package main

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func solve(root *TreeNode, start int) int {
	parent := map[int]*TreeNode{}
	markParent(root, parent)
	nodeToStart := findNode(root, start)
	s := []*TreeNode{}
	visitedNodes := map[int]int{}
	burnTime := -1
	s = append(s, nodeToStart)
	for len(s) != 0 {
		l := len(s)
		for i := 0; i < l; i++ {
			tempNode := s[0]
			visitedNodes[tempNode.Val] = tempNode.Val

			if tempNode.Left != nil {
				if _, found := visitedNodes[tempNode.Left.Val]; !found {
					s = append(s, tempNode.Left)
				}
			}
			if tempNode.Right != nil {
				if _, found := visitedNodes[tempNode.Right.Val]; !found {
					s = append(s, tempNode.Right)
				}
			}
			parentNode := parent[tempNode.Val]
			if parentNode != nil {
				if _, found := visitedNodes[parentNode.Val]; !found {
					s = append(s, parentNode)
				}
			}
			s = s[1:]
		}
		burnTime++
	}
	return burnTime
}

func findNode(root *TreeNode, B int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == B {
		return root
	}

	lNode := findNode(root.Left, B)
	if lNode != nil {
		return lNode
	}
	rNode := findNode(root.Right, B)
	if rNode != nil {
		return rNode
	}
	return nil
}

func markParent(A *TreeNode, parent map[int]*TreeNode) {
	if A == nil {
		return
	}
	if A.Left != nil {
		parent[A.Left.Val] = A
		markParent(A.Left, parent)
	}
	if A.Right != nil {
		parent[A.Right.Val] = A
		markParent(A.Right, parent)
	}
}
