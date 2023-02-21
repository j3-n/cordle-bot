package results

import (
	"cordle/database"
	"fmt"
)

func LogWin(attempts int, id int) {
	user := database.GetUser(id)
	winCalc(attempts, &user)
	fmt.Println(user.ToString())

	database.UpdateUser(&user)
}

func LogLoss(attempts int, id int) {
	user := database.GetUser(id)
	loseCalc(attempts, &user)
	fmt.Println(user.ToString())

	database.UpdateUser(&user)
}

func LogDraw(id int) {
	user := database.GetUser(id)
	drawCalc(&user)
	fmt.Println(user.ToString())

	database.UpdateUser(&user)
}
