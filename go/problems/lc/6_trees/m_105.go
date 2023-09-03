// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// https://www.youtube.com/watch?v=-YbXySKJsX8&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=37
// https://www.youtube.com/watch?v=wGmJatvjANY

package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 || len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := &TreeNode{Val: preorder[0]}

	position := findPosition(inorder, root.Val)
	root.Left = buildTree(preorder[1:position+1], inorder[:position])
	root.Right = buildTree(preorder[position+1:], inorder[position+1:])
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
func main() {

}
