package main

import (
	"cordle/database"
	"fmt"
)

func main() {
	dave := database.User{
		Id: 678678,
		Name: "dave",
		Wins: 135,
		Losses: 10,
		Draws: 50,
		Games: 195,
		Elo: 1261,
		Level: 120}

	database.AddUser(dave)

	user := database.GetUser(456123)
	fmt.Println(user.ToString())

	fmt.Println("")

	stats := database.GetStats(456123)
	fmt.Println(stats.ToString())
}
