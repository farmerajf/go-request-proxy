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

	var result string
	var err error
	switch command {
	case "start-server":
		startServer()
	case "start-client":
		fmt.Println(startClient())
	case "add-queue":
		result, err = addQueue(arg)
	case "add-key":
		result, err = addKey(arg)
	case "remove-queue":
		result, err = removeQueue(arg)
	case "remove-key":
		result, err = removeQueue(arg)
	default:
		fmt.Println("Unknown command")
		help()
	}

	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
}

func help() {
	fmt.Println("Usage: command")
	fmt.Println("commands")
	fmt.Println("	start-server")
	fmt.Println("	start-client")
	fmt.Println("	add-queue name")
	fmt.Println("	add-key name")
	fmt.Println("	remove-queue name")
	fmt.Println("	remove-key queue-name key-name")
}
