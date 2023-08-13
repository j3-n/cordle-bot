package game

import (
	"cordle/internal/pkg/util"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// GetStats renders a players stats into an embed
func GetStats(u *discordgo.User) *discordgo.MessageEmbed {
	e, err := db.CheckUser(u.ID)
	util.PrintErr(err)
	if !e {
		return &discordgo.MessageEmbed{
			Description: "That user has never played a game!",
		}
	}

	p, err := db.ReadUser(u.ID)
	util.PrintErr(err)

	return &discordgo.MessageEmbed{
		Title: "Stats",
		Color: 0x005f8f,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    u.Username,
			IconURL: u.AvatarURL("64"),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Score",
				Value:  fmt.Sprintf("`%d`", p.Elo),
				Inline: true,
			},
			{
				Name:   "Wins",
				Value:  fmt.Sprintf("`%d`", p.Wins),
				Inline: true,
			},
			{
				Name:   "Losses",
				Value:  fmt.Sprintf("`%d`", p.Losses),
				Inline: true,
			},
			{
				Name:   "Draws",
				Value:  fmt.Sprintf("`%d`", p.Draws),
				Inline: true,
			},
			{
				Name:   "Games Played",
				Value:  fmt.Sprintf("`%d`", p.Wins+p.Losses+p.Draws),
				Inline: true,
			},
		},
	}
}
