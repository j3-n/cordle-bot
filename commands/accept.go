package commands

import (
	"fmt"

	"cordle/game"

	"github.com/bwmarrin/discordgo"
)

// duelAccept attempts to accept a duel challenge
func duelAccept(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Find the challenge
	c := game.FindChallenge(i.Interaction.User)
	if c != nil{
		// Accept the challenge
		game.CloseChallenge(c)
		// Notify the players that the challenge was accepted
		m := fmt.Sprintf(
			"%s, %s has accepted your duel! Please proceed to the breakout thread to play.",
			i.Interaction.User.Mention(),
			c.Target.Mention(),
		)
		respond(s, i, m, false)
		// TODO: start a new duel game
	} else{
		// No challenge was found against this user
		respond(s, i, "You currently have no active challenges against you.", true)
	}
}