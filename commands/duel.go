package commands

import (
	"fmt"
	"time"

	"cordle/game"
	"cordle/config"

	"github.com/bwmarrin/discordgo"
)

func duel(s *discordgo.Session, i *discordgo.InteractionCreate){
	// TODO: check whether target is valid
	
	// Check that the target does not already have a challenge against them
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if game.FindChallenge(user) == nil{
		// Create a new challenge
		c := game.NewChallenge(i.Interaction.Member.User, user)
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
		// After the delay, delete the challenge notification and expire the challenge
		time.AfterFunc(time.Duration(config.Config.Game.ChallengeDuration) * time.Second, func(){
			s.InteractionResponseDelete(i.Interaction)
			game.CloseChallenge(c)
			// TODO: check this works
		})
	} else {
		// Challenge already exists
		ephemeralResponse(s, i, "That player already has an active challenge against them!")
	}
}