package serverInstance

import (
	"encoding/json"
	"fmt"
	"gopherLand2/src/game/entity"
	"net"
	"os"
)

const HOST string = "0.0.0.0"
const PORT string = "12387"
const TYPE string = "tcp"

func (si *Serverinstance) startTCPserver() {
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

func (si Serverinstance) handleConnection(conn net.Conn, idPlayer int) {
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			conn.Close()
			break
		}

		data := string(buf)
		fmt.Println(data)
		switch data[0] {
		case '0':
			{
				fmt.Println(si)

				si.AddPlayer(data[1:], PlayerConn{
					Pos:  entity.Pos{X: 0, Y: 0},
					Conn: &conn,
				})
				fmt.Println(si)
				break
			}
		default:
			{
				strPlayerInfo := string(buf)
				pi, err := entity.ParsePlayerInfo(strPlayerInfo)
				if err != nil {
					fmt.Println("Cannot parse player's position: " + strPlayerInfo)
				} else {
					pc, ok := si.PlayersConnected[pi.Nickname]
					if ok {
						pc.Pos = pi.Pos
					} else {
						fmt.Println("No player named " + pi.Nickname)
					}
				}
			}
			break
		}

		go si.sendToAll()
	}
}

func (si Serverinstance) sendToAll() {
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
