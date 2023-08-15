package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// helpMenus will store the most up to date interaction for each user's help menu
var helpMenus map[string]*discordgo.Interaction

func init() {
	helpMenus = make(map[string]*discordgo.Interaction)
}

// help offers a help menu for users to get quick help on how to play Cordle
func help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Store the interaction
	helpMenus[i.Interaction.Member.User.ID] = i.Interaction
	// Send back the help menu
	s.InteractionRespond(i.Interaction, getHelpMenu(s))
}

// gettingStarted sends the user a general "Getting Started" help menu
func gettingStarted(s *discordgo.Session, i *discordgo.InteractionCreate) {
	emb := &discordgo.MessageEmbed{
		Author: getMessageEmbedAuthor(s),
		Title:  "Help: Getting Started",
		Description: "Welcome to Cordle! Cordle Bot allows you to play competitive Wordle against your friends inside Discord! " +
			"Games of Cordle are duels, between you and one other player. To start a game of Cordle, challenge a friend with `/duel`. " +
			"Once your opponent accepts, the game will begin. If a friend challenges you, simply use `/accept` to accept their challenge, " +
			"or `/decline` if you'd rather not. Head over to \"Playing the Game\" to get help on how to play Cordle.",
	}
	helpMenuEdit(s, i, emb)
}

// playingGame displays help on playing Cordle
func playingGame(s *discordgo.Session, i *discordgo.InteractionCreate) {
	emb := &discordgo.MessageEmbed{
		Author: getMessageEmbedAuthor(s),
		Title:  "Help: Playing the Game",
		Description: "Once your game begins, head into the breakout thread that the bot creates for you. " +
			"A game of Cordle is a race against your opponent to solve the Wordle puzzle. You are both working " +
			"towards the same word. To make a guess in your puzzle, use `/guess` followed by your five-letter guess. " +
			"The bot will then update your board to show your new guess. Your opponent's board is also shown, updated " +
			"in real-time as they guess. The first player to solve the puzzle is the winner. If you both fail to solve " +
			"the puzzle in the given number of guesses, the game ends in a draw. If you wish to leave the game, either use " +
			"`/surrender` or click the red \"Surrender\" button at the top of the breakout thread. Head over to the \"Scoring System\" " +
			"section to learn about how the competitive scoring system works, as well as its related commands.",
	}
	helpMenuEdit(s, i, emb)
}

// scoring displays help with Cordle's scoring system
func scoring(s *discordgo.Session, i *discordgo.InteractionCreate) {
	emb := &discordgo.MessageEmbed{
		Author: getMessageEmbedAuthor(s),
		Title:  "Help: Getting Started",
		Description: "Cordle has a competitive ranking system so you can truly prove you are better than your friends at Wordle. Each player " +
			"starts with a score of 1000 on their first game. Winning games will cause you to take points from your opponent, and likewise, losing games " +
			"will see your opponent taking your points. The amoint of points gained or lost depends on the difference in score between the two players. " +
			"You can view the career stats of any player using `/stats`. You can also view the Top 10 Leaderboard using the `/leaderboard` command.",
	}
	helpMenuEdit(s, i, emb)
}

// helpMenuEdit edits the current help menu and replaces it with the requested page
func helpMenuEdit(s *discordgo.Session, i *discordgo.InteractionCreate, emb *discordgo.MessageEmbed) {
	// Get the interaction response to edit
	menu, exists := helpMenus[i.Interaction.Member.User.ID]

	if exists {
		_, err := s.InteractionResponseEdit(menu, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{
				emb,
			},
			Components: getHelpMenuButtons(),
		})

		if err != nil {
			log.Println(err)
		}
	}

	// Respond to the interaction
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
	})
}

// getHelpMenu renders the help menu into a discord interaction response
func getHelpMenu(s *discordgo.Session) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				getHelpMenuEmbed(s),
			},
			Components: *getHelpMenuButtons(),
			Flags:      discordgo.MessageFlagsEphemeral,
		},
	}
}

// getHelpMenuEmbed creates an embed containing the main help menu
func getHelpMenuEmbed(s *discordgo.Session) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author:      getMessageEmbedAuthor(s),
		Color:       0x009919,
		Title:       "Help",
		Description: "Select the topic you would like help with.",
	}
}

// getHelpMenuButtons returns the ActionsRow containing the help menu buttons
func getHelpMenuButtons() *[]discordgo.MessageComponent {
	return &[]discordgo.MessageComponent{
		&discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				&discordgo.Button{
					Label:    "Getting Started",
					Style:    discordgo.PrimaryButton,
					CustomID: "help_getting_started",
				},
				&discordgo.Button{
					Label:    "Playing the Game",
					Style:    discordgo.PrimaryButton,
					CustomID: "help_game",
				},
				&discordgo.Button{
					Label:    "Scoring System",
					Style:    discordgo.PrimaryButton,
					CustomID: "help_scoring",
				},
			},
		},
	}
}

// getMessageEmbedAuthor returns the bot's information as a MessageEmbedAuthor
func getMessageEmbedAuthor(s *discordgo.Session) *discordgo.MessageEmbedAuthor {
	return &discordgo.MessageEmbedAuthor{
		Name:    s.State.Ready.User.Username,
		IconURL: s.State.Ready.User.AvatarURL("64"),
	}
}
