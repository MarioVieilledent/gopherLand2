package instance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/instance/gameWindow"
	"gopherLand2/src/instance/io"
)

type Instance struct {
	Io        io.Io
	LocalGame game.Game
}

func StartInstance() {
	localChannel := make(chan string)

	instance := Instance{
		Io:        io.New(localChannel),
		LocalGame: game.New(localChannel),
	}

	go instance.LocalGame.Run()
	go instance.LocalGame.RunPlayer()

	gameWindow.OpenWindow(instance.Io, instance.LocalGame)
}
