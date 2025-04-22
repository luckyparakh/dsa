package main

// better way
type MinStack struct {
	items    []int
	minItems []int
	min      int
}

func Constructor() MinStack {
	return MinStack{
		items:    []int{},
		minItems: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.items) == 0 {
		this.min = val
	}
	if val < this.min {
		this.min = val
	}
	this.items = append(this.items, val)
	this.minItems = append(this.minItems, this.min)
	// fmt.Println(this.items,this.minItems)
}

func (this *MinStack) Pop() {
	l := len(this.items)
	isMinGettingOut := false
	if this.min == this.GetMin() {
		isMinGettingOut = true
	}
	this.items = this.items[:l-1]
	this.minItems = this.minItems[:l-1]
	// if len of item is zero then it will fail or give error
	if isMinGettingOut && len(this.items) > 0 {
		this.min = this.GetMin()
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.minItems[len(this.items)-1]
}

// // One Way (not so good)
// type MinStack struct {
// 	items []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		items: []int{},
// 	}
// }

// func (this *MinStack) Push(val int) {
// 	this.items = append(this.items, val)
// 	this.Up(len(this.items) - 1)
// }
// func (this *MinStack) Up(index int) {
// 	parentIndex := getParentIndex(index)
// 	for this.items[parentIndex] > this.items[index] {
// 		this.swap(index, parentIndex)
// 		index = parentIndex
// 		parentIndex = getParentIndex(index)
// 	}
// }
// func (this *MinStack) Pop() {
// 	l := len(this.items)
// 	if l > 0 {
// 		this.items[0] = this.items[l-1]
// 		this.items = this.items[:l-1]
// 		this.Down(0)
// 	}
// }
// func (this *MinStack) Down(index int) {
// 	minChild := this.minChild(index)
// 	for minChild < len(this.items) {
// 		this.swap(index, minChild)
// 		index = minChild
// 		minChild = this.minChild(index)
// 	}
// }
// func (this *MinStack) Top() int {
// 	if len(this.items) > 0 {
// 		return this.items[0]
// 	}
// 	return -1 // error
// }

// func (this *MinStack) GetMin() int {
// 	if len(this.items) > 0 {
// 		return this.items[0]
// 	}
// 	return -1 // error
// }

// func getParentIndex(i int) int {
// 	return (i - 1) / 2
// }
// func (this *MinStack) swap(i, j int) {
// 	this.items[i], this.items[j] = this.items[j], this.items[i]
// }

// func (this *MinStack) minChild(index int) int {
// 	lc, rc := 2*index+1, 2*index+2
// 	if lc >= len(this.items) {
// 		return -1 // No children
// 	}
// 	if rc >= len(this.items) {
// 		return lc // Only left child exists, as previous check has passed which means lc lies in between [0,l-1]
// 	}
// 	if this.items[lc] <= this.items[rc] {
// 		return lc
// 	}
// 	return rc
// }
