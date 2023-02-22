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
			_, err := g.SubmitGuess(guess, p)
			if err == nil {
				// Guess was valid, display result
				emb := displayGame(g.PlayerGuessHistory(p))
				// Add some extra info to the embed
				emb.Author = &discordgo.MessageEmbedAuthor{
					Name: p.Username,
					IconURL: p.AvatarURL("128"),
				}
				// Send the response
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Flags: discordgo.MessageFlagsEphemeral,
						Embeds: []*discordgo.MessageEmbed{emb},
					},
				})
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
						s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("All players are out of guesses! The word was `%s`.", g.GoalWord(p)))
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

// displayGame returns a MessageEmbed displaying the given guess history
func displayGame(gs []*wordle.Guess) *discordgo.MessageEmbed {
	// Iterate over the slice to build the guess board
	var gb strings.Builder
	for _, g := range gs {
		gb.WriteString(fmt.Sprintf("%s | `%s`\n", displayGuess(g),g.GuessWord))
	}
	// Return the board inside an embed
	return &discordgo.MessageEmbed{
		Type: discordgo.EmbedTypeRich,
		Color: 0x00b503,
		Title: "Cordle Game",
		Description: gb.String(),
	}
}

// displayGuess returns a nicely formatted response from a guess result to send back to the user
func displayGuess(r *wordle.Guess) string {
	var s strings.Builder
	for _, gs := range r.GuessResult {
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
