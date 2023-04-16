package game

import (
	"encoding/json"
	"os"
)

type Config struct {
	Version string `json:"version"`
	Size    int    `json:"size"`
	TickMS  int    `json:"tickMS"`
}

func loadConfig() (cfg Config) {
	data, err := os.ReadFile("data/tables/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return
}
