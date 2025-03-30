package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToEat(t *testing.T) {
	player := NewPlayer()
	player.ToEat("*")

	assert.NotEmpty(t, player.Body)
	assert.Equal(t, "*", player.Body[0])

	player.ToEat("*")
	player.ToEat("*")
	assert.Equal(t, len(player.Body), 4)
}

func TestToReduce(t *testing.T) {
	player := NewPlayer()
	player.ToReduce()
	assert.Empty(t, player.Body)
	player.ToEat("*")
	player.ToEat("*")
	assert.Equal(t, len(player.Body), 2)
	player.ToReduce()
	assert.Equal(t, len(player.Body), 1)
}

func TestMoveLeft(t *testing.T) {
	player := NewPlayer()
	assert.Equal(t, player.Y, 15)
	player.MoveLeft()
	assert.Equal(t, player.Y, 14)
	player.Y = 0
	player.MoveLeft()
	assert.Equal(t, player.Y, 0)
}

func TestMoveRight(t *testing.T) {
	player := NewPlayer()
	assert.Equal(t, player.Y, 15)
	player.MoveRight()
	assert.Equal(t, player.Y, 16)
	player.Y = width - 1
	player.MoveRight()
	assert.Equal(t, player.Y, width-1)
}

func TestMoveUp(t *testing.T) {
	player := NewPlayer()
	assert.Equal(t, player.X, 3)
	player.MoveUp()
	assert.Equal(t, player.X, 2)
}

func TestMoveDown(t *testing.T) {
	player := NewPlayer()
	assert.Equal(t, player.X, 3)
	player.MoveDown()
	assert.Equal(t, player.X, 4)
}
