package TCPClient

import (
	"fmt"
	"gopherLand2/src/game/entity"
	"net"
	"os"
)

const TYPE string = "tcp"

func StartTCPClient(host, port, nickname string, playerPosChannel chan entity.PlayerInfo, allPlayerInfosChannel chan []byte) {
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

	// Send nickname to server
	_, err = conn.Write([]byte("0" + nickname))
	if err != nil {
		fmt.Println("Cannot send nickname to server:", err.Error())
		os.Exit(1)
	}
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		fmt.Println("Read data failed:", err.Error())
		os.Exit(1)
	} else {
		allPlayerInfosChannel <- received
	}

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
				allPlayerInfosChannel <- received
			}
		}
	}()

	// Send own player info
	for {
		playerInfo := <-playerPosChannel

		pi, err := playerInfo.Stringify()
		if err != nil {
			fmt.Println("Cannot parse playerInfo into JSON: " + err.Error())
		} else {
			_, err = conn.Write([]byte("1" + string(pi)))
			if err != nil {
				fmt.Println("Write data failed:", err.Error())
				os.Exit(1)
			}
		}
	}
}
