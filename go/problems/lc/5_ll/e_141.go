// https://www.youtube.com/watch?v=354J83hX7RI&list=PLgUwDviBIf0r47RKH7fdWN54AbWFgGuii&index=9
// https://leetcode.com/problems/linked-list-cycle/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type LinkedList struct {
	head *ListNode
}

func hasCycle(head *ListNode) bool {
	sp, fp := head, head

	for sp != nil && fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
		if sp == fp {
			return true
		}
	}
	return false
}
func main() {

}
