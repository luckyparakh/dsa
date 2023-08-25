// https://leetcode.com/problems/lru-cache/submissions/
// https://www.youtube.com/watch?v=Xc4sICC8m4M&list=PLgUwDviBIf0r47RKH7fdWN54AbWFgGuii&index=18


// Create a double LL and maintain it in such way that least recent node lies at tail and recent used node lies at head
// It means insert node at head and delete node from tail which will help to achieve put with o(1)
// Maintain Map of nodes so that getting value is in o(1)

package main

type Node struct {
	Val  int
	Key  int
	Next *Node
	Prev *Node
}
type LinkedList struct {
	head *Node
	Tail *Node
}

type LRUCache struct {
	List     LinkedList
	Map      map[int]*Node
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*Node),
	}
}

// Get
// @receiver this *LRUCache
// @param key int
// @return int
func (this *LRUCache) Get(key int) int {
	node, found := this.Map[key]
	if !found {
		return -1
	}
	// If found, then delete and re-insert, this will make node as last used node
	this.Delete(node)
	this.Insert(node)
	return node.Val
}
func (this *LRUCache) Put(key int, value int) {
	node, found := this.Map[key]
	if !found {
		// Create node
		node = &Node{Key: key, Val: value}
		if len(this.Map) < this.Capacity {
			// if there is capacity insert node in both LL and Map
			this.Insert(node)
			this.Map[node.Key] = node
		} else {
			// In case there is no capacity, then delete tail node from both LL and Map.
			// Tail node is the least recent used node
			delete(this.Map, this.List.Tail.Key)
			this.Delete(this.List.Tail)

			// Insert node in LL and Map
			this.Insert(node)
			this.Map[node.Key] = node
		}
	} else {
		// If found then delete node from LL and map both
		this.Delete(node)
		delete(this.Map, node.Key)

		// Create new node, as value may change
		node = &Node{Key: key, Val: value}
		this.Insert(node)
		this.Map[node.Key] = node
	}
}
func (this *LRUCache) Insert(n *Node) {
	// Inserts nodes just after Head
	if this.List.head == nil {
		// If LL is empty then both tail and head will point to same node
		this.List.Tail = n
		this.List.head = n
	} else {
		this.List.head.Prev = n
		n.Next = this.List.head
		n.Prev = nil
		this.List.head = n
	}
}

func (this *LRUCache) Delete(n *Node) {
	// If next and prev is nil then it means there is only 1 node present
	if n.Next == nil && n.Prev == nil {
		this.List.head = nil
		this.List.Tail = nil
	}

	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		// Next nil means it is last node
		this.List.Tail = n.Prev
	}

	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		// Prev nil means it is 1st node
		this.List.head = n.Next
	}
}
