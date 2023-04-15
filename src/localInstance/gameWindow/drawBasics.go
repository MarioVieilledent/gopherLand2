package gameWindow

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const BG_IMG_WIDTH int = 3840
const BG_IMG_HEIGHT int = 2160

// Draw background
func (g *Graphics) drawBasics(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	if g.windowSizeChanged(screen) {
		g.getScaleForBackground(screen.Bounds())
	}

	op.GeoM.Scale(g.bgScale, g.bgScale)
	screen.DrawImage(g.game.Ressources.BackgroundImage, op)
}

// Calculate scale for drawing nicely background image
func (g *Graphics) getScaleForBackground(screenBounds image.Rectangle) {
	sX := screenBounds.Dx()
	sY := screenBounds.Dy()

	if float64(BG_IMG_WIDTH)/float64(BG_IMG_HEIGHT) > float64(sX)/float64(sY) {
		g.bgScale = float64(sY) / float64(BG_IMG_HEIGHT)
	} else {
		g.bgScale = float64(sX) / float64(BG_IMG_WIDTH)
	}
}
