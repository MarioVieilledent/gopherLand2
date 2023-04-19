package entity

const DEFAULT_SPEED float64 = 0.08
const DEFAULT_JUMP_SPEED float64 = 0.21
const GRAVITY float64 = 0.0065

type Player struct {
	// Identification properties
	Nickname string

	// Movement properties
	Pos              Pos
	TouchesGround    bool
	VerticalVelocity float64
	MovesLeft        bool
	MovesRight       bool

	// Statistics of player
	EatBox    [4]float64
	JumpSpeed float64
	Speed     float64
}

// Create a new player
func NewPlayer(pos Pos) Player {
	return Player{
		Pos:           pos,
		TouchesGround: false,
		MovesLeft:     false,
		MovesRight:    false,

		EatBox:    [4]float64{0.1, 0.1, 0.9, 1.95},
		Speed:     DEFAULT_SPEED,
		JumpSpeed: DEFAULT_JUMP_SPEED,
	}
}
