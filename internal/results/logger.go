package results

import (
	"cordle/internal/database"
)

func LogWin(d *database.Db, attempts int, id string) {
	user := d.ReadUser(id)
	winCalc(attempts, &user)

	d.UpdateUser(&user)
}

func LogLoss(d *database.Db, attempts int, id string) {
	user := d.ReadUser(id)
	loseCalc(attempts, &user)

	d.UpdateUser(&user)
}

func LogDraw(d *database.Db, id string) {
	user := d.ReadUser(id)
	drawCalc(&user)

	d.UpdateUser(&user)
}
