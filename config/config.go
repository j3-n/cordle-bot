package config

import (
	"encoding/json"
	"os"
	
	"cordle/util"
)

// Config structs store configuration information after they are read
type ConfigData struct {
	Token 	string
	Status 	string
}

// The path to read the config from
const configPath = "config/config.json"

// Globally available config data
var Config ConfigData

// When the module is first imported, load the config from a JSON file
func init(){
	// Open the configuration file
	file, err := os.ReadFile(configPath)
	util.CheckError(err, "Failed to read config file")

	// Decode JSON into the struct
	err = json.Unmarshal(file, &Config)
	util.CheckError(err, "Failed to decode JSON from config file")
}