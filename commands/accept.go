package commands

import (
	"fmt"
	"log"

	"cordle/game"

	"github.com/bwmarrin/discordgo"
)

// duelAccept attempts to accept a duel challenge
func duelAccept(s *discordgo.Session, i *discordgo.InteractionCreate){
	// Find the challenge
	c := game.FindChallenge(i.Interaction.Member.User)
	if c != nil{
		// Accept the challenge
		game.CloseChallenge(c)
		// Notify the players that the challenge was accepted
		m := fmt.Sprintf(
			"%s, %s has accepted your duel! Please proceed to the breakout thread to play.",
			c.Source.Mention(),
			i.Interaction.Member.Mention(),
		)
		respond(s, i, m, false)
		// Create a new thread to duel in
		th, err := s.MessageThreadStartComplex(i.ChannelID, "", &discordgo.ThreadStart{
			Name: 					"test thread",
			AutoArchiveDuration: 	60,
		})
		if err != nil{
			// Report the error
			s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
				Content: "Failed to create a thread for your game",
			})
			log.Printf("Failed to create breakout thread for duel")
		} else{
			// Create and store the game
			game.NewGame(game.Duel, th.ID, []*discordgo.User{c.Source, c.Target})
		}
	} else{
		// No challenge was found against this user
		respond(s, i, "You currently have no active challenges against you.", true)
	}
}