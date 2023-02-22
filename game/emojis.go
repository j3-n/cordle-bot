package game

import (
	"cordle/util"
	"encoding/json"
	"os"
)

// Stores a map of emoji names to their raw data
var Emojis map[string]string

func init(){
	// Load the list of emojis into a map
	f, err := os.ReadFile("game/emojis.json")
	util.CheckError(err, "Failed to load emojis.json")
	err = json.Unmarshal(f, &Emojis)
	util.CheckError(err, "Failed to decode emojis.json")
}