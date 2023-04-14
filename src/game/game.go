package game

import (
	"gopherLand2/src/game/entity"
	"gopherLand2/src/game/gameMap"
	"gopherLand2/src/game/ressources"
	"time"
)

type Game struct {
	Config     Config                // A config
	Ressources ressources.Ressources // All single instance ressources
	GameMap    gameMap.GameMap       // A map
	Player     []entity.Player       // A list of players
	Channel    chan string           // DEBUG controls the player[0]
	Tick       int                   // Tick of the game
}

func New(ch chan string) Game {
	// Load game's config
	cfg := loadConfig()

	// Create Game
	g := Game{
		Config:     cfg,
		Ressources: ressources.New(cfg.Size),
		GameMap:    gameMap.New(),
		Player: []entity.Player{
			entity.NewPlayer(entity.Pos{X: 1.0, Y: 1.0}),
		},
		Channel: ch,
		Tick:    0,
	}

	return g
}

// Run the game
func (g *Game) Run() {
	for {
		g.Tick++
		time.Sleep(100 * time.Millisecond)
	}
}

func (g *Game) RunPlayer() {
	var action string

	for {
		action = <-g.Channel
		if action == "left" {
			g.Player[0].Pos.X -= 0.5
		}
		if action == "right" {
			g.Player[0].Pos.X += 0.5
		}
	}
}
