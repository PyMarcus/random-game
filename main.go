package main

import (
	"github.com/PyMarcus/snake-game/game"
)

func main() {
	board := game.NewBoard()
	player := game.NewPlayer()
	for {
		if game.Play(player, board) {
			break
		}
	}
}
