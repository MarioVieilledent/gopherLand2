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
		si.AddPlayer(PlayerConn{
			Id:   idPlayer,
			Pos:  entity.Pos{X: 0.0, Y: 0.0},
			Conn: &conn,
		})
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

		strPos := string(buf)
		pos, err := entity.ParsePos(strPos)
		if err != nil {
			fmt.Println("Cannot parse player's position: " + strPos)
		} else {
			si.PlayersConnected[idPlayer].Pos = pos
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
