package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type config struct {
	Queues map[string]*queue
}

type queue struct {
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

func writeConfigFile(config *config) error {
	json, err := json.Marshal(config)
	if err != nil {
		return err
	}

	ioutil.WriteFile("config.json", json, 0)
	return nil
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

func addQueue(name string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", err
	}

	queueName := generateString()
	config.Queues[queueName] = &queue{Name: name}

	err = writeConfigFile(config)
	if err != nil {
		return "", err
	}

	return queueName, nil
}

func removeQueue(queueName string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", err
	}

	delete(config.Queues, queueName)

	err = writeConfigFile(config)
	if err != nil {
		return "", err
	}

	return queueName, nil
}

func addKey(queueName string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", err
	}

	queue, exists := config.Queues[queueName]
	if !exists {
		return "", errors.New("invalid queue")
	}
	key := generateString()
	keys := queue.Keys
	keys = append(keys, key)
	//config.Queues[queueName].Keys = keys
	queue.Keys = keys

	err = writeConfigFile(config)
	if err != nil {
		return "", err
	}

	return key, nil
}

func removeKey(queueName string, key string) (string, error) {
	config, err := loadConfig()
	if err != nil {
		return "", err
	}

	oldKeys := config.Queues[queueName].Keys
	newKeys := []string{}
	for _, currentKey := range oldKeys {
		if key != currentKey {
			newKeys = append(newKeys, currentKey)
		}
	}
	config.Queues[queueName].Keys = newKeys

	err = writeConfigFile(config)
	if err != nil {
		return "", err
	}

	return key, nil

}

func generateString() string {
	charset := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 60)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
