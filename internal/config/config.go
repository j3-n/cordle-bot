package config

import (
	"cordle/internal/pkg/util"
	"encoding/json"
	"os"
)

// ConfigData structs store configuration information after they are read
type ConfigData struct {
	Discord  DiscordConfig
	Game     GameConfig
	Database DatabaseConfig
}

// Default config path
const DEFAULT_CONFIG string = "config/config.json"

// Globally available ConfigData instance
var Config ConfigData

// Called when the module is imported and automatically loads the config
func init() {
	LoadConfig(DEFAULT_CONFIG)
}

// Loads the global config, needs to be called when the program starts
func LoadConfig(path string) {
	Config = loadConfigFromFile(path)
}

// loadConfigFromFile loads a ConfigData struct from a given JSON file
func loadConfigFromFile(path string) ConfigData {
	f := loadFile(path)
	// Decode JSON into a new strut
	var c ConfigData
	err := json.Unmarshal(f, &c)
	util.CheckErrMsg(err, "Failed to decode JSON from config file")

	return c
}

func loadFile(p string) []byte {
	// Open the configuration file
	file, err := os.ReadFile(p)
	util.CheckErrMsg(err, "Failed to read config file")
	return file
}
