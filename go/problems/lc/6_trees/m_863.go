package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parent := map[int]*TreeNode{}
	markParent(root, parent)
	// Map of Visited Node, to avoid duplicate
	visitedNodes := map[int]int{}
	s := []*TreeNode{}
	s = append(s, target)
	distance := 0
	for len(s) != 0 {
		if distance == k {
			break
		}
		l := len(s)
		// Level Wise
		for i := 0; i < l; i++ {
			firstNode := s[0]
			visitedNodes[firstNode.Val] = firstNode.Val

			// Add node in left
			if firstNode.Left != nil {
				// If node is already visited, then we should not move in that direction
				if _, found := visitedNodes[firstNode.Left.Val]; !found {
					s = append(s, firstNode.Left)
				}
			}

			// Add node in right
			if firstNode.Right != nil {
				if _, found := visitedNodes[firstNode.Right.Val]; !found {
					s = append(s, firstNode.Right)
				}
			}

			// Add Parent node as to move in upwards direction 
			if parent[firstNode.Val] != nil {
				if _, found := visitedNodes[parent[firstNode.Val].Val]; !found {
					s = append(s, parent[firstNode.Val])
				}
			}
			s = s[1:]
		}
		// Increase distance by 1 after each level
		distance++
	}
	ans := []int{}
	for _, v := range s {
		ans = append(ans, v.Val)
	}
	return ans
}

// markParent - Mark Parent of each node using a Map
func markParent(root *TreeNode, parent map[int]*TreeNode) {
	s := []*TreeNode{}
	s = append(s, root)
	for len(s) != 0 {
		n := s[0]
		if n.Left != nil {
			parent[n.Left.Val] = n
			s = append(s, n.Left)
		}
		if n.Right != nil {
			parent[n.Right.Val] = n
			s = append(s, n.Right)
		}
		s = s[1:]
	}
}
func main() {

}
