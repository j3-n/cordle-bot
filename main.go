package main

import (
	"cordle/database"
	"cordle/statistics"
	"fmt"
)

func main() {
	// newUser := database.User{
	// 	Id:     4561123,
	// 	Name:   "derek",
	// 	Wins:   5,
	// 	Losses: 3,
	// 	Draws:  2,
	// 	Games:  10,
	// 	Elo:    510,
	// 	Level:  2}

	// database.AddUser(newUser)

	user := database.GetUser(4561123)
	fmt.Println(user.ToString())

	topTen := database.GetTop()
	fmt.Println(statistics.GetLeaderboard(topTen))
}
