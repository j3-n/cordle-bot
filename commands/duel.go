package commands

import (
	"fmt"

	"cordle/game"

	"github.com/bwmarrin/discordgo"
)

func duel(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Check that the target does not already have a challenge against them
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if game.FindChallenge(user) == nil{
		// Create a new challenge
		game.NewChallenge(i.Interaction.Member.User, user)
		fmt.Println("Made new challenge")
	} else {
		fmt.Println("Challenge already exists")
	}
}