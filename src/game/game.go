package game

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopherLand2/src/game/entity"
	"gopherLand2/src/game/gameMap"
	"gopherLand2/src/game/ressources"
	"time"
)

type Game struct {
	// Config
	Config Config // The game's config

	// Ressources and map
	Ressources ressources.Ressources // All single instance ressources
	GameMap    gameMap.GameMap       // A map

	// Local Player
	Player  entity.Player // Playable player
	Channel chan string   // Controls the player
	Tick    int           // Tick of the game
	TickMS  int           // Delay between each tick in ms

	// Multiplayer data
	PlayerPos map[string]entity.Pos

	// Channels for multiplayer
	PlayerPosChannel     chan entity.PlayerInfo // Channel for sending own player position if multiplayer
	AllPlayersPosChannel chan []byte            // Channel for receiving all players positions if multiplayer
}

// Create a new game with a channel that receive players' actions
func New(ch chan string) Game {
	// Load game's config
	cfg := loadConfig()

	// Create Game
	g := Game{
		Config:     cfg,
		Ressources: ressources.New(cfg.Size),
		GameMap:    gameMap.New(),
		Player:     entity.Player{},
		Channel:    ch,
		Tick:       0,
		TickMS:     cfg.TickMS,
	}

	return g
}

// Run the game
func (g *Game) Run() {
	for {
		g.Tick++
		time.Sleep(time.Duration(g.TickMS) * time.Millisecond)
		g.ComputeTick()
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
	g.Player = entity.NewPlayer(playerPos, nickname, characterKey)
}

// Sets Player's nickname for multiplayer
func (g *Game) SetPlayerNickname(nickname string) {
	g.Player.Nickname = nickname
}

// Bind a channel for sending to multiplayer server player's data
func (g *Game) BindMultiplayerChannels(playerPosChannel chan entity.PlayerInfo, allPlayersPosChannel chan []byte) {
	g.PlayerPosChannel = playerPosChannel
	g.AllPlayersPosChannel = allPlayersPosChannel
	go g.UpdateAllPlayers()
}

// Update all player in case of server send data
func (g *Game) UpdateAllPlayers() {
	for {
		data := <-g.AllPlayersPosChannel
		playersConnected := map[string]entity.Pos{}
		err := json.Unmarshal(bytes.Trim(data, string([]byte{0})), &playersConnected)
		if err != nil {
			fmt.Println("Error parsing players' data sent by server: " + err.Error())

			fmt.Println(string(data))
		}

		g.PlayerPos = playersConnected
	}
}
