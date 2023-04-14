package launcherWindow

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Graphics struct {
}

func (g *Graphics) Update() error {
	return nil
}

func (g *Graphics) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 200, 50, 255})
}

func (g *Graphics) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func OpenWindow() {
	ebiten.SetWindowSize(987, 610)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("GopherLand2 Launcher")

	graphics := Graphics{}

	if err := ebiten.RunGame(&graphics); err != nil {
		log.Fatal(err)
	}
}
