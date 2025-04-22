package main
// Get is like removeNode (if present) then add node to head
// Put is like updateNode (if present) and move to head. if not present then add node to head and remove tail node if size exceeds capacity.
type Node struct {
	value, key int
	next, prev *Node
}
type LRUCache struct {
	cache    map[int]*Node
	size     int
	capacity int
	tail     *Node
	head     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache:    map[int]*Node{},
		size:     0,
		capacity: capacity,
		tail:     nil,
		head:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) moveToHead(node *Node) {
	if node == this.head {
		return
	}
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *Node) {
	// think 2 cases
	// if this is 1st node
	// if this is not 1st node
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		// this is 1st node
		this.tail = node
	}
}

func (this *LRUCache) removeNode(node *Node) {
	// think 3 cases:
	// Ist when node is middle node ie both prev/next are not nil
	// 2nd when node's both next and prev is nil, which means its last node of list
	// 3rd when node's next is nil, which means it is tail node
	// Here in LRU we never remove head node (ie when prev is nil), as head node is most recent node and we always preserve it.
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// if node is head node
		this.head = node.next
		// Although this not needed in LRU case but for sake of completeness kept here, to keep it generic double ll operation
		if this.head != nil {
			// if node was not last node
			this.head.prev = nil
		}
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// if node is tail node
		this.tail = node.prev
		if this.tail != nil {
			// if node was not last node
			this.tail.next = nil
		}
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value // update to new value
		this.moveToHead(node)
	} else {
		nn := &Node{
			value: value,
			key:   key,
			next:  nil,
			prev:  nil,
		}
		this.cache[key] = nn
		// think if it can be 1st node being added then in that case update both tail/head
		this.addToHead(nn)
		this.size++
		if this.size > this.capacity {
			rm := this.cache[this.tail.key]
			delete(this.cache, rm.key)
			// this.tail = rm.prev
			// this.tail.next = nil
			this.removeNode(rm)
			this.size--
		}
	}
}
