package statistics

import (
	"cordle/database"
	"fmt"
)

func GetStats(id int) string {
	user := database.GetUser(id)

	return fmt.Sprintf(
		"``\n%s\n``",
		user.ToStat())
}
