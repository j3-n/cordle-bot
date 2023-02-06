package wordle

// Stores information about a game of Wordle (this has no information linking to Discord)
type WordleGame struct {
	Guesses 	[6]string
	GoalWord 	string
}