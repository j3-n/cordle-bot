package commands

import (
	"cordle/game"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// duelDecline declines a challenge against a given user
func duelDecline(s *discordgo.Session, i *discordgo.InteractionCreate){
	c := game.FindChallenge(i.Interaction.User)
	if c != nil{
		// Decline the challenge
		game.CloseChallenge(c)
		// Notify players that the challenge has been declined
		m := fmt.Sprintf(
			"%s, %s has declined your duel!",
			c.Target.Mention(),
			i.Interaction.User.Mention(),
		)
		respond(s, i, m, false)
	} else{
		respond(s, i, "You currently have no active challenges against you.", true)
	}
}