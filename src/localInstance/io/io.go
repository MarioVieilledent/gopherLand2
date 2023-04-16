package io

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Io struct {
	channels []chan string
}

func New(chs []chan string) Io {
	return Io{
		channels: chs,
	}
}

func (io Io) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsKeyPressed(ebiten.KeySpace) {
		io.send("up")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) ||
		ebiten.IsKeyPressed(ebiten.KeyD) {
		io.send("right")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyS) {
		io.send("down")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) ||
		ebiten.IsKeyPressed(ebiten.KeyA) {
		io.send("left")
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustReleased(ebiten.KeyD) {
		io.send("released_right")
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustReleased(ebiten.KeyA) {
		io.send("released_left")
	}
}

func (io Io) send(action string) {
	for _, ch := range io.channels {
		ch <- action
	}
}
