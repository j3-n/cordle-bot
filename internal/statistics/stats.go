package statistics

import (
	"cordle/internal/database"
	"fmt"
)

func GetStats(d *database.Db, id string) string {
	user := d.ReadUser(id)

	return fmt.Sprintf(
		"``\n%s\n``",
		user.ToStat())
}
