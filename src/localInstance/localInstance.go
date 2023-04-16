package localInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/gameWindow"
	"gopherLand2/src/localInstance/io"
)

type LocalInstance struct {
	Io        io.Io
	LocalGame game.Game
}

func StartInstance() {
	localChannel := make(chan string)

	game := game.New(localChannel)
	game.AddPlayer(entity.Pos{X: 7, Y: -1})

	instance := LocalInstance{
		Io:        io.New(localChannel),
		LocalGame: game,
	}

	go instance.LocalGame.Run()
	go instance.LocalGame.RunPlayer(0)

	gameWindow.OpenWindow(instance.Io, instance.LocalGame)
}

func ConnectToServer(host string) {
	localChannel := make(chan string)

	game := game.New(localChannel)
	game.AddPlayer(entity.Pos{X: 7, Y: -1})

	instance := LocalInstance{
		Io:        io.New(localChannel),
		LocalGame: game,
	}

	go instance.LocalGame.Run()
	go instance.LocalGame.RunPlayer(0)

	gameWindow.OpenWindow(instance.Io, instance.LocalGame)
}
