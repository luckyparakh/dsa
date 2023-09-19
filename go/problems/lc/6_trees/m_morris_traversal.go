// https://www.youtube.com/watch?v=80Zug6D1_r4&t=600s
package main

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func morrisTraversalInOrder(root *TreeNode) []int {
	// If input root is nil return nil
	if root == nil {
		return nil
	}
	ans := []int{}
	var solve func(root *TreeNode)
	solve = func(current *TreeNode) {
		// if root comes nil then return and this would also be base condition and also nil pointer at line 22
		if current == nil {
			return
		}
		// If root left is nil, it means this root hence append to slice 
		// And move to right
		if current.Left == nil {
			ans = append(ans, current.Val)
			solve(current.Right)
		} else {
			// Else find the in order predecessor of current predecessor
			// means find right most predecessor in left side of tree 
			// that is the predecessor that the predecessor of current predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			// 1st time predecessor right will be nil, so we will set it right to current and move left
			if predecessor.Right == nil {
				predecessor.Right = current
				solve(current.Left)
			} else {
				// It means predecessor right is already pointing to current, which also means that left part of current has
				// visited, hence move to right
				// Also add val in slice since this is a root node
				ans = append(ans, current.Val)
				predecessor.Right = nil
				solve(current.Right)
			}
		}
	}
	solve(root)
	return ans
}

func morrisTraversalPreOrder(root *TreeNode) []int {
	// If input root is nil return nil
	if root == nil {
		return nil
	}
	ans := []int{}
	var solve func(root *TreeNode)
	solve = func(current *TreeNode) {
		// if root comes nil then return and this would also be base condition and also nil pointer at line 22
		if current == nil {
			return
		}
		// If root left is nil, it means this root hence append to slice 
		// And move to right
		if current.Left == nil {
			ans = append(ans, current.Val)
			solve(current.Right)
		} else {
			// Else find the in order predecessor of current predecessor
			// means find right most predecessor in left side of tree 
			// that is the predecessor that the predecessor of current predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			// 1st time predecessor right will be nil, so we will set it right to current and move left
			if predecessor.Right == nil {
				// Also add val in slice because when pred is pointing to cur 
				// at that point cur is root
				ans = append(ans, current.Val)
				predecessor.Right = current
				solve(current.Left)
			} else {
				// It means predecessor right is already pointing to current, which also means that left part of current has
				// visited, hence move to right
				predecessor.Right = nil
				solve(current.Right)
			}
		}
	}
	solve(root)
	return ans
}
