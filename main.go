package main

import (
	"time"

	"github.com/PyMarcus/snake-game/game"
)

func main() {
	board := game.NewBoard()
	player := game.NewPlayer()

	board.CreateMap()
	board.ShowMap()

	board.AddIntoMap(player, false, false, false, false)
	board.ShowMap()
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second * 1)
		player.MoveDown()
		board.AddIntoMap(player, false, false, false, true)
		board.ShowMap()
	}
}
