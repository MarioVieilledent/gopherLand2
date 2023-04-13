package game

import (
	"gopherLand2/src/game/gameMap"
	"gopherLand2/src/game/ressources"
)

type Game struct {
	Ressources ressources.Ressources
	GameMap    gameMap.GameMap
}

func New() Game {
	g := Game{
		Ressources: ressources.New(),
		GameMap:    gameMap.New(),
	}

	return g
}
