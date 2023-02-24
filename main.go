package main

import (
	"cordle/database"
	"cordle/statistics"
	"fmt"
)

func main() {
	user := database.GetStats(4561123)
	fmt.Println(user.ToString(), "\n\n")
	fmt.Println(statistics.GetStats(4561123))

	fmt.Println(statistics.GetLeaderboard())
	defer database.Disconnect()
}
