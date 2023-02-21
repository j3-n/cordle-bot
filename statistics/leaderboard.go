package statistics

import (
	"cordle/database"
	"fmt"
)

func GetLeaderboard(topTen [10]database.User) string {
	var output string
	for index, user := range topTen {
		output += fmt.Sprintf(
			"``%d : %s``",
			index+1,
			user.ToLeaderboard())
	}
	return output
}
