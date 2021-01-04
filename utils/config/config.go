package config

import (
	"encoding/json"
	"io/ioutil"
)

type (
	ConfFile struct {
		Debug    bool
		Database ConfSQL
		Web      ConfWeb
	}
	ConfSQL struct {
		Username string
		Password string
		Host     string
		Port     int
		DB       string
	}
	ConfWeb struct {
		Port int
	}
)

//Read config
func Read() (*ConfFile, error) {
	configJSON, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var conf ConfFile
	err = json.Unmarshal([]byte(configJSON), &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
