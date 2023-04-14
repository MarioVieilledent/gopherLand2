package entity

type Player struct {
	Pos Pos
}

func NewPlayer(pos Pos) Player {
	return Player{
		Pos: pos,
	}
}
