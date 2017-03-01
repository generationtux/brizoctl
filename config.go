package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
)

type brizoctlConfig struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

// ConfigPath contains file path to brizo config file
var ConfigPath = ""

// Config contains the parsed config struct
var Config = &brizoctlConfig{}

// parseConfig will ensure the config file exists, and contains required properties
func parseConfig() error {
	// expand home directory
	actualPath, err := homedir.Expand(ConfigPath)
	if err != nil {
		return errors.New("Unable to parse config path: " + ConfigPath)
	}

	// validate file exists
	if _, err = os.Stat(actualPath); os.IsNotExist(err) {
		return errors.New("Config file does not exist at " + ConfigPath)
	}

	// read file to struct
	data, err := ioutil.ReadFile(actualPath)
	if err != nil {
		return errors.New("Unable to read config file: " + ConfigPath)
	}
	c := new(brizoctlConfig)
	err = json.Unmarshal(data, c)
	if err != nil {
		return errors.New("Unable to read JSON in config file " + ConfigPath)
	}

	// validate required properties
	if c.Endpoint == "" || c.Token == "" {
		return errors.New("Config file does not contain required properties (endpoint, token).")
	}

	// set config
	Config = c
	return nil
}
