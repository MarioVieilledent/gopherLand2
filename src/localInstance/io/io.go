package io

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Io struct {
	channel chan string
}

func New(ch chan string) Io {
	return Io{
		channel: ch,
	}
}

func (io Io) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeySpace) {
		io.channel <- "up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) ||
		ebiten.IsKeyPressed(ebiten.KeyD) {
		io.channel <- "right"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyS) {
		io.channel <- "down"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyA) {
		io.channel <- "left"
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustReleased(ebiten.KeyD) {
		io.channel <- "released_right"
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustReleased(ebiten.KeyA) {
		io.channel <- "released_left"
	}
}
