package results

import (
	"cordle/database"
	"fmt"
)

func LogWin(attempts int, player *database.User) {
	winCalc(attempts, player)
	fmt.Println(player.ToString())

	database.UpdateUser(player)
}

func LogLoss(attempts int, player *database.User) {
	loseCalc(attempts, player)
	fmt.Println(player.ToString())

	database.UpdateUser(player)
}

func LogDraw(player *database.User) {
	drawCalc(player)
	fmt.Println(player.ToString())

	database.UpdateUser(player)
}
