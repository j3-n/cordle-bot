package commands

import "github.com/bwmarrin/discordgo"

// resetMenu deletes the old interaction response that contained a game board and
// responds with a new one.
func resetMenu(s *discordgo.Session, i *discordgo.InteractionCreate) {
	respond(s, i, "Penis", true)
}
