package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var queueConfigs []*queueConfig
var baseURL string
var transportConfig http.RoundTripper

func startClient() error {
	configClient()
	if len(queueConfigs) == 0 {
		queueConfigs = []*queueConfig{}
	}

	client := http.Client{Transport: transportConfig}
	for true {

		for _, config := range queueConfigs {
			url := baseURL + "/retrieve?" + config.id
			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println(err)
			}
			request.Header.Add("x-api-key", config.key)

			response, err := client.Do(request)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if response.StatusCode == 200 {
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					fmt.Println(err)
					continue
				}
				config.handler(string(body))
			}
		}

		time.Sleep(time.Second)
	}
	return nil
}

type queueConfig struct {
	id      string
	key     string
	handler func(response string)
}
