package entity

import (
	"encoding/json"
	"gopherLand2/src/localInstance/input"
)

// Used for multiplayer data sharing
type PlayerInfo struct {
	Nickname   string           `json:"nickname"`
	Pos        Pos              `json:"pos"`
	KeyPressed input.KeyPressed `json:"keyPressed"`
}

func (pi PlayerInfo) Stringify() ([]byte, error) {
	data, err := json.Marshal(pi)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func Parse(data []byte) (PlayerInfo, error) {
	var pi PlayerInfo
	err := json.Unmarshal(data, &pi)
	if err != nil {
		return PlayerInfo{}, err
	}
	return pi, nil
}
