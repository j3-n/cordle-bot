package commands

import (
	"cordle/util"

	"github.com/bwmarrin/discordgo"
)

// Big list of available commands
var commands = []*discordgo.ApplicationCommand{
	// Test command to ensure this works
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
	// Iterate over each command and register it
	for _, cmd := range commands{
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		util.CheckError(err, "Failed to create command: /" + cmd.Name)
	}

	// Create a handler to map commands to their handlers
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate){
		h, exists := commandHandlers[i.ApplicationCommandData().Name]
		if exists{
			h(s, i)
		}
	})
}

// ClearCommands retrieves all commands registered with this bot from discord and deletes them
func ClearCommands(s *discordgo.Session){
	
}