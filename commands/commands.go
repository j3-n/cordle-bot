package commands

import (
	"fmt"

	"cordle/util"

	"github.com/bwmarrin/discordgo"
)

// Big list of available commands
var commands = []*discordgo.ApplicationCommand{
	{
		Name: "duel",
		Description: "Send a duel challenge to another player",
	},
	{
		Name: "accept",
		Description: "Accept a duel challenge against you",
	},
	{
		Name: "decline",
		Description: "Decline a duel challenge against you",
	},
	{
		Name: "guess",
		Description: "Submit a guess to a game of wordle",
	},
}

var regCommands []*discordgo.ApplicationCommand

// Big map linking command names to their handling functions
// Handler functions are stored in separate go files
var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"duel": 	duel,
	"accept": 	duelAccept,
	"decline": 	duelDecline,
	"guess": 	guess,
}

// RegisterCommands registers all command with Discord, this is necessary to allow users to run them
func RegisterCommands(s *discordgo.Session){
	// Initialise the registered command array
	regCommands = make([]*discordgo.ApplicationCommand, len(commands))
	// Iterate over each command and register it
	for i, cmd := range commands{
		c, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		util.CheckError(err, "Failed to create command: /" + cmd.Name)
		// Log the command for later deletion
		regCommands[i] = c
	}

	// Create a handler to map commands to their handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate){
		h, exists := commandHandlers[i.ApplicationCommandData().Name]
		if exists{
			h(s, i)
		}
	})
}

// ClearCommands deletes all commands created during this session
func ClearCommands(s *discordgo.Session){
	// Iterate over every registered command and delete it
	for _, cmd := range regCommands{
		err := s.ApplicationCommandDelete(s.State.User.ID, "", cmd.ID)
		util.CheckError(err, fmt.Sprintf("Failed to delete command /%s", cmd.Name))
	}
}