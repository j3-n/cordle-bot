package config

import (
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	Token string
}

func LoadConfig(path string) (Config){
	// Open the configuration file
	file, err := os.ReadFile(path)
	if err != nil{
		log.Fatal("Failed to read config file")
	}

	// Create a new Config struct to return
	ret := Config{}
	// Decode JSON into the struct
	err = json.Unmarshal(file, &ret)
	if err != nil{
		log.Fatal("Failed to decode JSON from config file")
	}

	return ret
}