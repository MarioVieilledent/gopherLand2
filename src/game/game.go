package game

import (
	"gopherLand2/src/game/entity"
	"gopherLand2/src/game/gameMap"
	"gopherLand2/src/game/ressources"
	"time"
)

type Game struct {
	Config     Config                // The game's config
	Ressources ressources.Ressources // All single instance ressources
	GameMap    gameMap.GameMap       // A map
	Player     entity.Player         // Playable player
	Channel    chan string           // DEBUG controls the player[0]
	Tick       int                   // Tick of the game
	TickMS     int                   // Delay between each tick in ms

	MultiplayerChannel chan entity.Pos // Channel to send if multiplayer
}

// Create a new game with a channel that recieve players' actions
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
func (g *Game) SetPlayer(playerPos entity.Pos) {
	g.Player = entity.NewPlayer(playerPos)
}

// Bind a channel for sending to multiplayer server player's data
func (g *Game) BindMultiplayerChannel(ch chan entity.Pos) {
	g.MultiplayerChannel = ch
}
