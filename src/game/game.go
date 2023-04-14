package game

import (
	"gopherLand2/src/game/gameMap"
	"gopherLand2/src/game/ressources"
)

type Game struct {
	Ressources ressources.Ressources
	GameMap    gameMap.GameMap
	Config     Config
}

func New() Game {
	// Load game's config
	cfg := loadConfig()

	// Create Game
	g := Game{
		Ressources: ressources.New(cfg.Size),
		GameMap:    gameMap.New(),
		Config:     cfg,
	}

	return g
}
