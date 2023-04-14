package gameMap

import (
	"os"
	"strings"
)

type GameMap struct {
	Name   string
	Blocks [][]string
}

func New() GameMap {
	gm := GameMap{
		Name:   "default",
		Blocks: [][]string{},
	}

	gm.loadMap("map1.txt")

	return gm
}

// Load a map from a txt file
func (gm *GameMap) loadMap(file string) {
	data, err := os.ReadFile("data/maps/" + file)

	if err != nil {
		panic(err)
	}

	for y, line := range strings.Split(string(data), "\r\n") {
		gm.Blocks = append(gm.Blocks, []string{})
		for _, ru := range line {
			str := string(ru)
			gm.Blocks[y] = append(gm.Blocks[y], str)
		}
	}
}
