package commands

import (
	"github.com/bwmarrin/discordgo"
)

func help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	r :=
		`Cordle Bot commands:
		'/duel <player>' duel a given player
		'/accept' accept a duel
		'/reject' reject a duel
		'/stats' get the stats of a given player
		'/leaderboard' get the top ten players in your server

		Whilst in a game:
		'/guess <word>' when you are in a duel use this command to guess a word.
		'[surrender]' when in a game you can use a surrender button
		`

	respond(s, i, r, true)
}
