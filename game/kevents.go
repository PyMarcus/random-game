package game

import (
	"fmt"
	"os"

	hook "github.com/eiannone/keyboard"
)

func Play(player *Player, board *Board) bool {
	return eventListener(player, board)
}

func eventListener(player *Player, board *Board) bool {
	board.CreateMap()
	board.ShowMap()

	board.AddIntoMap(player, false, false, false, false)
	board.ShowMap()

	err := hook.Open()
	if err != nil {
		fmt.Println("Fail to listen keyboard " + err.Error())
		return true
	}
	defer hook.Close()
	status := false

	key, _, err := hook.GetKey()
	if err != nil {
		fmt.Println("Erro ao capturar tecla:", err)
		status = true
		panic(err)
	}
	switch key {
	case 'w':
		player.MoveUp()
		board.AddIntoMap(player, false, false, true, false)
	case 'a':
		player.MoveLeft()
		board.AddIntoMap(player, true, false, false, false)
	case 's':
		player.MoveDown()
		board.AddIntoMap(player, false, false, false, true)
	case 'd':
		player.MoveRight()
		board.AddIntoMap(player, false, true, false, false)
	case 'i':
		fmt.Println("Saindo...")
		status = true
		os.Exit(0)
	}
	board.ShowMap()

	return status
}
