package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var queues map[string][]string
var keys map[string][]string
var logger simpleLogger

func main() {
	logger = &consoleLogger{}
	queues = make(map[string][]string)
	keys = make(map[string][]string)

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

	key := r.Header.Get("x-api-key")
	queueName := r.URL.RawQuery

	queue, exists := queues[queueName]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		logger.log("Queue " + queueName + " does not exist ")
		return
	}
	keyFound := false
	for _, configKey := range keys[queueName] {
		if key == configKey {
			keyFound = true
		}
	}
	if !keyFound {
		w.WriteHeader(http.StatusNotFound)
		logger.log("Invalid key for queue provided")
		return
	}

	queue = append(queue, body)
	queues[queueName] = queue

	logger.log("Pushed to queue " + queueName)
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("x-api-key")
	queueName := r.URL.RawQuery

	queue, exists := queues[queueName]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		logger.log("Queue " + queueName + " does not exist ")
		return
	}
	keyFound := false
	for _, configKey := range keys[queueName] {
		if key == configKey {
			keyFound = true
		}
	}
	if !keyFound {
		w.WriteHeader(http.StatusNotFound)
		logger.log("Invalid key for queue provided")
		return
	}
	if len(queue) < 1 {
		w.WriteHeader(http.StatusNotFound)
		logger.log("Attempt to retrieve from empty queue " + queueName)
		return
	}

	request, newQueue := queue[0], queue[1:]
	queues[queueName] = newQueue

	fmt.Fprintf(w, request)

	logger.log("Popped from queue " + queueName)
}

func initConfig(config *config) {
	for queueName, queue := range config.Queues {
		queues[queueName] = []string{}
		logger.log("Initialised queue " + queueName)

		keys[queueName] = []string{}
		for _, key := range queue.Keys {
			keys[queueName] = append(keys[queueName], key)
			logger.log("Initialised a key for queue " + queueName)
		}
	}
}
