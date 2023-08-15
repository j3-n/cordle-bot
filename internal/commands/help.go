package commands

import (
	"github.com/bwmarrin/discordgo"
)

// help offers a help menu for users to get quick help on how to play Cordle
func help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Send back the help menu
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				getHelpMenu(s),
			},
			Components: []discordgo.MessageComponent{
				&discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						&discordgo.Button{
							Label:    "Getting Started",
							Style:    discordgo.PrimaryButton,
							CustomID: "help_getting_started",
						},
						&discordgo.Button{
							Label:    "Playing the Game",
							Style:    discordgo.PrimaryButton,
							CustomID: "help_game",
						},
						&discordgo.Button{
							Label:    "Scoring System",
							Style:    discordgo.PrimaryButton,
							CustomID: "help_scoring",
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}

// getHelpMenu creates an embed containing the main help menu
func getHelpMenu(s *discordgo.Session) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    s.State.Ready.User.Username,
			IconURL: s.State.Ready.User.AvatarURL("64"),
		},
		Color:       0x009919,
		Title:       "Help",
		Description: "Select the topic you would like help with.",
	}
}
