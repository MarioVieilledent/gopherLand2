package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/game/gameMap"
	"gopherLand2/src/game/ressources"
	"gopherLand2/src/localInstance/input"
)

type Game struct {
	// Config
	Config Config // The game's config

	// Ressources and map
	Ressources ressources.Ressources // All single instance ressources
	GameMap    gameMap.GameMap       // A map

	// Players
	Nickname string                   // Nickname of playable player
	Players  map[string]entity.Player // Other players

	// Game tick
	Tick   int // Tick of the game
	TickMS int // Delay between each tick in ms

	// PlayerInputChannel for getting input of player
	PlayerInputChannel chan input.KeyPressed // Controls the player

	// Channels for multiplayer
	PlayerPosChannel      chan entity.PlayerInfo // Channel for sending own player position if multiplayer
	AllPlayerInfosChannel chan []byte            // Channel for receiving all players positions if multiplayer
}

// Create a new game with a channel that receive players' actions
func New(ch chan input.KeyPressed, nickname string) Game {
	// Load game's config
	cfg := loadConfig()

	// Create Game
	g := Game{
		Config:             cfg,
		Ressources:         ressources.New(cfg.Size),
		GameMap:            gameMap.New(),
		Nickname:           nickname,
		Players:            map[string]entity.Player{},
		PlayerInputChannel: ch,
		Tick:               0,
		TickMS:             cfg.TickMS,
	}

	return g
}

// Run the game, function called by ebiten's Update() function (in file gameWindow.go)
func (g *Game) Run() {
	g.Tick++
	for nickname := range g.Players {
		g.ManagePlayerMovement(nickname)
	}
}

// Add a new player
func (g *Game) SetPlayer(playerPos entity.Pos, nickname, character string) {
	characterKey := "k"
	// If name is found among ressource elements,
	// sets player's Character property to key in map g.Ressources.Elements
	for k, v := range g.Ressources.Elements {
		if v.Name == "player_"+character {
			characterKey = k
		}
	}
	g.Players[g.Nickname] = entity.NewPlayer(playerPos, nickname, characterKey)
}

// Bind a channel for sending to multiplayer server player's data
func (g *Game) BindMultiplayerChannels(playerPosChannel chan entity.PlayerInfo, allPlayerInfosChannel chan []byte) {
	g.PlayerPosChannel = playerPosChannel
	g.AllPlayerInfosChannel = allPlayerInfosChannel
	go g.UpdateAllPlayers()
}

// Update all player in case of server send data
func (g *Game) UpdateAllPlayers() {
	for {
		data := <-g.AllPlayerInfosChannel
		playersConnected := map[string]entity.PlayerInfo{}
		err := json.Unmarshal(bytes.Trim(data, string([]byte{0})), &playersConnected)
		if err != nil {
			fmt.Println("Error parsing players' data sent by server: " + string(data) + " - Error:" + err.Error())
		}

		for nickname, pi := range playersConnected {
			if _, ok := g.Players[nickname]; ok {
				g.PlayerInputChannel <- pi.KeyPressed
			} else {
				g.Players[nickname] = entity.NewPlayer(pi.Pos, pi.Nickname, pi.Character)
				g.PlayerInputChannel <- pi.KeyPressed
			}
		}
	}
}
