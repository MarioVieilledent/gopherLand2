package main

import (
	"fmt"
	"gopherLand2/src/localInstance"
	"gopherLand2/src/serverInstance"
	"os"
)

func main() {
	args := os.Args

	for k, a := range args {
		fmt.Println(k, a)
	}

	if len(args) >= 3 && args[1] == "server" {
		host := args[2]
		port := args[3]
		serverInstance.StartInstance(host, port)
	} else if len(args) >= 6 && args[1] == "multiplayer" {
		host := args[2]
		port := args[3]
		nickname := args[4]
		character := args[5]
		localInstance.ConnectToServer(host, port, nickname, character)
	} else if len(args) >= 2 {
		character := args[1]
		localInstance.StartInstance("solo", character)
	} else if len(args) == 1 {
		localInstance.StartInstance("solo", "gopher")
	} else {
		fmt.Println("Wrong arguments provided.")
	}
}
