package commands

import (
	"cordle/game"

	"github.com/bwmarrin/discordgo"
)

func duel(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Check that the target does not already have a challenge against them
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if game.FindChallenge(user) == nil{
		// Create a new challenge
		game.NewChallenge(i.Interaction.Member.User, user)
		// Respond to the interaction
		ephemeralResponse(s, i, "Challenge created, good luck!")
	} else {
		// Challenge already exists
		ephemeralResponse(s, i, "That player already has an active challenge against them!")
	}
}