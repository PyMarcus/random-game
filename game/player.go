package game

type Player struct {
	Body []string
	Y    int // w
	X    int // h
}

func NewPlayer() *Player {
	return &Player{
		X:    3,
		Y:    15,
		Body: []string{"*"},
	}
}

func (p *Player) ToEat(food string) {
	p.Body = append(p.Body, food)
}

func (p *Player) ToReduce() {
	p.Body = p.Body[:len(p.Body)-1]
}

func (p *Player) MoveLeft() {
	if p.Y > 0 {
		p.Y--
	}
}

func (p *Player) MoveRight() {
	if p.Y < width-1 {
		p.Y++
	}
}

func (p *Player) MoveUp() {
	if p.X > 0 {
		p.X--
	}
}

func (p *Player) MoveDown() {
	if p.X < height-1 {
		p.X++
	}
}
