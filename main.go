package main

import (
	"time" // Temporary

	"cordle/config"
	"cordle/util"

	"github.com/bwmarrin/discordgo"
)

// Path to read config from
const ConfigPath = "config/config.json"
// Path to read command metadata from
const CommandsPath string = "commands/commands.json"

func main() {
	// Load config file
	config := config.LoadConfig(ConfigPath)

	// Start discord bot
	session, err := discordgo.New("Bot " + config.Token)
	util.CheckError(err, "Failed to initialise discord session")
	session.Open()
	defer session.Close()

	// Set the bot's status
	err = session.UpdateGameStatus(0, config.Status)
	util.CheckError(err, "Failed to set status")

	// Temporary, stops the bot from instantly logging out
	time.Sleep(8 * time.Second)
}