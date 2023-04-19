package entity

import (
	"errors"
	"strings"
)

// Used for multiplayer data sharing
type PlayerInfo struct {
	Nickname string
	Pos      Pos
}

func (pi PlayerInfo) ToString() string {
	return pi.Nickname + "|" + pi.Pos.ToString()
}

func ParsePlayerInfo(s string) (PlayerInfo, error) {
	split := strings.Split(s, "|")
	if len(split) == 2 {
		pp, err := ParsePos(split[1])
		if err != nil {
			return PlayerInfo{}, err
		} else {
			pi := PlayerInfo{
				Nickname: split[0],
				Pos:      pp,
			}
			return pi, nil
		}
	} else {
		return PlayerInfo{}, errors.New("there's extra | symbols in pseudo")
	}
}
