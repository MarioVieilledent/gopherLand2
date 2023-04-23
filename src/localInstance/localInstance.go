package localInstance

import (
	"gopherLand2/src/game"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/localInstance/TCPClient"
	"gopherLand2/src/localInstance/gameWindow"
	"gopherLand2/src/localInstance/input"
)

// Game solo
func StartInstance(nickname, character string) {
	playerInputChannel := make(chan input.KeyPressed)

	game := game.New(playerInputChannel)
	game.SetPlayer(entity.Pos{X: 7.0, Y: -1.0}, nickname, character)

	input := input.New(playerInputChannel)

	go game.Run()
	go game.RunPlayer()

	gameWindow.OpenWindow(input, &game)
}

// Game multiplayer
func ConnectToServer(host, port, nickname, character string) {
	playerInputChannel := make(chan input.KeyPressed) // Send user's input to game to move own player
	playerPosChannel := make(chan entity.PlayerInfo)  // Send own player position to TCPClient
	allPlayerInfosChannel := make(chan []byte)        // Send all players data from TCPClient to game instance

	game := game.New(playerInputChannel)
	game.SetPlayer(entity.Pos{X: 7.0, Y: -1.0}, nickname, character)
	game.SetPlayerNickname(nickname)

	game.BindMultiplayerChannels(playerPosChannel, allPlayerInfosChannel)

	input := input.New(playerInputChannel)

	go game.Run()
	go game.RunPlayer()

	go TCPClient.StartTCPClient(host, port, nickname, playerPosChannel, allPlayerInfosChannel)

	gameWindow.OpenWindow(input, &game)
}
