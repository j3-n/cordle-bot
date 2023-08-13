package game

import (
	"cordle/internal/pkg/util"
	"fmt"
)

func GetStats(n string, i string) string {
	u, err := db.ReadUser(i)
	util.PrintErr(err)

	return fmt.Sprintf(
		"``%s\n%s\n``",
		n,
		u.ToStat(),
	)
}
