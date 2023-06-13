package config

import (
	"cordle/internal/pkg/util"
	"encoding/json"
	"log"
	"os"
)

// ConfigData structs store configuration information after they are read
type ConfigData struct {
	Discord  DiscordConfig  `json:"discord"`
	Game     GameConfig     `json:"game"`
	Database DatabaseConfig `json:"database"`
}

// Default config path
const DEFAULT_CONFIG_PATH string = "config/config.json"

// Globally available ConfigData instance
var Config ConfigData

func init() {
	// Check if the config path has been overwritten
	p := os.Getenv("CORDLE_CONFIG_PATH")
	if p == "" {
		p = DEFAULT_CONFIG_PATH
	} else {
		// Use debug logs to indicate that an alternative config is being used
		log.Printf("Using alternate config (%s)", p)
	}
	// Load config from the determined path
	loadConfig(p)
}

// LoadConfig initialises the global Config instance. This MUST be called at the start of the program
func loadConfig(path string) {
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
