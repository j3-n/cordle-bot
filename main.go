package main

import (
	"cordle/database"
	"cordle/statistics"
	"fmt"
)

func main() {
	user := database.GetUser(4561123)
	fmt.Println(user.ToString())

	topTen := database.GetTop()
	fmt.Println(statistics.GetLeaderboard(topTen))
	defer database.Disconnect()
}
