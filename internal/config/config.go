package config

import (
	"cordle/internal/pkg/util"
	"encoding/json"
	"os"
)

// Config structs store configuration information after they are read
type Config struct {
	Discord DiscordConfig
	Game    GameConfig
	Sql     SqlConfig
}

// The path to read the config from
const (
	discPath = "config/config.json"
	sqlPath  = "configs/db-key.json"
)

// Globally available config data
var Conf Config

// When the module is first imported, load the config from a JSON file
func init() {
	// Open the configuration file
	file, err := os.ReadFile(discPath)
	util.CheckErrMsg(err, "Failed to read config file")

	// Decode JSON into the struct
	var d DiscordConfig
	err = json.Unmarshal(file, &d)
	util.CheckErrMsg(err, "Failed to decode JSON from config file")

	Conf.Discord = d
}
