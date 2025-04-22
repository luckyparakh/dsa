package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	nodesVal := []int{}
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		goodNode := true
		for _, v := range nodesVal {
			if v > tn.Val {
				goodNode = false
				break
			}
		}
		if goodNode {
			ans++
		}
		nodesVal = append(nodesVal, tn.Val)
		if tn.Left != nil {
			dfs(tn.Left)
		}
		if tn.Right != nil {
			dfs(tn.Right)
		}
		nodesVal = nodesVal[:len(nodesVal)-1]
	}
	dfs(root)
	return ans
}

func goodNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, maxVal int) {
		if tn.Val >= maxVal {
			maxVal = tn.Val
			ans++
		}
		if tn.Left != nil {
			dfs(tn.Left, maxVal)
		}
		if tn.Right != nil {
			dfs(tn.Right, maxVal)
		}
	}
	dfs(root, math.MinInt)
	return ans
}
