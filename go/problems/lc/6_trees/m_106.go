// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
// https://www.youtube.com/watch?v=-YbXySKJsX8&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=37
// https://www.youtube.com/watch?v=wGmJatvjANY

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(postorder) == 1 || len(inorder) == 1 {
		return &TreeNode{Val: postorder[len(postorder)-1]}
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	position := findPosition(inorder, root.Val)
	root.Left = buildTree(inorder[:position], postorder[:position])
	root.Right = buildTree(inorder[position+1:], postorder[position:len(postorder)-1])
	return root
}
func findPosition(slice []int, value int) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
