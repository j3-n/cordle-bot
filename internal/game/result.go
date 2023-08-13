package game

// Scoring constants
const (
	ScoreWin  = 1
	ScoreLoss = 0
	ScoreDraw = 0.5
)

type Result struct {
	Winner string
	Loser  string
	Score  float64
}
