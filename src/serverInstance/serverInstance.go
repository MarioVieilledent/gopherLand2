package serverInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/serverInstance/server"
)

type Serverinstance struct {
	LocalGame game.Game
}

func StartInstance() {
	localChannel := make(chan string)

	instance := Serverinstance{
		LocalGame: game.New(localChannel),
	}

	go instance.LocalGame.Run()

	server.StartTCPserver()
}
