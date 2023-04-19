package serverInstance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopherLand2/src/game/entity"
	"net"
	"os"
)

const HOST string = "0.0.0.0"
const PORT string = "12387"
const TYPE string = "tcp"

func (si *ServerInstance) startTCPserver() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on " + HOST + ":" + PORT)

	for {
		idPlayer := len(si.PlayersConnected)
		conn, err := listen.Accept()
		fmt.Println("New client connected")

		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// defer conn.Close()

		go si.handleConnection(conn, idPlayer)
	}

}

func (si *ServerInstance) handleConnection(conn net.Conn, idPlayer int) {
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			conn.Close()
			break
		}
		command := buf[0]
		data := string(bytes.Trim(buf[1:], "\x00"))
		switch command {
		case byte('0'):
			{
				si.AddPlayer(data, PlayerConn{
					Pos:  entity.Pos{X: 0, Y: 0},
					Conn: &conn,
				})
				break
			}
		case byte('1'):
			{
				pi, err := entity.ParsePlayerInfo(data)
				if err != nil {
					fmt.Println("Cannot parse player's position: " + data)
				} else {
					pc, ok := si.PlayersConnected[pi.Nickname]
					if ok {
						pc.Pos = pi.Pos
					} else {
						fmt.Println("No player named " + pi.Nickname)
					}
				}
				break
			}
		default:
			{
				fmt.Println("Unsupported TCP command")
				break
			}
		}

		go si.sendToAll()
	}
}

func (si *ServerInstance) sendToAll() {
	var list []entity.Pos
	for _, pc := range si.PlayersConnected {
		list = append(list, pc.Pos)
	}
	for _, pc := range si.PlayersConnected {
		playerPosStr, err := json.Marshal(list)
		if err != nil {
			fmt.Println("Cannot encode players' data into json.")
		} else {
			(*pc.Conn).Write(playerPosStr)
		}
	}
}
