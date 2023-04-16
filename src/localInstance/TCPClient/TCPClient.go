package TCPClient

import (
	"net"
	"os"
)

const TYPE string = "tcp"

func StartTCPClient(host, port string, TCPChannel chan string) {
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
		action := <-TCPChannel

		_, err = conn.Write([]byte(action))
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
