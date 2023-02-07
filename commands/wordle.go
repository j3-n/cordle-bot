package commands

import (
	"cordle/wordle"

	"github.com/bwmarrin/discordgo"
)

var game wordle.WordleGame

// Creates a new game of wordle
func newWordle(s *discordgo.Session, i *discordgo.InteractionCreate){
	game = wordle.NewRandomGame()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Starting new game of Wordle",
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})

	s.FollowupMessageCreate(i.Interaction, false, &discordgo.WebhookParams{
		Content: game.GoalWord,
		Flags: discordgo.MessageFlagsEphemeral,
	})
}