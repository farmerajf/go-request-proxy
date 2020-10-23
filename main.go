package main

import (
	"bytes"
	"fmt"
	"net/http"
)

var requests = []string{}
var logger simpleLogger

func main() {
	logger = &consoleLogger{}

	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)

	http.HandleFunc("/request", handleRequest)
	http.HandleFunc("/retrieve", handleRetrieve)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	requests = append(requests, body)

	fmt.Fprintf(w, "OK")
}

func handleRetrieve(w http.ResponseWriter, _ *http.Request) {
	r, nr := requests[0], requests[1:]
	requests = nr

	fmt.Fprintf(w, r)
}
