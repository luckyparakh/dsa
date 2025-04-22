package main

import "fmt"

func ratInMaze(mat [][]int) []string {
	ans := []string{}
	r := []rune{}
	rows, cols := len(mat), len(mat[0])
	d := make([][]int, rows)
	for i := range d {
		t := make([]int, cols)
		d[i] = t
	}
	var h func(x, y int, r []rune)
	h = func(x, y int, r []rune) {
		if x == rows-1 && y == cols-1 {
			ans = append(ans, string(r))
			return
		}
		if x >= 0 && y >= 0 && x < rows && y < cols && d[x][y] == 0 && mat[x][y] == 1 {
			d[x][y] = 1
			r = append(r, 'D')
			h(x+1, y, r)
			r = r[:len(r)-1] // backtrack
			r = append(r, 'L')
			h(x, y-1, r)
			r = r[:len(r)-1] // backtrack
			r = append(r, 'R')
			h(x, y+1, r)
			r = r[:len(r)-1] // backtrack
			r = append(r, 'U')
			h(x-1, y, r)
			r = r[:len(r)-1] // backtrack
			d[x][y] = 0      // backtrack
		} else {
			return
		}
	}
	if mat[0][0] == 1 {
		h(0, 0, r)
	}
	return ans
}

func main() {
	// Test Cases
	maze1 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	fmt.Println("Maze 1:", ratInMaze(maze1)) // Output: [DDRR DRDR]

	maze2 := [][]int{
		{1, 1, 1},
		{0, 1, 0},
		{0, 1, 1},
	}
	fmt.Println("Maze 2:", ratInMaze(maze2)) // Output: [RRDD]

	maze3 := [][]int{
		{0, 0},
		{0, 0},
	}
	fmt.Println("Maze 3:", ratInMaze(maze3)) // Output: []

	maze4 := [][]int{
		{1},
	}
	fmt.Println("Maze 4:", ratInMaze(maze4)) // Output: []

	maze5 := [][]int{
		{1, 0},
		{1, 1},
	}
	fmt.Println("Maze 5:", ratInMaze(maze5)) // Output: [DR]

	maze6 := [][]int{
		{1, 1},
		{0, 1},
	}
	fmt.Println("Maze 6:", ratInMaze(maze6)) // Output: [RD]

	maze7 := [][]int{
		{0, 1},
		{1, 1},
	}
	fmt.Println("Maze 7:", ratInMaze(maze7)) // Output: []

	maze8 := [][]int{
		{1, 0, 1},
		{1, 1, 1},
		{0, 0, 1},
	}
	fmt.Println("Maze 8:", ratInMaze(maze8)) // Output: [DDRR]

	maze9 := [][]int{
		{1, 1, 0, 0},
		{0, 1, 1, 1},
		{0, 0, 0, 1},
	}
	fmt.Println("Maze 9:", ratInMaze(maze9)) // Output: [RRDD]

	maze10 := [][]int{
		{1, 0, 0},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println("Maze 10:", ratInMaze(maze10)) // Output: [DDRR]

	maze11 := [][]int{
		{1, 0},
	}
	fmt.Println("Maze 11:", ratInMaze(maze11)) //output: []

	maze12 := [][]int{
		{1, 1},
	}
	fmt.Println("Maze 12:", ratInMaze(maze12)) //output: [RD]

}
