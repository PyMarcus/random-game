package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIntoMap(t *testing.T) {
	player := NewPlayer()
	board := NewBoard()
	board.CreateMap()

	player.ToEat("*")
	player.ToEat("*")
	board.AddIntoMap(player, false, false, false, false)
	board.ShowMap()
	assert.Equal(t, board.Map[player.X][player.Y], "*")
	assert.Equal(t, board.Map[player.X][player.Y], "*")
}
