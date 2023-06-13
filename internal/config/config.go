package config

import (
	"cordle/internal/pkg/util"
	"encoding/json"
	"flag"
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

var configPath = flag.String("config", DEFAULT_CONFIG_PATH, "path to a JSON config file to use")

func init() {
	flag.Parse()
	// Load config from the determined path
	loadConfig(*configPath)
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
