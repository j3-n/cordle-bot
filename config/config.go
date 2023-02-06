package config

import (
	"encoding/json"
	"os"
	"cordle/util"
)

type Config struct {
	Token string
}

func LoadConfig(path string) (Config){
	// Open the configuration file
	file, err := os.ReadFile(path)
	util.CheckError(err, "Failed to read config file")

	// Create a new Config struct to return
	ret := Config{}
	// Decode JSON into the struct
	err = json.Unmarshal(file, &ret)
	util.CheckError(err, "Failed to decode JSON from config file")

	return ret
}