package commands

import (
	"cordle/game"

	"github.com/bwmarrin/discordgo"
)

// guess submits a guess to an ongoing game of Cordle
func guess(s *discordgo.Session, i *discordgo.InteractionCreate){
	// TODO: Checks
	// - Check that user has guesses remaining
	g, exists := game.FindGame(i.Interaction.ChannelID)
	if exists {
		if(!g.PlayerInGame(i.Interaction.Member.User)){
			c := g.PlayerCanGuess(i.Interaction.Member.User)
			if(c == game.CanGuess){

			} else if c == game.NoGuesses {
				// The player has run out of guesses
				respond(s, i, "You have run out of guesses!", true)
			}
		} else {
			// The player does not belong to this game
			respond(s, i, "You are not part of this game!", true)
		}
	} else {
		// No game in this channel
		respond(s, i, "There are no active games in this channel", true)
	}
}