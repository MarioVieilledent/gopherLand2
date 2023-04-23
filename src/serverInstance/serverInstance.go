package serverInstance

import (
	"gopherLand2/src/game/entity"
	"net"
)

type ServerInstance struct {
	PlayerInfos        map[string]entity.PlayerInfo
	PlayersConnections map[string]*net.Conn
}

func (si *ServerInstance) AddPlayer(nickname string, playerPos entity.PlayerInfo, conn *net.Conn) {
	si.PlayerInfos[nickname] = playerPos
	si.PlayersConnections[nickname] = conn
}

func StartInstance(host, port string) {
	serverInstance := ServerInstance{
		PlayerInfos:        map[string]entity.PlayerInfo{},
		PlayersConnections: map[string]*net.Conn{},
	}

	serverInstance.startTCPserver(host, port)
}
