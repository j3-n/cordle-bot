package config

import (
	"cordle/internal/pkg/util"
	"encoding/json"
	"os"
)

// ConfigData structs store configuration information after they are read
type ConfigData struct {
	Discord DiscordConfig
	Game    GameConfig
	Sql     SqlConfig
}

// The path to read the config from
const (
	discPath = "config/discord-config.json"
	gamePath = "config/game-config.json"
	sqlPath  = "configs/db-key.json"
)

// Globally available config data
var Config ConfigData

// When the module is first imported, load the config from a JSON file
func init() {
	loadDiscordConfig()
	loadGameConfig()
}

func loadDiscordConfig() {
	file := loadFile(discPath)

	// Decode JSON into the struct
	var d DiscordConfig
	err := json.Unmarshal(file, &d)
	util.CheckErrMsg(err, "Failed to decode JSON from discord config file")

	Config.Discord = d
}

func loadGameConfig() {
	file := loadFile(gamePath)
	// Decode JSON into the struct
	var g GameConfig
	err := json.Unmarshal(file, &g)
	util.CheckErrMsg(err, "Failed to decode JSON from game config file")

	Config.Game = g
}

func loadFile(p string) []byte {
	// Open the configuration file
	file, err := os.ReadFile(p)
	util.CheckErrMsg(err, "Failed to read config file")
	return file
}
