package serverInstance

import (
	"gopherLand2/src/game"
)

type Serverinstance struct {
	LocalGame game.Game
}

func StartInstance() {
	localChannel := make(chan string)

	instance := Serverinstance{
		LocalGame: game.New(localChannel),
	}

	instance.LocalGame.Run()

}
