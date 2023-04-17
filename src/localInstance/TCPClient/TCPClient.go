package TCPClient

import (
	"fmt"
	"gopherLand2/src/game/entity"
	"net"
	"os"
)

const TYPE string = "tcp"

func StartTCPClient(host, port string, multiplayerChannel chan entity.Pos) {
	tcpServer, err := net.ResolveTCPAddr(TYPE, host+":"+port)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		action := <-multiplayerChannel
		fmt.Println(action.ToString())

		_, err = conn.Write([]byte(action.ToString()))
		if err != nil {
			println("Write data failed:", err.Error())
			// os.Exit(1)
		}

		// buffer to get data
		received := make([]byte, 1024)
		_, err = conn.Read(received)
		if err != nil {
			println("Read data failed:", err.Error())
			// os.Exit(1)
		}

		println("Received message:", string(received))
	}
}
