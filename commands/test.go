package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Test command
func test(s *discordgo.Session, i *discordgo.InteractionCreate){
	fmt.Println("Received test command")
}