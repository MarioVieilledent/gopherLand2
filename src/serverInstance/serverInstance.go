package serverInstance

import (
	"gopherLand2/src/game/entity"
	"net"
)

type ServerInstance struct {
	PlayersPositions   map[string]entity.Pos
	PlayersConnections map[string]*net.Conn
}

func (si *ServerInstance) AddPlayer(nickname string, playerPos entity.Pos, conn *net.Conn) {
	si.PlayersPositions[nickname] = playerPos
	si.PlayersConnections[nickname] = conn
}

func StartInstance() {
	serverInstance := ServerInstance{
		PlayersPositions:   map[string]entity.Pos{},
		PlayersConnections: map[string]*net.Conn{},
	}

	serverInstance.startTCPserver()
}
