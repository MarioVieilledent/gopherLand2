package gameWindow

import (
	"log"

	"gopherLand2/src/game"
	"gopherLand2/src/localInstance/input"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const WIDTH int = 987
const HEIGHT int = 610

type Graphics struct {
	game  *game.Game  // Instance of the game
	input input.Input // Instance of input/output for player control

	fullScreen bool // Window is in fullscreen mode or not

	size       float64 // Precomputed Size
	width      int     // current window width
	height     int     // current window height
	halfWidth  float64
	halfHeight float64
	bgScale    float64 // Scale for 4K 16/9 Background to screen size (calculated only when window is resized)
}

func (g *Graphics) Update() error {
	if inpututil.IsKeyJustReleased(ebiten.KeyF11) {
		g.fullScreen = !g.fullScreen
		ebiten.SetFullscreen(g.fullScreen)
	}

	g.input.Update()

	g.game.Run()
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

func OpenWindow(input input.Input, game *game.Game) {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GopherLand2")

	ebiten.SetTPS(ebiten.SyncWithFPS)
	ebiten.SetTPS(144)

	graphics := Graphics{
		game: game,

		fullScreen: false,

		size:       float64(game.Config.Size),
		input:      input,
		width:      0,
		height:     0,
		halfWidth:  0.0,
		halfHeight: 0.0,
		bgScale:    1.0,
	}

	err := ebiten.RunGame(&graphics)
	if err != nil {
		log.Fatal(err)
	}
}
