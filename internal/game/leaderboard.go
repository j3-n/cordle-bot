package game

import (
	"cordle/internal/pkg/util"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// GetLeaderboard renders the leaderboard into a discord message embed
func GetLeaderboard(s *discordgo.Session) *discordgo.MessageEmbed {
	t, err := db.ReadTop()
	util.PrintErr(err)

	// Create array of top users
	fields := make([]*discordgo.MessageEmbedField, len(t))
	for index, user := range t {
		u, err := s.User(user.Id)
		util.PrintErr(err)
		n := u.Username

		// Render user info into a field
		fields[index] = &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("#%d : %s", index+1, n),
			Value:  fmt.Sprintf("Score: `%d`", user.Elo),
			Inline: false,
		}
	}
	return &discordgo.MessageEmbed{
		Title:  "Top 10 Cordle Users",
		Color:  0xdb4500,
		Fields: fields,
	}
}
