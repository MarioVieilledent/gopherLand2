package gameWindow

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Graphics) drawBlocks(screen *ebiten.Image) {
	xPlayer := g.game.Player.Pos.X
	yPlayer := g.game.Player.Pos.Y

	for y, line := range g.game.GameMap.Blocks {
		for x, res := range line {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(
				(float64(x)-xPlayer)*g.size+g.halfWidth,
				(float64(y)-yPlayer)*g.size+g.halfHeight,
			)

			val, ok := g.game.Ressources.Elements[res]
			if ok && val.Code != 0 {
				screen.DrawImage(val.Img, op)
			}
		}
	}
}
