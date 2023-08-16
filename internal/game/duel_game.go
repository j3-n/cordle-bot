package game

import (
	"cordle/internal/config"
	"cordle/internal/wordle"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// DuelGame holds the information about a DuelGame
type DuelGame struct {
	// session stores the session that the DuelGame belongs to
	session *discordgo.Session
	// channel stores a string ChannelID that the game is taking place in
	channel string
	// games stores a map of users to their game
	games map[string]*wordle.WordleGame
	// menus stores the interaction to edit to display games to each user
	menus map[string]*discordgo.InteractionCreate
	// timer stores a pointer to a Timer that tracks inactivity in the game
	timer *time.Timer
}

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func NewDuelGame(th string, p []*discordgo.User, s *discordgo.Session) {
	// Create the shared game
	g0 := wordle.NewRandomGame()
	// Manually create a second game with the same goal word
	// This is more efficient than doing a deep copy
	g1 := &wordle.WordleGame{
		Guesses:  []*wordle.Guess{},
		GoalWord: g0.GoalWord,
	}
	// Create the game struct
	g := &DuelGame{
		session: s,
		channel: th,
		games: map[string]*wordle.WordleGame{
			p[0].ID: g0,
			p[1].ID: g1,
		},
		menus: make(map[string]*discordgo.InteractionCreate),
	}
	// Create an inactivity timer that will warn players that the game is ending
	g.ResetInactivityTimer()
	// Store the game
	games.mu.Lock()
	games.g[th] = g
	games.mu.Unlock()
}

// PlayerInGame returns true if the given player is a part of the game
func (g *DuelGame) PlayerInGame(p *discordgo.User) bool {
	_, exists := g.games[p.ID]
	return exists
}

// Menus returns the map of menus for this game
func (g *DuelGame) Menus() map[string]*discordgo.InteractionCreate {
	return g.menus
}

// PlayerHasGuesses returns true if the player has guesses remaining in the game
func (g *DuelGame) PlayerHasGuesses(p *discordgo.User) bool {
	return g.games[p.ID].GuessesRemaining() > 0
}

// GetPlayerInteractionMenu searches for and returns the interaction menu for the given player
// Returns a boolean to indicate whether or not the menu was found
func (g *DuelGame) GetPlayerInteractionMenu(p *discordgo.User) (*discordgo.InteractionCreate, bool) {
	r, exists := g.menus[p.ID]
	return r, exists
}

// SetPlayerInteractionMenu stores an interaction to be used as the user's menu
func (g *DuelGame) SetPlayerInteractionMenu(p *discordgo.User, m *discordgo.InteractionCreate) {
	g.menus[p.ID] = m
}

// SubmitGuess allows a guess to be submitted to the game of a given player
// Returns the result as an array of wordle.GuessState
func (g *DuelGame) SubmitGuess(guess string, p *discordgo.User) (*wordle.Guess, error) {
	pg := g.games[p.ID]
	return pg.Guess(guess)
}

// PlayerGuessHistory returns the formatted game history of the player and their opponent
func (g *DuelGame) PlayerGameBoard(p *discordgo.User) *discordgo.MessageEmbed {
	var ghp, gho []*wordle.Guess
	for id, game := range g.games {
		// If it is the local player, store in ghp
		if id == p.ID {
			ghp = game.Guesses
		} else {
			gho = game.Guesses
		}
	}
	gbp := displayGame(ghp, false)
	gbo := displayGame(gho, true)
	return renderGameBoard(gbp, gbo, p)
}

// GoalWord returns the goal word for this game
func (g *DuelGame) GoalWord() string {
	return g.games[getPlayers(g)[0]].GoalWord
}

// PlayerSurrender allows a player to quit an ongoing game
// In a duel game, this should immediately end the game
func (g *DuelGame) PlayerSurrender(p *discordgo.User) {
	for id, g := range g.games {
		if p.ID != id {
			g.Won = true
			return
		}
	}
}

// GameWon returns true if the game has been won, as well as the ID of the winner and loser
func (g *DuelGame) GameWon() (bool, string, string) {
	var w string
	var l string
	for id, g := range g.games {
		if g.Won {
			w = id
		} else {
			l = id
		}
	}
	return w != "", w, l
}

// ShouldEndInDraw returns true if the current game has reached a stalemate and should end in a draw
func (g *DuelGame) ShouldEndInDraw() bool {
	for _, g := range g.games {
		if g.GuessesRemaining() > 0 {
			return false
		}
	}
	return true
}

// ResetInactivityTimer restarts the activity timer for the current game
func (g *DuelGame) ResetInactivityTimer() {
	// If the game has an inactivity timer, cancel it
	if g.timer != nil {
		g.timer.Stop()
	}
	// Start a new inactivity timer
	g.timer = time.AfterFunc(time.Duration(config.Config.Game.InactivityTimeout-config.Config.Game.InactivityWarning)*time.Second, func() {
		// Notify the players that the inactivity limit is near
		g.SendInactivityWarning()
		// Start a new timer that will close the whole game
		g.timer = time.AfterFunc(time.Duration(config.Config.Game.InactivityWarning)*time.Second, func() {
			// Close the game once the inactivity period is up
			g.SendInactivityExpired()
			g.EndGame()
		})
	})
}

// SendInactivityWarning sends a warning that the game is about to be closed for inactivity
func (g *DuelGame) SendInactivityWarning() {
	p := getPlayers(g)
	m := fmt.Sprintf(
		"<@%s>, <@%s>, your game will soon expire due to inactivity. Please submit a new guess within %d seconds.",
		p[0],
		p[1],
		config.Config.Game.InactivityWarning,
	)
	_, err := g.session.ChannelMessageSend(g.channel, m)
	if err != nil {
		log.Printf("Failed to send inactivity warning. [%s]", err.Error())
	}
}

// SendInactivityExpired sends a message to notify players that the game has been closed due to inactivity
func (g *DuelGame) SendInactivityExpired() {
	p := getPlayers(g)
	m := fmt.Sprintf(
		"<@%s>, <@%s>, your game has ended in a draw due to inactivity. The word was `%s`.",
		p[0],
		p[1],
		g.GoalWord(),
	)
	_, err := g.session.ChannelMessageSend(g.channel, m)
	if err != nil {
		log.Printf("Failed to send inactivity expiration notice. [%s]", err.Error())
	}
}

// RegisterResult is called to store the result of the game and update elo scores
func (g *DuelGame) RegisterResult(r *Result) {
	// If the game was a draw, infer the participants
	ps := getPlayers(g)
	if r.Score == ScoreDraw {
		r.Winner = ps[0]
		r.Loser = ps[1]
	}
	// Update the elo scores and return them
	ws, wd, ls, ld := updateScores(r.Winner, r.Loser, r.Score)
	// Send scores to users
	ma := fmt.Sprintf("<@%s>", r.Winner)
	mb := fmt.Sprintf("<@%s>", r.Loser)
	// Render embeds
	ea := renderScore(ws, wd)
	eb := renderScore(ls, ld)
	// Send the messages
	sendScoreEmbed(g, ma, ea)
	sendScoreEmbed(g, mb, eb)
}

// EndGame is called to finish a game
func (g *DuelGame) EndGame() {
	// Remove the game internally
	CloseGame(g.channel)
	// Stop the inactivity timer
	if g.timer != nil {
		g.timer.Stop()
	}

	// Archive and lock the thread from discord after delay
	time.AfterFunc(time.Duration(config.Config.Game.PostGameDelay)*time.Second, func() {
		archived := true
		locked := true
		_, err := g.session.ChannelEditComplex(g.channel, &discordgo.ChannelEdit{
			Archived: &archived,
			Locked:   &locked,
		})
		if err != nil {
			log.Println(err.Error())
		}
	})
}

// getPlayers returns a slice of the IDs of users currently playing in a game
func getPlayers(g *DuelGame) []string {
	p := make([]string, len(g.games))
	i := 0
	for k := range g.games {
		p[i] = k
		i++
	}
	return p
}

// sendScoreEmbed sends a message containing a score embed to the given game channel
func sendScoreEmbed(g *DuelGame, msg string, emb *discordgo.MessageEmbed) {
	_, err := g.session.ChannelMessageSendComplex(g.channel, &discordgo.MessageSend{
		Content: msg,
		Embeds:  []*discordgo.MessageEmbed{emb},
	})
	if err != nil {
		log.Printf("Failed to send elo score message. [%s]", err.Error())
	}
}

// renderScore returns an embed containing an update on a player's score after a game
func renderScore(s int, d int) *discordgo.MessageEmbed {
	// Get the appropriate emoji and colour to use
	var emoji string
	var color int
	if d >= 0 {
		emoji = "eloup"
		color = 0x0f9100
	} else {
		emoji = "elodown"
		color = 0x8a0e00
	}
	return &discordgo.MessageEmbed{
		Color: color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Your Score",
				Value:  fmt.Sprintf("**%d** [ %s %d ]", s, Emojis[emoji], d),
				Inline: false,
			},
		},
	}
}

// displayGame returns a string displaying the given guess history.
// If hide is true, returns the game board without the letters
func displayGame(gh []*wordle.Guess, hide bool) string {
	// Iterate over the slice to build the guess board
	var gb strings.Builder
	for i := 0; i < wordle.MaxGuesses; i++ {
		// Write the player's board
		if i < len(gh) {
			gb.WriteString(displayGuess(gh[i], hide))
		} else {
			// If not all guesses have been filled, add a blank line
			gb.WriteString(blankLine())
		}
		gb.WriteRune('\n')
	}
	// Return the board as a string
	return gb.String()
}

// Given 2 game boards as a string, creates a MessageEmbed containing the board
func renderGameBoard(gbp string, gbo string, p *discordgo.User) *discordgo.MessageEmbed {
	// Create a message embed with the game board inside
	return &discordgo.MessageEmbed{
		Type:  discordgo.EmbedTypeRich,
		Title: "Cordle Game | Duel",
		Color: 0x00b503,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    p.Username,
			IconURL: p.AvatarURL("64"),
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "You",
				Value:  gbp,
				Inline: true,
			},
			{
				Name:   "Your opponent",
				Value:  gbo,
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Your opponent is guessing at the same time as you, try to solve the puzzle before they do! Use /guess to guess again.",
		},
	}
}

// displayGuess returns a nicely formatted response from a guess result to send back to the user.
// If hide is true, returns blank boxes instead of boxes with letters
func displayGuess(r *wordle.Guess, hide bool) string {
	var s strings.Builder
	runes := []rune(r.GuessWord)
	for i, gs := range r.GuessResult {
		prefix := ""
		if gs == wordle.CorrectCharacter {
			prefix = "green"
		} else if gs == wordle.IncorrectPosition {
			prefix = "yellow"
		} else {
			prefix = "grey"
		}
		// Calculate the name of the required emoji and write it
		var e string
		if hide {
			e = fmt.Sprintf("%s_blank", prefix)
		} else {
			e = fmt.Sprintf("%s_%c", prefix, runes[i])
		}
		s.WriteString(Emojis[e])
	}
	return s.String()
}

// blankLine generates a line of five blank emojis
func blankLine() string {
	var s strings.Builder
	for i := 0; i < 5; i++ {
		s.WriteString(Emojis["blank"])
	}
	return s.String()
}
