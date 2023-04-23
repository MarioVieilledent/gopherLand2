package serverInstance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopherLand2/src/game/entity"
	"net"
	"os"
)

const TYPE string = "tcp"

func (si *ServerInstance) startTCPserver(host, port string) {
	listen, err := net.Listen(TYPE, host+":"+port)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on " + host + ":" + port)

	for {
		conn, err := listen.Accept()
		fmt.Println("New client connected")

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// defer conn.Close()

		go si.handleConnection(conn)
	}

}

func (si *ServerInstance) handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			conn.Close()
			break
		}
		command := buf[0]
		data := bytes.Trim(buf[1:], "\x00")
		switch command {
		case byte('0'):
			{
				si.AddPlayer(string(data), entity.PlayerInfo{}, &conn)
				break
			}
		case byte('1'):
			{
				pi, err := entity.Parse(data)
				if err != nil {
					fmt.Println("Cannot parse player's position: " + string(data) + " - Error: " + err.Error())
				} else {
					_, ok := si.PlayerInfos[pi.Nickname]
					if ok {
						si.PlayerInfos[pi.Nickname] = pi
					} else {
						fmt.Println("No player named " + pi.Nickname)
					}
				}
				si.sendToAll(pi.Nickname)
				break
			}
		default:
			{
				fmt.Println("Unsupported TCP command")
				break
			}
		}
	}
}

func (si *ServerInstance) sendToAll(except string) {
	playerInfosBytes, err := json.Marshal(si.PlayerInfos)
	if err != nil {
		fmt.Println("Cannot encode players' data into json.")
	} else {
		for nicknameConn, conn := range si.PlayersConnections {
			if nicknameConn != except {
				(*conn).Write(playerInfosBytes)
			}
		}
	}
}
