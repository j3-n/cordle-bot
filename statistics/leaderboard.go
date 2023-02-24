package statistics

import (
	"cordle/database"
	"fmt"
)

func GetLeaderboard() string {
	topTen := database.GetTop()
	output := "``\n"
	for index, user := range topTen {
		output += fmt.Sprintf(
			"%d : %s\n",
			index+1,
			user.ToLeaderboard())
	}
	return output + "``"
}
