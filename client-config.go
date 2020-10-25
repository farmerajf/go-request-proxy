package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func configClient() {
	baseURL = "https://localhost:8080"
	transportConfig = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	queueConfigs = []*queueConfig{
		{
			id:  "tsGrwks8CS5kCG5ZKjbZ8dUQduylVyl1r8s5buaiZpbbt8O6J8E9hXxgeLnC",
			key: "clhLTI2en4XioOgoiZZbMB6ZILUvPSqvQ1L2QoFjk5fIWLXRiJQgFhlTimf0",
			handler: func(response string) {
				fmt.Println(response)
			},
		},
	}
}
