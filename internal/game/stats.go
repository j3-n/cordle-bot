package game

import (
	"cordle/internal/pkg/util"
	"fmt"
)

func GetStats(n string, i string) string {
	e, err := db.CheckUser(i)
	util.PrintErr(err)
	if !e {
		return fmt.Sprintf(
			"%s has never played a game!\n",
			n,
		)
	}

	u, err := db.ReadUser(i)
	util.PrintErr(err)

	return fmt.Sprintf(
		"``%s\n%s\n``",
		n,
		u.ToStat(),
	)
}
