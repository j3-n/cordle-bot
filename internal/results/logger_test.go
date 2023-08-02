//go:build integration

package results

import (
	"cordle/internal/config"
	db "cordle/internal/database"
	"cordle/internal/users"
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d *db.Db

func TestMain(m *testing.M) {
	d = db.NewDb(config.Config.Database)
	m.Run()
}

func TestPing(t *testing.T) {
	err := d.Ping()
	assert.NoError(t, err)
}

func TestLogWin(t *testing.T) {
	assert.NotNil(t, d)

	ub := users.User{
		Id:     "123987",
		Wins:   4,
		Losses: 3,
		Draws:  2,
		Elo:    512,
	}

	d.AddUser(&ub)
	e, err := d.CheckUser(ub.Id)
	assert.NoError(t, err)

	if !e {
		log.Fatalln(errors.New("Failed to add user"))
	}

	LogWin(d, 5, ub.Id)

	ua, err := d.ReadUser(ub.Id)
	assert.NoError(t, err)

	if ua.Draws != ub.Draws {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log win failed, draw mismatch"))
	}

	if ua.Wins != ub.Wins+1 && ua.Losses != ub.Losses {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log win failed, win / loss mismatch"))
	}

	if ua.Elo != ub.Elo+8 {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log win failed, elo mismatch"))
	}

	d.DeleteUser(ub.Id)
}

func TestLogLoss(t *testing.T) {
	d := db.NewDb(config.Config.Database)
	defer d.Close()

	assert.NotNil(t, d)

	ub := users.User{
		Id:     "123987",
		Wins:   4,
		Losses: 3,
		Draws:  2,
		Elo:    512,
	}

	d.AddUser(&ub)
	e, err := d.CheckUser(ub.Id)
	assert.NoError(t, err)

	if !e {
		log.Fatalln(errors.New("Failed to add user"))
	}

	LogLoss(d, 5, ub.Id)

	ua, err := d.ReadUser(ub.Id)
	assert.NoError(t, err)

	if ua.Draws != ub.Draws {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log loss failed, draw mismatch"))
	}

	if ua.Wins != ub.Wins && ua.Losses != ub.Losses+1 {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log loss failed, win / loss mismatch"))
	}

	if ua.Elo != ub.Elo-20 {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log loss failed, elo mismatch"))
	}

	d.DeleteUser(ub.Id)
}

func TestLogDraw(t *testing.T) {
	d := db.NewDb(config.Config.Database)
	defer d.Close()

	assert.NotNil(t, d)

	ub := users.User{
		Id:     "123987",
		Wins:   4,
		Losses: 3,
		Draws:  2,
		Elo:    512,
	}

	d.AddUser(&ub)
	e, err := d.CheckUser(ub.Id)
	assert.NoError(t, err)

	if !e {
		log.Fatalln(errors.New("Failed to add user"))
	}

	LogDraw(d, ub.Id)

	ua, err := d.ReadUser(ub.Id)
	assert.NoError(t, err)

	if ua.Draws != ub.Draws+1 {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log draw failed, draw mismatch"))
	}

	if ua.Wins != ub.Wins && ua.Losses != ub.Losses {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log draw failed, win / loss mismatch"))
	}

	if ua.Elo != ub.Elo+3 {
		d.DeleteUser(ub.Id)
		log.Fatalln(errors.New("Log draw failed, elo mismatch"))
	}

	d.DeleteUser(ub.Id)
}
