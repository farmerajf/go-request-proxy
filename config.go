package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	Queues map[string]queue
}

type queue struct {
	ID   string
	Name string
	Keys []string
}

func loadConfig() (*config, error) {
	data, err := readConfigFile()
	if err != nil {
		return nil, err
	}

	config, err := mapConfig(data)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func readConfigFile() ([]byte, error) {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	logger.log("Read config" + string(byteValue))
	return byteValue, nil
}

func mapConfig(data []byte) (*config, error) {
	var config config

	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
