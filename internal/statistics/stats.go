package statistics

import (
	"cordle/internal/database"
	"cordle/internal/pkg/util"
	"fmt"
)

func GetStats(d *database.Db, id string) string {
	user, err := d.ReadUser(id)
	util.PrintErr(err)

	return fmt.Sprintf(
		"``\n%s\n``",
		user.ToStat())
}
