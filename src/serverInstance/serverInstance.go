package serverInstance

import (
	"gopherLand2/src/game/entity"
	"net"
)

type Serverinstance struct {
	PlayersConnected []PlayerConn
}

type PlayerConn struct {
	Id   int
	Pos  entity.Pos
	Conn *net.Conn
}

func (si *Serverinstance) AddPlayer(player PlayerConn) {
	si.PlayersConnected = append(si.PlayersConnected, player)
}

func StartInstance() {
	serverInstance := Serverinstance{
		PlayersConnected: []PlayerConn{},
	}

	serverInstance.startTCPserver()
}
