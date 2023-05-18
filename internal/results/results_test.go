package results

import (
	"cordle/internal/users"
	"fmt"
	"log"
	"testing"
)

func TestWinCalc(t *testing.T) {
	uB := users.User{
		Id:     123,
		Wins:   10,
		Losses: 2,
		Draws:  5,
		Elo:    500,
	}

	uA := uB

	winCalc(3, &uA)

	if uA.Wins <= uB.Wins {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Elo <= uB.Elo {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Draws != uB.Draws {
		log.Fatalln(fmt.Errorf("error calculating win %s\n->\n %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Losses != uB.Losses {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}
}

func TestLoseCalc(t *testing.T) {
	uB := users.User{
		Id:     123,
		Wins:   10,
		Losses: 2,
		Draws:  5,
		Elo:    500,
	}

	uA := uB

	loseCalc(3, &uA)

	if uA.Losses <= uB.Losses {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Elo >= uB.Elo {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Draws != uB.Draws {
		log.Fatalln(fmt.Errorf("error calculating win %s\n->\n %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Wins != uB.Wins {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}
}

func TestDrawCalc(t *testing.T) {
	uB := users.User{
		Id:     123,
		Wins:   10,
		Losses: 2,
		Draws:  5,
		Elo:    500,
	}

	uA := uB

	drawCalc(&uA)

	if uA.Elo <= uB.Elo {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Draws <= uB.Draws {
		log.Fatalln(fmt.Errorf("error calculating win %s\n->\n %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Losses != uB.Losses {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}

	if uA.Wins != uB.Wins {
		log.Fatalln(fmt.Errorf("error calculating win %s -> %s", uA.ToStr(), uB.ToStr()))
	}
}
