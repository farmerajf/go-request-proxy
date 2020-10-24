package main

import (
	"fmt"
	"os"
)

var logger simpleLogger

func main() {
	logger = &consoleLogger{}

	if len(os.Args) < 2 {
		help()
		return
	}

	command := os.Args[1]

	arg := ""
	if len(os.Args) > 2 {
		arg = os.Args[2]
	}

	switch command {
	case "start-server":
		startServer()
	case "create-queue":
		fmt.Println(createQueue(arg))
	case "add-key":
		fmt.Println(createKey(arg))
	default:
		fmt.Println("Unknown command")
		help()
	}
}

func help() {
	fmt.Println("Usage: command")
	fmt.Println("commands")
	fmt.Println("	start-server")
	fmt.Println("	start-client")
}
