package main

import (
	"fmt"
)

func tohDriver(numDisc int) {
	fmt.Println("Tower of Hanoi")
	auxPole := "aux"
	startPole := "start"
	endPole := "end"
	toh(numDisc, startPole, auxPole, endPole)
}
func toh(numDisc int, s, a, e string) {
	if numDisc == 1 {
		fmt.Printf("Move %d from %s to %s\n", numDisc, s, e)
		return
	}
	toh(numDisc-1, s, e, a)
	fmt.Printf("Move %d from %s to %s\n", numDisc, s, e)
	toh(numDisc-1, a, s, e)
	return
}
func main() {
	tohDriver(1)
	tohDriver(2)
	tohDriver(3)
}
