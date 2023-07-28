// https://www.youtube.com/watch?v=U81n0UYtk98&list=PL_z_8CaSLPWeT1ffjiImo0sYTcnLzo-wY&index=18
// Print N-bit binary numbers having more 1’s than 0’s
package main

import (
	"fmt"
)

func printAnswer(n int) {
	mOp := [][]int{}
	driver(&mOp, []int{}, n, 0, 0)
	fmt.Println(mOp)
}
func driver(mOp *[][]int, op []int, total, numberof1, numberof0 int) {
	// if numberof0 > total/2 {
	// 	return
	// }
	if numberof0+numberof1 == total {
		*mOp = append(*mOp, op)
		fmt.Println(op)
		return
	}
	if numberof0 < numberof1 {
		driver(mOp, append(op, 0), total, numberof1, numberof0+1)
	}
	driver(mOp, append(op, 1), total, numberof1+1, numberof0)
}
func main() {
	printAnswer(2)
	printAnswer(3)
	printAnswer(4)
}
