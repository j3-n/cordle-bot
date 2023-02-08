package commands

import (
	"fmt"

	"cordle/game"
	"cordle/config"

	"github.com/bwmarrin/discordgo"
)

func duel(s *discordgo.Session, i *discordgo.InteractionCreate){
	// TODO: check whether a role is mentioned
	
	// Check that the target does not already have a challenge against them
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if game.FindChallenge(user) == nil{
		// Create a new challenge
		game.NewChallenge(i.Interaction.Member.User, user)
		// Respond to the interaction, notifying the other player of the duel
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					"%s, %s has challenged you to a duel! You have %d seconds to either `/accept` or `/decline` this duel",
					user.Mention(),
					i.Interaction.Member.Mention(),
					config.Config.Game.ChallengeDuration,
				),
			},
		})
	} else {
		// Challenge already exists
		ephemeralResponse(s, i, "That player already has an active challenge against them!")
	}
}