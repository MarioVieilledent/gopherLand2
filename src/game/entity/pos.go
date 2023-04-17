package entity

import "fmt"

type Pos struct {
	X float64
	Y float64
}

func (p Pos) ToString() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func ParsePos(s string) (Pos, error) {
	var p Pos
	_, err := fmt.Sscanf(s, "(%f, %f)", &p.X, &p.Y)
	if err != nil {
		return Pos{}, err
	}
	return p, nil
}
