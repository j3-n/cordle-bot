package commands

import (
	"fmt"
	"time"

	"cordle/config"
	"cordle/game"

	"github.com/bwmarrin/discordgo"
)

func duel(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Check whether the target is valid (not a role and not a bot)
	o := i.ApplicationCommandData().Options[0]
	if(o.Type == discordgo.ApplicationCommandOptionUser && !o.UserValue(s).Bot){
		// Check that the target does not already have a challenge against them
		user := o.UserValue(s)
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
			})
		} else {
			// Challenge already exists
			ephemeralResponse(s, i, "That player already has an active challenge against them!")
		}
	} else {
		ephemeralResponse(s, i, "Please challenge a valid user!")
	}
}