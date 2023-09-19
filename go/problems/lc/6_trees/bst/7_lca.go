// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == q.Val || root.Val == p.Val {
		return root
	} else if root.Val > p.Val && root.Val < q.Val {
		return root
	} else if root.Val < p.Val && root.Val > q.Val {
		return root
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return nil
}
