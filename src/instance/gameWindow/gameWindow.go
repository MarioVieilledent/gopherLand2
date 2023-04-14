package gameWindow

import (
	"log"

	"gopherLand2/src/game"
	"gopherLand2/src/instance/io"

	"github.com/hajimehoshi/ebiten/v2"
)

type Graphics struct {
	game game.Game
	io   io.Io
}

func (g *Graphics) Update() error {
	g.io.Update()
	return nil
}

func (g *Graphics) Draw(screen *ebiten.Image) {
	screen.Clear()
	for y, line := range g.game.GameMap.Blocks {
		for x, res := range line {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(
				(float64(x)-g.game.Player[0].Pos.X)*float64(g.game.Config.Size),
				float64(y*g.game.Config.Size),
			)

			val, ok := g.game.Ressources.Blocks[res]
			if ok && val.Code != 0 {
				screen.DrawImage(val.Img, op)
			}
		}
	}
}

func (g *Graphics) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func OpenWindow(io io.Io, game game.Game) {
	ebiten.SetWindowSize(987, 610)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GopherLand2")
	ebiten.SetTPS(60)

	graphics := Graphics{
		game: game,
		io:   io,
	}

	if err := ebiten.RunGame(&graphics); err != nil {
		log.Fatal(err)
	}
}
