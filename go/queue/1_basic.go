package main

import "fmt"

// Keep in consideration if queue is empty

type customQueue struct {
	queue []string
}

func initQ() *customQueue {
	return &customQueue{}
}
func (cq *customQueue) Len() int {
	return len(cq.queue)
}

func (cq *customQueue) enqueue(item string) {
	cq.queue = append(cq.queue, item)
}

func (cq *customQueue) dequeue() error {
	// as 1st element is removed from queue and it may cause memory leak
	if len(cq.queue) != 0 {
		cq.queue = cq.queue[1:]
		return nil
	}
	return fmt.Errorf("queue is empty")
}

func (cq *customQueue) front() (string, error) {
	if len(cq.queue) != 0 {
		return cq.queue[0], nil
	}
	return "", fmt.Errorf("queue is empty")
}

func (cq *customQueue) Size() int {
	return len(cq.queue)
}

func (cq *customQueue) isEmpty() bool {
	return len(cq.queue) == 0
}

func main() {
	q := initQ()
	q.enqueue("a")
	q.enqueue("b")
	q.enqueue("c")
	fmt.Println(q)

	q.dequeue()
	q.dequeue()
	fmt.Println(q)
	fmt.Println(q.front())
	fmt.Println(q.isEmpty())
}
