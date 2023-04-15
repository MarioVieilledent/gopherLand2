package gameWindow

import (
	"log"

	"gopherLand2/src/game"
	"gopherLand2/src/localInstance/io"

	"github.com/hajimehoshi/ebiten/v2"
)

const WIDTH int = 987
const HEIGHT int = 610

type Graphics struct {
	game    game.Game // Instance of the game
	io      io.Io     // Instance of input/output for player control
	width   int       // current window width
	height  int       // current window height
	bgScale float64   // Scale for 4K 16/9 Background to screen size (calculed only when window is resized)
}

func (g *Graphics) Update() error {
	g.io.Update()
	return nil
}

func (g *Graphics) Draw(screen *ebiten.Image) {
	g.drawBasics(screen)
	g.drawBlocks(screen)
	g.drawEntities(screen)
	g.drawDebug(screen)
}

func (g *Graphics) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func OpenWindow(io io.Io, game game.Game) {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GopherLand2")
	ebiten.SetTPS(144)

	graphics := Graphics{
		game:    game,
		io:      io,
		width:   0,
		height:  0,
		bgScale: 1.0,
	}

	err := ebiten.RunGame(&graphics)
	if err != nil {
		log.Fatal(err)
	}
}
