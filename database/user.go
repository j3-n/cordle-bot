package database

import (
	"fmt"
)

type User struct {
	Id		int
	Name	string
	Wins	int
	Losses	int
	Draws	int
	Games	int
	Elo		int
	Level	int
}

func (u User) ToString() string {
	return fmt.Sprintf("ID: %d\nName: %s\nWins: %d\nLosses: %d\nDraws: %d\nGames Played: %d\nElo: %d\nLevel: %d", 
			u.Id, 
			u.Name, 
			u.Wins, 
			u.Losses, 
			u.Draws, 
			u.Games, 
			u.Elo, 
			u.Level)
}

func (u User) WinPercentage() float64 {
	return float64(u.Wins) / float64(u.Games)
}