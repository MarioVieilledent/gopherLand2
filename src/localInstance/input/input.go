package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	playerInputChannel chan KeyPressed
	keyPressed         KeyPressed
}

func New(ch chan KeyPressed, nickname string) Input {
	return Input{
		playerInputChannel: ch,
		keyPressed: KeyPressed{
			Nickname: nickname,
			Up:       false,
			Right:    false,
			Down:     false,
			Left:     false,
		},
	}
}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !i.keyPressed.Up {
			i.keyPressed.Up = true
			i.sendPlayerInput()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) ||
		ebiten.IsKeyPressed(ebiten.KeyD) {
		if !i.keyPressed.Right {
			i.keyPressed.Right = true
			i.sendPlayerInput()
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyA) {
		if !i.keyPressed.Left {
			i.keyPressed.Left = true
			i.sendPlayerInput()
		}
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustReleased(ebiten.KeyW) ||
		inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		if i.keyPressed.Up {
			i.keyPressed.Up = false
			i.sendPlayerInput()
		}
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustReleased(ebiten.KeyD) {
		if i.keyPressed.Right {
			i.keyPressed.Right = false
			i.sendPlayerInput()
		}
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustReleased(ebiten.KeyA) {
		if i.keyPressed.Left {
			i.keyPressed.Left = false
			i.sendPlayerInput()
		}
	}
}

func (i Input) sendPlayerInput() {
	i.playerInputChannel <- i.keyPressed
}
