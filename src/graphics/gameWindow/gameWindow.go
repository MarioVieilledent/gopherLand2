package gameWindow

import (
	"fmt"
	"log"

	"gopherLand2/src/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Graphics struct {
	game game.Game
}

func (g *Graphics) Update() error {
	return nil
}

func (g *Graphics) Draw(screen *ebiten.Image) {
	for y, line := range g.game.GameMap.Blocks {
		for x, res := range line {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*g.game.Config.Size), float64(y*g.game.Config.Size))
			val, ok := g.game.Ressources.Blocks[res]
			if ok {
				screen.DrawImage(val.Img, op)
			} else {
				fmt.Println("ressouce innexistante : " + res)
			}
		}
	}
}

func (g *Graphics) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func OpenWindow() {
	ebiten.SetWindowSize(987, 610)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GopherLand2")

	graphics := Graphics{
		game: game.New(),
	}

	if err := ebiten.RunGame(&graphics); err != nil {
		log.Fatal(err)
	}
}
