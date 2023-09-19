// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := []string{}
	stk := []*TreeNode{}
	stk = append(stk, root)
	for len(stk) != 0 {
		node := stk[0]
		if node != nil {
			ans = append(ans, strconv.Itoa(node.Val))
			stk = append(stk, node.Left)
			stk = append(stk, node.Right)
		} else {
			ans = append(ans, "null")
		}
		stk = stk[1:]
	}
	return strings.Join(ans, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	dataSlice := strings.Split(data, ",")
	stk := []*TreeNode{}
	root := createNode(dataSlice[0])
	stk = append(stk, root)
	position := 0
	for len(stk) != 0 {
		node := stk[0]
		position++
		leftNode := createNode(dataSlice[position])
		node.Left = leftNode
		if leftNode != nil {
			stk = append(stk, leftNode)
		}

		position++
		rightNode := createNode(dataSlice[position])
		node.Right = rightNode
		if rightNode != nil {
			stk = append(stk, rightNode)
		}

		stk = stk[1:]
	}
	return root
}

func createNode(val string) *TreeNode {
	if val == "null" {
		return nil
	}
	v, _ := strconv.Atoi(val)
	return &TreeNode{Val: v}
}
