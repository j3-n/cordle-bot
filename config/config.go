package config

import (
	"encoding/json"
	"os"
	"log"
	"cordle/util"
)

type Config struct {
	Token string
}

func LoadConfig(path string) (Config){
	// Open the configuration file
	file, err := os.ReadFile(path)
	util.CheckError("Failed to read config file")

	// Create a new Config struct to return
	ret := Config{}
	// Decode JSON into the struct
	err = json.Unmarshal(file, &ret)
	util.CheckError("Failed to decode JSON from config file")

	return ret
}