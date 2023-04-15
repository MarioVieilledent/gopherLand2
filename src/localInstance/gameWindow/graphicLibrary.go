package gameWindow

import "github.com/hajimehoshi/ebiten/v2"

// Check if the size of the window has changed
func (g *Graphics) windowSizeChanged(screen *ebiten.Image) bool {
	if g.width != screen.Bounds().Dx() || g.height != screen.Bounds().Dy() {
		g.width = screen.Bounds().Dx()
		g.height = screen.Bounds().Dy()
		return true
	}
	return false
}
