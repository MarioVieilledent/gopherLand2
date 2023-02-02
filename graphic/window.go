package graphic

import (
	"gopherLand2/engine3D"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const windowWidth int = 1280
const windowHeight int = 720

var backgroundImage *ebiten.Image

type Controller struct {
	engine engine3D.Engine3D // 3D engine
	tick   uint64            // Ticks of the game
	theta  float64           // A suppr
}

//////////////////////////////
// INITIALIZATION FUNCTIONS //
//////////////////////////////

func initController() Controller {
	objs := []engine3D.Line{
		{
			Start: engine3D.Point{
				X: -1,
				Y: -1,
				Z: -1,
			},
			End: engine3D.Point{
				X: -1,
				Y: -1,
				Z: 1,
			},
		},
		{
			Start: engine3D.Point{
				X: -1,
				Y: -1,
				Z: 1,
			},
			End: engine3D.Point{
				X: -1,
				Y: 1,
				Z: 1,
			},
		},
		{
			Start: engine3D.Point{
				X: -1,
				Y: 1,
				Z: 1,
			},
			End: engine3D.Point{
				X: -1,
				Y: 1,
				Z: -1,
			},
		},
		{
			Start: engine3D.Point{
				X: -1,
				Y: 1,
				Z: -1,
			},
			End: engine3D.Point{
				X: -1,
				Y: -1,
				Z: -1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: -1,
				Z: -1,
			},
			End: engine3D.Point{
				X: 1,
				Y: -1,
				Z: 1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: -1,
				Z: 1,
			},
			End: engine3D.Point{
				X: 1,
				Y: 1,
				Z: 1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: 1,
				Z: 1,
			},
			End: engine3D.Point{
				X: 1,
				Y: 1,
				Z: -1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: 1,
				Z: -1,
			},
			End: engine3D.Point{
				X: 1,
				Y: -1,
				Z: -1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: 1,
				Z: 1,
			},
			End: engine3D.Point{
				X: -1,
				Y: 1,
				Z: 1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: 1,
				Z: -1,
			},
			End: engine3D.Point{
				X: -1,
				Y: 1,
				Z: -1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: -1,
				Z: 1,
			},
			End: engine3D.Point{
				X: -1,
				Y: -1,
				Z: 1,
			},
		},
		{
			Start: engine3D.Point{
				X: 1,
				Y: -1,
				Z: -1,
			},
			End: engine3D.Point{
				X: -1,
				Y: -1,
				Z: -1,
			},
		},
	}

	return Controller{engine3D.Engine3D{
		Cam: engine3D.Camera{
			Name:  "FPS camera",
			Focal: 4,
			Angle: 0,
			Pos:   engine3D.Point{X: 0.0, Y: 0.0, Z: -6.0},
		},
		Scale:          120.0,
		Objs:           objs,
		ScreenSettings: engine3D.CreateScreenSettings(windowWidth, windowHeight),
	}, 0, 0.0}
}

func init() {
	var err error
	backgroundImage, _, err = ebitenutil.NewImageFromFile("data/images/test.png")
	if err != nil {
		log.Fatal(err)
	}
}

//////////////////////
// UPDATE FUNCTIONS //
//////////////////////

// Update function, called each frame
func (c *Controller) Update() error {
	c.manageButtonClicks()
	c.tick++
	// c.engine.Cam.Focal += 0.01
	// c.engine.Cam.Pos.Translate(engine3D.Point{0.0, 0.0, 0.01})
	// c.engine.Cam.Angle += 0.01
	return nil
}

// Manages input for controlling player
func (c *Controller) manageButtonClicks() {
	// Left button clicked
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.engine.Cam.Pos.Translate(engine3D.Point{-0.1, 0.0, 0.0})
		c.theta += 0.01
	}

	// Right button clicked
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.engine.Cam.Pos.Translate(engine3D.Point{0.1, 0.0, 0.0})
		c.theta -= 0.01
	}

	// Up button clicked
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		c.engine.Cam.Pos.Translate(engine3D.Point{0.0, 0.0, 0.1})
	}

	// Down button clicked
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		c.engine.Cam.Pos.Translate(engine3D.Point{0.0, 0.0, -0.1})
	}
}

///////////////////////
// DRAWING ON WINDOW //
///////////////////////

func (c *Controller) Draw(screen *ebiten.Image) {
	for _, line := range c.engine.Objs {
		// c.DrawLine(screen, line)
		c.DrawLine(screen, line.Rotate(c.theta))
	}
}

func (c *Controller) DrawLine(screen *ebiten.Image, line engine3D.Line) {
	for i := 0.0; i <= 1.0; i += 0.007 {
		op := &ebiten.DrawImageOptions{}
		x1, y1 := c.engine.PlotPoint(line.Start)
		x2, y2 := c.engine.PlotPoint(line.End)
		op.GeoM.Translate((1-i)*x1+i*x2, (1-i)*y1+i*y2)
		screen.DrawImage(backgroundImage, op)
	}
}

/////////////////////
// OTHER FUNCTIONS //
/////////////////////

func (c *Controller) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func OpenWindow() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("GopherLand2")
	ebiten.SetWindowIcon([]image.Image{backgroundImage})

	controler := initController()

	err := ebiten.RunGame(&controler)

	if err != nil {
		log.Fatal(err)
	}
}
