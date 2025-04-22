package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r, c := 0, 0
	var prev *ListNode
	var ans *ListNode
	for l1 != nil && l2 != nil {
		s := l1.Val + l2.Val + r
		r = s / 10
		n := ListNode{Val: s % 10}
		if c == 0 {
			ans = &n
		} else {
			// fmt.Println(c,prev,n)
			prev.Next = &n
			// fmt.Println(c,prev,n)
		}
		// fmt.Println(ans)
		prev = &n
		l1 = l1.Next
		l2 = l2.Next
		c++
	}
	for l1 != nil {
		s := l1.Val + r
		r = s / 10
		n := ListNode{Val: s % 10}
		prev.Next = &n
		prev = &n
		l1 = l1.Next
	}
	for l2 != nil {
		s := l2.Val + r
		r = s / 10
		n := ListNode{Val: s % 10}
		prev.Next = &n
		prev = &n
		l2 = l2.Next
	}
	if r != 0 {
		n := ListNode{Val: r}
		prev.Next = &n
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}
