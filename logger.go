package main

import "fmt"

type simpleLogger interface {
	log(message string)
}

type consoleLogger struct {
}

func (c *consoleLogger) log(message string) {
	fmt.Println(message)
}
