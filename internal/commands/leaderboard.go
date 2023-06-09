package commands

import (
	"cordle/internal/users"

	"github.com/bwmarrin/discordgo"
)

func leaderboard(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, "Invalid target", true)
		return
	}

	u := users.User{
		Id:     "123",
		Wins:   123,
		Losses: 123,
		Draws:  123,
		Elo:    567,
	}

	respond(s, i, u.ToLeaderboard(), false)
}
