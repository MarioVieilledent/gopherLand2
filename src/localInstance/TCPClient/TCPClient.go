package TCPClient

import (
	"fmt"
	"gopherLand2/src/game/entity"
	"net"
	"os"
)

const TYPE string = "tcp"

func StartTCPClient(host, port string, playerPosChannel chan entity.Pos, allPlayersPosChannel chan []byte) {
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

	fmt.Println("Connected to server " + host + ":" + port)

	// Receive data from server
	go func() {
		for {
			received := make([]byte, 1024)
			_, err = conn.Read(received)
			if err != nil {
				fmt.Println("Read data failed:", err.Error())
				os.Exit(1)
			} else {
				allPlayersPosChannel <- received
			}
		}
	}()

	// Send own player position
	for {
		action := <-playerPosChannel

		_, err = conn.Write([]byte(action.ToString()))
		if err != nil {
			fmt.Println("Write data failed:", err.Error())
			os.Exit(1)
		}
	}
}
