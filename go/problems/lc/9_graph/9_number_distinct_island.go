// https://github.com/doocs/leetcode/blob/main/solution/0600-0699/0694.Number%20of%20Distinct%20Islands/README_EN.md

package main

import (
	"fmt"
)

func NumberofDistinctIslands(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	vn := make([][]bool, row)
	islands := map[string]bool{}
	for k := 0; k < row; k++ {
		vn[k] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 && !vn[i][j] {
				island := []rune{}
				dfs(i, j, &grid, &vn, i, j, &island)
				// fmt.Println(island)
				islands[string(island)] = true
			}
		}
	}
	// fmt.Println(islands)
	return len(islands)
}

func dfs(i, j int, grid *[][]int, vn *[][]bool, baseI, baseJ int, island *[]rune) {
	(*vn)[i][j] = true
	g := (*grid)
	*island = append((*island), rune(i-baseI), rune(j-baseJ))
	dirsRow := [4]int{-1, 1, 0, 0}
	dirsCol := [4]int{0, 0, -1, 1}
	for h := 0; h < 4; h++ {
		x, y := i+dirsRow[h], j+dirsCol[h]
		if x >= 0 && y >= 0 && x < len(g) && y < len(g[0]) && !(*vn)[x][y] && g[x][y] == 1 {
			dfs(x, y, grid, vn, baseI, baseJ, island)
		}
	}
}

func main() {

	// 3
	fmt.Println(NumberofDistinctIslands([][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
	}))
	// 2
	fmt.Println(NumberofDistinctIslands([][]int{
		{1, 1, 0, 1, 1},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{1, 1, 0, 1, 0},
	}))
}
