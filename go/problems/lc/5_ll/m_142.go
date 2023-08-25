// https://leetcode.com/problems/linked-list-cycle-ii/description/
// https://www.youtube.com/watch?v=QfbOhn0WZ88&list=PLgUwDviBIf0r47RKH7fdWN54AbWFgGuii&index=11

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	sp, fp := head, head
	for fp != nil && fp.Next != nil {
		if sp == fp {
			break
		}
		sp = sp.Next
		fp = fp.Next.Next
	}
	if fp == nil || fp.Next == nil {
		return nil
	}
	
	oneMorePointer := head
	for sp != oneMorePointer {
		sp = sp.Next
		oneMorePointer = oneMorePointer.Next
	}
	return sp
}

func main() {

}
