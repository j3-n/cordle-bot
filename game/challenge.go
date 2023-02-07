package game

import (
	"github.com/bwmarrin/discordgo"
)

// Challenge stores information about an active challenge between two players
type Challenge struct{
	Source	*discordgo.User
	Target	*discordgo.User
}