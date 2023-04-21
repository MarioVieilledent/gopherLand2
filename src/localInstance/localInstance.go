package localInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/TCPClient"
	"gopherLand2/src/localInstance/gameWindow"
	"gopherLand2/src/localInstance/io"
)

// Game solo
func StartInstance(nickname, character string) {
	localChannel := make(chan string)

	game := game.New(localChannel)
	game.SetPlayer(entity.Pos{X: 7.0, Y: -1.0}, nickname, character)

	io := io.New(localChannel)

	go game.Run()
	go game.RunPlayer()

	gameWindow.OpenWindow(io, &game)
}

// Game multiplayer
func ConnectToServer(host, port, nickname, character string) {
	localChannel := make(chan string)                // Send user's input to game to move own player
	playerPosChannel := make(chan entity.PlayerInfo) // Send own player position to TCPClient
	allPlayersPosChannel := make(chan []byte)        // Send all players data from TCPClient to game instance

	game := game.New(localChannel)
	game.SetPlayer(entity.Pos{X: 7.0, Y: -1.0}, nickname, character)
	game.SetPlayerNickname(nickname)

	game.BindMultiplayerChannels(playerPosChannel, allPlayersPosChannel)

	io := io.New(localChannel)

	go game.Run()
	go game.RunPlayer()

	go TCPClient.StartTCPClient(host, port, nickname, playerPosChannel, allPlayersPosChannel)

	gameWindow.OpenWindow(io, &game)
}
