package entity

import (
	"fmt"
)

type Game struct {
	board   Board
	players []Player
}

func NewGame(board Board, players []Player) *Game {
	return &Game{board: board, players: players}

}

func (g *Game) Start() {
	// Start the game
	runGame := true
	playerNumber := 0
	winner := ""
	for runGame {
		if g.board.isEmpty() {
			x, y, err := readInput(playerNumber)
			if err != nil {
				continue
			}
			if x < 0 || x >= g.board.size || y < 0 || y >= g.board.size || g.board.board[x][y] != E {
				fmt.Println("Invalid coordinates")
				continue
			}
			g.board.board[x][y] = g.players[playerNumber].piece
			g.board.freeSpace -= 1
			g.board.print()
			if g.board.checkWin(g.players[playerNumber].piece, x, y) {
				runGame = false
				winner = fmt.Sprintf("Player %d", playerNumber)
			}
			playerNumber = (playerNumber + 1) % len(g.players)
		} else {
			runGame = false
			winner = "tie"
		}
	}
	fmt.Printf("Winner: %s\n", winner)
}

func readInput(playerNumber int) (int, int, error) {
	var x, y int
	fmt.Printf("Player %d: Enter the x and y coordinates: ", playerNumber)
	_, err := fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println(err)
		return x, y, err
	}
	return x, y, nil
}
