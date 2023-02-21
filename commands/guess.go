package commands

import (
	"cordle/game"
	"cordle/wordle"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// guess submits a guess to an ongoing game of Cordle
func guess(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, exists := game.FindGame(i.Interaction.ChannelID)
	if exists {
		p := i.Interaction.Member.User
		if g.PlayerInGame(p) {
			// Retrive the guess
			guess := strings.ToLower(i.ApplicationCommandData().Options[0].StringValue())
			r, err := g.SubmitGuess(guess, p)
			if err == nil {
				// Guess was valid, return result
				respond(s, i, guess+"\n"+displayGuess(r), true)
				// Check if a player has won the game
				won, id := g.GameWon()
				if won {
					// Notify the players that the game has been won
					s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("<@%s> has won the game! The word was `%s`.", id, g.GoalWord(p)))
					// Close the game
					closeGame(s, i.ChannelID)
				} else if !g.PlayerHasGuesses(p) {
					// Notify the players that one has run out of guesses
					s.ChannelMessageSend(i.ChannelID, p.Mention()+" has run out of guesses!")
					// Check if both players have run out of guesses
					if g.ShouldEndInDraw() {
						// End the game in a draw
						s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("All players are out of guesses! the word was `%s`", g.GoalWord(p)))
						closeGame(s, i.ChannelID)
					}
				}
			} else {
				// Send the error message back to the user
				respond(s, i, err.Error(), true)
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

// displayGuess returns a nicely formatted response from a guess result to send back to the user
func displayGuess(r [5]wordle.GuessState) string {
	var s strings.Builder
	for _, gs := range r {
		if gs == wordle.CorrectCharacter {
			s.WriteRune('🟩')
		} else if gs == wordle.IncorrectPosition {
			s.WriteRune('🟨')
		} else {
			s.WriteRune('🟥')
		}
	}
	return s.String()
}

// Close game closes the current game
func closeGame(s *discordgo.Session, th string) {
	// Remove the game internally
	game.CloseGame(th)
	// Archive and lock the thread from discord
	archived := true
	locked := true
	_, err := s.ChannelEditComplex(th, &discordgo.ChannelEdit{
		Archived: &archived,
		Locked:   &locked,
	})
	if err != nil {
		log.Println(err.Error())
	}
}
