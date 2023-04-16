package localInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/TCPClient"
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
		Io:        io.New([]chan string{localChannel}),
		LocalGame: game,
	}

	go instance.LocalGame.Run()
	go instance.LocalGame.RunPlayer(0)

	gameWindow.OpenWindow(instance.Io, instance.LocalGame)
}

func ConnectToServer(host, port string) {
	localChannel := make(chan string)
	TCPChannel := make(chan string)

	game := game.New(localChannel)
	game.AddPlayer(entity.Pos{X: 7, Y: -1})

	instance := LocalInstance{
		Io:        io.New([]chan string{localChannel, TCPChannel}),
		LocalGame: game,
	}

	go instance.LocalGame.Run()
	go instance.LocalGame.RunPlayer(0)

	go TCPClient.StartTCPClient(host, port, TCPChannel)

	gameWindow.OpenWindow(instance.Io, instance.LocalGame)
}
