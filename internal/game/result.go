package game

// Scoring constants
const (
	SCORE_WIN  = 1
	SCORE_LOSS = 0
	SCORE_DRAW = 0.5
)

type Result struct {
	Winner string
	Loser  string
	Score  float64
}
