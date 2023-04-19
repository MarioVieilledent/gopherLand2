package serverInstance

import (
	"gopherLand2/src/game/entity"
	"net"
)

type Serverinstance struct {
	PlayersConnected map[string]PlayerConn
}

type PlayerConn struct {
	Pos  entity.Pos
	Conn *net.Conn
}

func (si *Serverinstance) AddPlayer(nickname string, player PlayerConn) {
	si.PlayersConnected[nickname] = player
}

func StartInstance() {
	serverInstance := Serverinstance{
		PlayersConnected: map[string]PlayerConn{},
	}

	serverInstance.startTCPserver()
}
