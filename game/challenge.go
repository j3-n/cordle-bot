package game

import (
	"github.com/bwmarrin/discordgo"
)

// Challenge stores information about an active challenge between two players
type Challenge struct{
	Source	*discordgo.User
	Target	*discordgo.User
}

// Stores all currently active challenges
var challenges []*Challenge

// NewChallenge creates a new challenge between two players
func NewChallenge(s *discordgo.User, t *discordgo.User){
	challenges = append(challenges, &Challenge{
		Source: s,
		Target: t,
	})
}

// Locates and returns a challenge between two players given the target user object
// Returns nil if one is not found
func FindChallenge(t *discordgo.User) (*Challenge){
	for _, c := range challenges{
		if c.Target.ID == t.ID{
			return c
		}
	}

	return nil
}