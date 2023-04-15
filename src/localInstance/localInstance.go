package localInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/localInstance/gameWindow"
	"gopherLand2/src/localInstance/io"
)

type LocalInstance struct {
	Io        io.Io
	LocalGame game.Game
}

func StartInstance() {
	localChannel := make(chan string)

	instance := LocalInstance{
		Io:        io.New(localChannel),
		LocalGame: game.New(localChannel),
	}

	go instance.LocalGame.Run()
	go instance.LocalGame.RunPlayer()

	gameWindow.OpenWindow(instance.Io, instance.LocalGame)

}
