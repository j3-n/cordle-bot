package statistics

import (
	"cordle/internal/database"
	"cordle/internal/pkg/util"
	"fmt"
)

func GetLeaderboard(d *database.Db) string {
	topTen, err := d.ReadTop()
	util.PrintErr(err)

	output := "``\n"
	for index, user := range topTen {
		output += fmt.Sprintf(
			"%d : %s\n",
			index+1,
			user.ToLeaderboard())
	}
	return output + "``"
}
