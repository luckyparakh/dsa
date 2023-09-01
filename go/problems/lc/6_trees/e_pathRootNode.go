// https://www.interviewbit.com/problems/path-to-given-node/
// https://www.youtube.com/watch?v=fmflMqVOC7k&list=PLgUwDviBIf0q8Hkd7bK2Bpryj2xVJk8Vk&index=27

package main

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func solve(A *treeNode, B int) []int {
	if A == nil {
		return []int{}
	}
	if A.value == B {
		return []int{B}
	}
	if A.left != nil {
		leftArr := solve(A.left, B)
		if len(leftArr) != 0 {
			temp:=[]int{A.value}
			temp = append(temp, leftArr...)
			return temp
		}
	}
	if A.right != nil {
		rightArr := solve(A.right, B)
		if len(rightArr) != 0 {
			temp:=[]int{A.value}
			temp = append(temp, rightArr...)
			return temp
		}
	}
	return []int{}
}
func main() {

}
