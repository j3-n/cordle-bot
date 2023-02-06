package main

import (
	"log"
	"time"

	"cordle/config"
	"cordle/util"
	"cordle/commands"

	"github.com/bwmarrin/discordgo"
)

// Path to read config from
const ConfigPath = "config/config.json"
// Path to read command metadata from
const CommandsPath string = "commands/commands.json"

func main() {
	// Load config file
	config := config.LoadConfig(ConfigPath)

	// Create discord bot
	session, err := discordgo.New("Bot " + config.Token)
	util.CheckError(err, "Failed to initialise discord session")

	// Add a handler to print a happy message when the bot logs in successfully
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready){
		log.Printf("Logged in as %s#%s", r.User.Username, r.User.Discriminator)
	})

	// Start the bot
	err = session.Open()
	util.CheckError(err, "Failed to open session")
	defer session.Close()

	// Register the bot's commands
	commands.RegisterCommands(session)

	// Set the bot's status
	err = session.UpdateGameStatus(0, config.Status)
	util.CheckError(err, "Failed to set status")

	// Temporary, stops the bot from instantly logging out
	time.Sleep(8 * time.Second)
}