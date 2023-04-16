package server

import (
	"fmt"
	"net"
	"os"
)

const CONN_HOST string = "0.0.0.0"
const CONN_PORT string = "12345"
const CONN_TYPE string = "tcp"

func StartTCPserver() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// defer conn.Close()

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		mes, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			conn.Close()
			break
		}
		fmt.Println("recu : " + string(mes))

		conn.Write([]byte("ok"))
	}
}
