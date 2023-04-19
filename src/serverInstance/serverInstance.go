package serverInstance

import (
	"gopherLand2/src/game/entity"
	"net"
)

type ServerInstance struct {
	PlayersConnected map[string]PlayerConn
}

type PlayerConn struct {
	Pos  entity.Pos
	Conn *net.Conn
}

func (si *ServerInstance) AddPlayer(nickname string, player PlayerConn) {
	si.PlayersConnected[nickname] = player
}

func StartInstance() {
	serverInstance := ServerInstance{
		PlayersConnected: map[string]PlayerConn{},
	}

	serverInstance.startTCPserver()
}
