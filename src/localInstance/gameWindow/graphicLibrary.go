package gameWindow

import "github.com/hajimehoshi/ebiten/v2"

// Check if the size of the window has changed
// Recalculate values for window graphics
func (g *Graphics) windowSizeChanged(screen *ebiten.Image) bool {
	if g.width != screen.Bounds().Dx() || g.height != screen.Bounds().Dy() {
		g.width = screen.Bounds().Dx()
		g.height = screen.Bounds().Dy()
		g.halfWidth = float64(g.width) / 2
		g.halfHeight = float64(g.height) / 2
		return true
	}
	return false
}
