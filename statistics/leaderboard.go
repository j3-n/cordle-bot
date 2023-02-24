package statistics

import (
	"cordle/database"
	"fmt"
)

func GetLeaderboard(topTen []database.User) string {
	output := "``\n"
	for index, user := range topTen {
		output += fmt.Sprintf(
			"%d : %s\n",
			index+1,
			user.ToLeaderboard())
	}
	return output + "``"
}
