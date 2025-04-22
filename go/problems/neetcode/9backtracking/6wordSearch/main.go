package main

func exist(board [][]byte, word string) bool {
	r := len(board)
	c := len(board[0])
	visited := make([][]int, r)
	for i := range r {
		t := make([]int, c)
		visited[i] = t
	}
	var h func(int, int, int, []byte) bool
	h = func(x, y int, index int, t []byte) bool {
		if index == len(word) {
			if word == string(t) {
				return true
			}
		}
		if x >= 0 && x < r && y >= 0 && y < c && visited[x][y] == 0 && board[x][y] == word[index] {
			visited[x][y] = 1
			t = append(t, word[index])
			if h(x+1, y, index+1, t) || h(x-1, y, index+1, t) || h(x, y+1, index+1, t) || h(x, y-1, index+1, t) {
				return true
			}
			visited[x][y] = 0
			t = t[:len(t)-1]
		}
		return false
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if board[i][j] == word[0] {
				if h(i, j, 0, []byte{}) {
					return true
				}
			}
		}
	}
	return false
}
