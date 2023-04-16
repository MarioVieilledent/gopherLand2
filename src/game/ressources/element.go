package ressources

import "github.com/hajimehoshi/ebiten/v2"

// A block
type Element struct {
	Name  string `json:"name"`  // Name
	Code  int    `json:"code"`  // Unique code
	Solid bool   `json:"solid"` // Collision with entities
	X     int    `json:"x"`     // X position in the ressourcePack
	Y     int    `json:"y"`     // Y position in the ressourcePack
	W     int    `json:"w"`     // Width (number of blocks taken by element going right)
	H     int    `json:"h"`     // Height (number of blocks taken by element going down)
	Img   *ebiten.Image
}
