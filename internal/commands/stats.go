package commands

import (
	"cordle/internal/game"

	"github.com/bwmarrin/discordgo"
)

func stats(s *discordgo.Session, i *discordgo.InteractionCreate) {
	user := i.ApplicationCommandData().Options[0].UserValue(s)
	if user.Bot {
		respond(s, i, "Invalid target", true)
		return
	}

	r := game.GetStats(user)

	embedRespond(s, i, r)
}
