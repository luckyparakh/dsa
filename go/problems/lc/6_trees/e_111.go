// https://www.youtube.com/watch?v=tZS4VHtbYoo&t=496s
// https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
package main

type TreeNode struct {
	VAl   int
	Right *TreeNode
	Left  *TreeNode
}

// To level search or bfs and check for 1st leaf node encountered
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{}
	nodes = append(nodes, root)
	depth := 1
	for len(nodes) != 0 {
		l := len(nodes)
		for i := 0; i < l; i++ {
			node := nodes[0]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
			nodes = nodes[1:]
		}
		// Depth should be increased for each level not for each node, hence outside of above loop
		depth++
	}
	return depth
}

func main() {

}
