// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := pathToNode(root, p.Val)
	qPath := pathToNode(root, q.Val)
	c := 0
	for c = 0; c < len(pPath) && c < len(qPath); c++ {
		if pPath[c].Val != qPath[c].Val {
			break
		}
	}
	return pPath[c-1]
}

func pathToNode(A *TreeNode, B int) []*TreeNode {
	if A == nil {
		return []*TreeNode{}
	}
	if A.Val == B {
		return []*TreeNode{A}
	}
	if A.Left != nil {
		leftArr := pathToNode(A.Left, B)
		if len(leftArr) != 0 {
			temp := []*TreeNode{A}
			temp = append(temp, leftArr...)
			return temp
		}
	}
	if A.Right != nil {
		rightArr := pathToNode(A.Right, B)
		if len(rightArr) != 0 {
			temp := []*TreeNode{A}
			temp = append(temp, rightArr...)
			return temp
		}
	}
	return []*TreeNode{}
}
func main() {

}
