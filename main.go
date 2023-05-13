package main

import (
	"log"
	"os"
	"os/signal"

	"cordle/commands"
	"cordle/config"
	"cordle/util"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create discord bot
	session, err := discordgo.New("Bot " + config.Config.Token)
	util.CheckError(err, "Failed to initialise discord session")

	// Add a handler to print a happy message when the bot logs in successfully
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s#%s", r.User.Username, r.User.Discriminator)
	})

	// Start the bot
	err = session.Open()
	util.CheckError(err, "Failed to open session")
	defer session.Close()

	// Register the bot's commands
	log.Println("Registering slash commands:")
	commands.RegisterCommands(session)
	log.Println("Done registering commands")

	// Set the bot's status
	err = session.UpdateGameStatus(0, config.Config.Status)
	util.CheckError(err, "Failed to set status")

	// Keep the program running until interrupted
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Unregister commands
	log.Println("Clearing slash commands:")
	commands.ClearCommands(session)
	log.Println("Done clearing commands")
}
