// https://www.hackerrank.com/challenges/tree-top-view/problem

package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}
type temp struct {
	node        *TreeNode
	colPosition int
}

func topView(root *TreeNode) {
	m := map[int]int{}
	driver(root, m)
	fmt.Println(m)
	keys := []int{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	ans := []int{}
	for _, v := range keys {
		ans = append(ans, m[v])
	}
	fmt.Println(ans)
}

func driver(root *TreeNode, ans map[int]int) {
	s := []temp{}
	s = append(s, temp{root, 0})
	for len(s) != 0 {
		l := len(s)
		for i := 0; i < l; i++ {
			tNode := s[0]
			if _, found := ans[tNode.colPosition]; !found {
				ans[tNode.colPosition] = tNode.node.Val
			}
			if tNode.node.Left != nil {
				s = append(s, temp{tNode.node.Left, tNode.colPosition - 1})
			}
			if tNode.node.Right != nil {
				s = append(s, temp{tNode.node.Right, tNode.colPosition + 1})
			}
			s = s[1:]
		}
	}
}
func main() {

}
