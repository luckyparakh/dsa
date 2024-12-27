package main

import "tt/entity"

func initializeGame(players []entity.Player, size int) {
	board := entity.NewBoard(size)
	game := entity.NewGame(*board, players)
	game.Start()
}
func main() {
	player1 := entity.NewPlayer("John", entity.X)
	player2 := entity.NewPlayer("John", entity.O)
	players := []entity.Player{*player1, *player2}
	initializeGame(players, 3)
}
