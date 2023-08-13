package commands

import (
	"cordle/internal/game"

	"github.com/bwmarrin/discordgo"
)

func leaderboard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	r := game.GetLeaderboard(s)

	respond(s, i, r, false)
}
