package localInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/TCPClient"
	"gopherLand2/src/localInstance/gameWindow"
	"gopherLand2/src/localInstance/io"
)

// Game solo
func StartInstance() {
	localChannel := make(chan string)

	game := game.New(localChannel)
	game.SetPlayer(entity.Pos{X: 7.0, Y: -1.0})

	io := io.New(localChannel)

	go game.Run()
	go game.RunPlayer()

	gameWindow.OpenWindow(io, game)
}

// Game multiplayer
func ConnectToServer(host, port string) {
	localChannel := make(chan string)
	multiplayerChannel := make(chan entity.Pos)

	game := game.New(localChannel)
	game.SetPlayer(entity.Pos{X: 7.0, Y: -1.0})
	game.BindMultiplayerChannel(multiplayerChannel)

	io := io.New(localChannel)

	go game.Run()
	go game.RunPlayer()

	go TCPClient.StartTCPClient(host, port, multiplayerChannel)

	gameWindow.OpenWindow(io, game)
}
