package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// Find mid of list
	slow, fast := head, head
	// In LL sometime it very easy to find base condition, check to safeguard against nil pointer like fast can't be nil else fast.next will give nil pointer error
	// Similar fast.next can't be nil, else fast.next.next will give nil pointer
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	// This needed, as we are braking LL in 2 part, so 1st LL end should be nil
	slow.Next = nil

	// Reverse second
	var second *ListNode
	for mid != nil {
		nn := mid.Next
		mid.Next = second
		second = mid
		mid = nn
	}

	// merge
	first := head
	// Checking of only second nil, if len(original ll) is odd then 2nd part will be one element less than 1st part.
	// If len(original ll) is even then equal than both part will be equal in that case check nil of any 1st or 2nd part will work.
	// Hence to server both the case (odd or even len), then checking 2nd part as nil is suff. for both cases
	for second != nil {
		fn := first.Next
		sn := second.Next
		first.Next = second
		second.Next = fn
		first = fn
		second = sn
	}

}
