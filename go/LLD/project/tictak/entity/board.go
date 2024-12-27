package entity

import "fmt"

type Board struct {
	board     [][]Piece
	size      int
	freeSpace int
}

func NewBoard(size int) *Board {
	b := make([][]Piece, size)
	for i := range b {
		v := make([]Piece, size)
		for j := range v {
			v[j] = E
		}
		b[i] = v
	}
	return &Board{size: size, board: b, freeSpace: size * size}
}

func (b *Board) GetSize() int {
	return b.size
}

func (b *Board) checkWin(p Piece, row, column int) bool {
	status := true
	for i := 0; i < b.size; i++ {
		if b.board[i][i] != p {
			status = false
			break
		}
	}
	if status {
		return status
	}
	status = true
	for i := 0; i < b.size; i++ {
		if b.board[i][b.size-1-i] != p {
			status = false
			break
		}
	}
	if status {
		return status
	}
	status = true

	for j := 0; j < b.size; j++ {
		if b.board[row][j] != p {
			status = false
			break
		}
	}
	if status {
		return status
	}
	status = true

	for i := 0; i < b.size; i++ {
		if b.board[i][column] != p {
			status = false
			break
		}
	}
	return status
}

func (b *Board) print() {
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if j == 0 {
				fmt.Print(b.board[i][j])
			} else {
				fmt.Print(" | ", b.board[i][j])
			}

		}
		fmt.Println("")
	}
}

func (b *Board) isEmpty() bool {
	return b.freeSpace > 0
}
