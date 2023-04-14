package io

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		io.channel <- "up"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		io.channel <- "right"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		io.channel <- "down"
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		io.channel <- "left"
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		io.channel <- "space"
	}
}
