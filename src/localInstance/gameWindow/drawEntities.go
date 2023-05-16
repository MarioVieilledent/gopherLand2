package gameWindow

import "github.com/hajimehoshi/ebiten/v2"

func (g *Graphics) drawEntities(screen *ebiten.Image) {

	// Draw own player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		g.halfWidth,
		g.halfHeight,
	)
	screen.DrawImage(g.game.Ressources.Elements[g.game.Players[g.game.Nickname].Character].Img, op)

	// Draw other players
	xPlayer := g.game.Players[g.game.Nickname].Pos.X
	yPlayer := g.game.Players[g.game.Nickname].Pos.Y

	for nickname, pi := range g.game.Players {
		if g.game.Nickname != nickname {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(
				(float64(pi.Pos.X)-xPlayer)*g.size+g.halfWidth,
				(float64(pi.Pos.Y)-yPlayer)*g.size+g.halfHeight,
			)

			screen.DrawImage(g.game.Ressources.Elements[pi.Character].Img, op)
		}
	}
}
