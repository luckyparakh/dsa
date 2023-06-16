package main

import (
	"container/list"
	"fmt"
)

type llQueue struct {
	q *list.List
}

func initQ() *llQueue {
	return &llQueue{
		q: list.New(),
	}
}
func (ql *llQueue) Len() int {
	return ql.q.Len()
}

func (ql *llQueue) enqueue(item string) {
	ql.q.PushBack(item)
}

func (ql *llQueue) dequeue() error {
	if ql.Len() > 0 {
		ql.q.Remove(ql.q.Front())
		return nil
	}
	return fmt.Errorf("Empty queue")
}

func (ql *llQueue) front() (string, error) {
	if ql.q.Len() > 0 {
		return ql.q.Front().Value.(string), nil
	}
	return "", fmt.Errorf("Empty queue")
}
func (ql *llQueue) traverse() {

	for e := ql.q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(string))
	}
}

func (ql *llQueue) isEmpty() bool {
	return ql.Len() == 0
}

func main() {
	q := initQ()
	q.enqueue("a")
	q.enqueue("b")
	q.enqueue("c")
	q.traverse()

	q.dequeue()
	q.dequeue()
	q.traverse()
	fmt.Println(q.front())
	fmt.Println(q.isEmpty())
}
