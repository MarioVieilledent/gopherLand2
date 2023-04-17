package gameWindow

import "github.com/hajimehoshi/ebiten/v2"

func (g *Graphics) drawEntities(screen *ebiten.Image, playingPlayerId int) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		g.halfWidth,
		g.halfHeight,
	)
	screen.DrawImage(g.game.Ressources.Elements["P"].Img, op)

	/* Old Method for preparing drawing all players
	for playerId, player := range g.game.Players {
		op := &ebiten.DrawImageOptions{}

		if playingPlayerId == playerId {
			op.GeoM.Translate(
				g.halfWidth,
				g.halfHeight,
			)
		} else {
			op.GeoM.Translate(
				player.Pos.X*g.size,
				player.Pos.Y*g.size,
			)
		}

		screen.DrawImage(g.game.Ressources.Elements["P"].Img, op)
	}
	*/
}
