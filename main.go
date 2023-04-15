package main

import (
	"fmt"
	"gopherLand2/src/localInstance"
	"gopherLand2/src/serverInstance"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		localInstance.StartInstance()
	} else if len(args) == 2 && args[1] == "server" {
		serverInstance.StartInstance()
	} else {
		fmt.Println("Wrong arguments provided.")
	}
}
