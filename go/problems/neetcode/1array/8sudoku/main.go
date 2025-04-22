package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[int]bool, 9)
	cols := make([]map[int]bool, 9)
	squares := make([]map[int]bool, 9)
	for i := range 9 {
		rows[i] = make(map[int]bool)
		cols[i] = make(map[int]bool)
		squares[i] = make(map[int]bool)
	}
	for i := range 9 {
		for j := range 9 {
			val := board[i][j]
			if val == '.' {
				continue
			}
			// the formula efficiently maps a 2D cell coordinate (row, column) to a 1D subgrid index
			square := (i/3)*3 + j/3
			if rows[i][int(val)] {
				return false
			} else {
				rows[i][int(val)] = true
			}
			if cols[j][int(val)] {
				return false
			} else {
				cols[j][int(val)] = true
			}
			if squares[square][int(val)] {
				return false
			} else {
				squares[square][int(val)] = true
			}
		}
	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)
	for i := range 9 {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}
	for r := range 9 {
		for c := range 9 {
			val := board[r][c]
			if val == '.' {
				continue
			}
			// the formula efficiently maps a 2D cell coordinate (row, column) to a 1D subgrid index
			square := (r/3)*3 + c/3
			if rows[r][val] || cols[c][val] || squares[square][val] {
				return false
			}
			rows[r][val] = true
			cols[c][val] = true
			squares[square][val] = true

		}
	}
	return true
}
