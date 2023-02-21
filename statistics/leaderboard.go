package statistics

import (
	"cordle/database"
	"fmt"
)

func GetLeaderboard(topTen [10]database.User) string {
	var output string
	for i := 0; i < len(topTen); i++ {
		output += fmt.Sprintf(
			"``%d : %s``",
			i+1,
			topTen[i].ToLeaderboard())
	}
	return output
}
