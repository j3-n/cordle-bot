package main

import (
	"log"
	"github.com/bwmarrin/discordgo"
	"cordle/config"
)

// Path to read config from
const ConfigPath = "config/config.json"

func main() {
	// Load config file
	config := config.LoadConfig(ConfigPath)

	// Start discord bot
	_, err := discordgo.New("Bot " + config.Token)
	if err != nil{
		log.Fatal("Failed to start discord")
	}
}