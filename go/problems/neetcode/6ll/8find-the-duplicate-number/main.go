package main

// Same finding loop in LL
// Assume index as node and value at index as pointer to next node
func findDuplicate(nums []int) int {
    slow,fast:=0,0

    for {
        slow=nums[slow] // Moving one node
        fast=nums[nums[fast]] // Moving 2 nodes
        if slow==fast{
            break
        }
    }
    slow=0
    for{
        slow=nums[slow]
        fast=nums[fast]
        if slow==fast{
            break
        }
    }
    return slow
}