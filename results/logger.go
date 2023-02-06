package results

import(
	"fmt"
	"cordle/database"
)

func LogWin(attempts int, player database.User) {
	player = winCalc(attempts, player)
	fmt.Println(player.ToString())

	// update database 
}

func LogLoss(attempts int, player database.User) {
	player = loseCalc(attempts, player)
	fmt.Println(player.ToString())

	// update database
}

func LogDraw(player database.User) {
	player = drawCalc(player)
	fmt.Println(player.ToString())

	// update database
}