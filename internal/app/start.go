package app

import (
	"cordle/internal/commands"
	"cordle/internal/config"
	"cordle/internal/database"
	"cordle/internal/pkg/util"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Run() {
	db := database.NewDb(config.Config.Database)
	defer db.Close()

	// Create discord bot
	session, err := discordgo.New("Bot " + config.Config.Discord.Token)
	util.CheckErrMsg(err, "Failed to initialise discord session")

	// Add a handler to print a happy message when the bot logs in successfully
	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s#%s", r.User.Username, r.User.Discriminator)
	})

	// Start the bot
	err = session.Open()
	util.CheckErrMsg(err, "Failed to open session")
	defer session.Close()

	// Register the bot's commands
	log.Println("Registering slash commands:")
	commands.RegisterCommands(session)
	log.Println("Done registering commands")

	// Set the bot's status
	err = session.UpdateGameStatus(0, config.Config.Discord.Status)
	util.CheckErrMsg(err, "Failed to set status")

	// Keep the program running until SIGTERM or SIGINT is received
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	// Unregister commands
	log.Println("Clearing slash commands:")
	commands.ClearCommands(session)
	log.Println("Done clearing commands")
}

// getConnStr creates a connection string for the database from the loaded config
func getConnStr() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Address,
		config.Config.Database.Port,
		config.Config.Database.Database)
}
