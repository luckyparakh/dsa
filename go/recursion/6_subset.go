package main

// https://www.youtube.com/watch?v=Yg5a2FxU4Fo&list=PL_z_8CaSLPWeT1ffjiImo0sYTcnLzo-wY&index=12
// https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b

import "fmt"

func subsets(nums []int) {
	// return getSubsets(nums, []int{})
	op := [][]int{}
	getSubsetsA(nums, []int{}, &op)
	fmt.Println(op)
}
func getSubsets(nums []int, subSet []int) [][]int {
	fmt.Println("inputs", subSet, nums)
	op := [][]int{}
	if len(nums) == 0 {
		op = append(op, subSet)
		fmt.Println("*********", op)
		return op
	}
	op1 := getSubsets(nums[1:], subSet)
	fmt.Println("*********op1", op1)
	op = append(op, op1...)
	subSet = append(subSet, nums[0])
	op2 := getSubsets(nums[1:], subSet)
	fmt.Println("*********op2", op2)

	op = append(op, op2...)
	fmt.Println("###########", op)
	// op = getSubsets(nums[0:len(nums)-1], subSet, op)
	// op = getSubsets(nums[0:len(nums)-1], append(subSet, nums[len(nums)-1]), op)
	return op
}
func getSubsetsA(nums []int, subSet []int, op *[][]int) {
	if len(nums) == 0 {
		*op = append(*op, subSet)
		// fmt.Println("*********", *op)
		return
	}
	// fmt.Println("@@@@@@@@@", subSet, nums)
	getSubsetsA(nums[1:], subSet, op)
	subSet = append(subSet, nums[0])
	getSubsetsA(nums[1:], subSet, op)
	// op = append(op, op1...)
	// op = append(op, op2...)
	// fmt.Println("###########", *op)
	// op = getSubsets(nums[0:len(nums)-1], subSet, op)
	// op = getSubsets(nums[0:len(nums)-1], append(subSet, nums[len(nums)-1]), op)
	return
}

func main() {
	// subsets([]int{1, 2})
	// subsets([]int{1, 2, 3, 4})
	subsets([]int{1, 2, 3, 4, 5}) // does not give correct o/p as [1,2,3,4] set repeats
	// fmt.Println(subsets([]int{1, 2, 3, 4, 5, 6}))
	// fmt.Println(subsets([]int{}))
	// fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

// func main() {
// 	slice := []string{"a", "a"}
// 	func(slice []string) {
// 		slice[0] = "b"
// 		slice[1] = "b"
// 		fmt.Print(slice)
// 	}(slice)
// 	fmt.Print(slice)

// 	slice1 := []string{"a", "a"}
// 	func(slice []string) {
// 		slice = append(slice, "a")
// 		slice[0] = "b"
// 		slice[1] = "b"
// 		fmt.Print(slice)
// 	}(slice1)
// 	fmt.Print(slice1)
// }
