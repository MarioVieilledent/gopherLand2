package gameWindow

import "github.com/hajimehoshi/ebiten/v2"

func (g *Graphics) drawEntities(screen *ebiten.Image) {

	// Draw own player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		g.halfWidth,
		g.halfHeight,
	)
	screen.DrawImage(g.game.Ressources.Elements[g.game.Player.Character].Img, op)

	// Draw other players
	xPlayer := g.game.Player.Pos.X
	yPlayer := g.game.Player.Pos.Y

	for nickname, pos := range g.game.PlayerPos {
		if g.game.Player.Nickname != nickname {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(
				(float64(pos.X)-xPlayer)*g.size+g.halfWidth,
				(float64(pos.Y)-yPlayer)*g.size+g.halfHeight,
			)

			screen.DrawImage(g.game.Ressources.Elements["p"].Img, op)
		}
	}
}
