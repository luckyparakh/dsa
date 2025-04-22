package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if min(p.Val, q.Val) > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
