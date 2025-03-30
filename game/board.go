package game

import "fmt"

type Board struct {
	Map [][]string
}

func NewBoard() *Board {
	mapa := make([][]string, height)
	for i := range mapa {
		mapa[i] = make([]string, width)
	}
	return &Board{Map: mapa}
}

const width int = 40
const height int = 20

func (b *Board) CreateMap() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == 0 || j == 0 || i == height-1 || j == width-1 {
				b.Map[i][j] = "#"
			} else {
				b.Map[i][j] = " "
			}
		}
	}
}

func (b Board) ShowMap() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(b.Map[i][j])
		}
		fmt.Println()
	}
}

func (b *Board) AddIntoMap(player *Player, left bool, right bool, up bool, down bool) {
	fmt.Println(left, len(player.Body))

	playerSize := len(player.Body)
	height := player.X
	for i := player.Y; i < player.Y+playerSize; i++ {
		b.Map[player.X][i] = player.Body[0]
		b.clearPath(player, i, height, left, right, up, down)
	}
}

func (b *Board) clearPath(player *Player, i, j int, left, right, up, down bool) {
	if left {
		b.Map[player.X][i+1] = " "
	} else if right {
		b.Map[player.X][i-1] = " "
	} else if up {
		if j > -1 {
			b.Map[j+1][player.Y] = " "
		}
	} else if down {
		if j < height-1 {
			b.Map[j-1][player.Y] = " "
		}
	}
}

// 0,0 0,1 0,2 0,3
// 1,0 1,1 1,2 1,3
// 2,0 2,1 2,2 2,3
// 3,0 3,1 3,2 3,3
