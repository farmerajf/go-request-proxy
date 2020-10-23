package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var queues map[string][]string
var logger simpleLogger

func main() {
	logger = &consoleLogger{}
	queues = make(map[string][]string)

	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	initConfig(config)

	http.HandleFunc("/request", handleRequest)
	http.HandleFunc("/retrieve", handleRetrieve)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	queueName := r.URL.RawQuery

	logger.log("New request for queue " + queueName)

	queue, exists := queues[queueName]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		logger.log("Queue " + queueName + " does not exist ")

		return
	}
	queue = append(queue, body)
	queues[queueName] = queue

	logger.log("Added request to queue " + body)
}

func handleRetrieve(w http.ResponseWriter, _ *http.Request) {
	//r, nr := requests[0], requests[1:]
	//requests = nr

	//fmt.Fprintf(w, r)
}

func initConfig(config *config) {
	for queueName, _ := range config.Queues {
		queues[queueName] = []string{}
		logger.log("Initialised queue " + queueName)
	}
}
