package ressources

import "github.com/hajimehoshi/ebiten/v2"

// A block
type Block struct {
	Name string `json:"name"`
	Code int    `json:"code"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Img  *ebiten.Image
}
