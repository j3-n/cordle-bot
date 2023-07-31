package results

import (
	"cordle/internal/database"
	"cordle/internal/pkg/util"
)

func LogWin(d *database.Db, attempts int, id string) {
	user, err := d.ReadUser(id)
	util.PrintErr(err)
	winCalc(attempts, &user)

	d.UpdateUser(&user)
}

func LogLoss(d *database.Db, attempts int, id string) {
	user, err := d.ReadUser(id)
	util.PrintErr(err)
	loseCalc(attempts, &user)

	d.UpdateUser(&user)
}

func LogDraw(d *database.Db, id string) {
	user, err := d.ReadUser(id)
	util.PrintErr(err)
	drawCalc(&user)

	d.UpdateUser(&user)
}
