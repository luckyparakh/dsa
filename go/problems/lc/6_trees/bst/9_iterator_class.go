// https://www.youtube.com/watch?v=D2jMcmxU4bs
// https://leetcode.com/problems/binary-search-tree-iterator/
package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}
type BSTIterator struct {
	stack []*TreeNode
}

// Another approach can be that in constructor itself create a slice which has inorder of tree
// But due to this s(O) will be n as in that slice all n nodes will be stored.

// Hence this approach where stack is used and this stack will store (max) all left nodes of a specific node 
// Hence in worst case s(O) will be O(h) where h is hg of tree.
func Constructor(root *TreeNode) BSTIterator {
	s := []*TreeNode{}
	for root != nil {
		s = append(s, root)
		root = root.Left
	}
	return BSTIterator{
		stack: s,
	}
}

func (this *BSTIterator) Next() int {
	nodeToPop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if nodeToPop.Right != nil {
		temp := nodeToPop.Right
		for temp != nil {
			this.stack = append(this.stack, temp)
			temp = temp.Left
		}
	}

	return nodeToPop.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}
