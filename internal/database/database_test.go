package database

import (
	"cordle/internal/config"
	"cordle/internal/users"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d *Db

func TestMain(m *testing.M) {
	config.LoadConfig("../../config/config.json")
	os.Exit(m.Run())
}

func TestDb(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	assert.NotNil(t, d)
}

func TestAddUser(t *testing.T) {
	u := users.User{
		Id:     "7567",
		Wins:   20,
		Losses: 53,
		Draws:  151,
		Elo:    341,
	}

	d = NewDb(config.Config.Database)
	defer d.Close()

	d.AddUser(&u)

	e := d.CheckUser(u.Id)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u.Id)
}

func TestAddUsers(t *testing.T) {
	u1 := users.User{
		Id:     "7567",
		Wins:   20,
		Losses: 53,
		Draws:  151,
		Elo:    341,
	}
	u2 := users.User{
		Id:     "1577",
		Wins:   20,
		Losses: 13,
		Draws:  51,
		Elo:    541,
	}

	d = NewDb(config.Config.Database)
	defer d.Close()

	u := make([]users.User, 2)
	u[0] = u1
	u[1] = u2
	d.AddUsers(&u)

	e := d.CheckUser(u1.Id)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u1.Id)

	e = d.CheckUser(u2.Id)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u2.Id)
}

func TestUpdateUser(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	d.AddUser(&users.User{
		Id:     "123",
		Wins:   2,
		Losses: 1,
		Draws:  3,
		Elo:    521,
	})
	e := d.CheckUser("123")
	assert.True(t, e)

	u := d.ReadUser("123")
	draws := u.Draws
	u.Draws += 1
	d.UpdateUser(&u)
	u = d.ReadUser("123")
	if u.Draws != draws+1 {
		log.Fatalln(errors.New("error updating draw count"))
	}

	d.DeleteUser("123")
}

func TestUpdateUsers(t *testing.T) {

}

func TestReadUser(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	u := d.ReadUser("388395158397517824")
	if u.Id != "388395158397517824" {
		log.Fatalln(errors.New("read nil user error"))
	}
	assert.NotNil(t, u)
}

func TestReadAllUsers(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	u := d.ReadAllUsers()
	assert.NotNil(t, u)

	if len(u) == 0 {
		log.Fatalln(fmt.Errorf(
			"incorrect array length for all users %d", len(u)),
		)
	}
}

func TestReadTop(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	tt := d.ReadTop()
	assert.NotNil(t, tt)

	for i := 0; i < len(tt)-2; i++ {
		if tt[i].Elo < tt[i+1].Elo {
			continue
		}
		log.Fatalln(fmt.Errorf(
			"top ten wrong order %d", len(tt)),
		)
	}
}

func TestCheckUser(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	e := d.CheckUser("388395158397517824")
	if !e {
		log.Fatalln(fmt.Errorf(
			"wrong existing value returned %t", e),
		)
	}
	assert.True(t, e)

	e = d.CheckUser("517123")
	if e {
		log.Fatalln(fmt.Errorf(
			"wrong existing value returned %t", e),
		)
	}
	assert.False(t, e)
}

func TestDeleteUser(t *testing.T) {
	u := users.User{
		Id:     "61567",
		Wins:   22,
		Losses: 51,
		Draws:  101,
		Elo:    371,
	}

	d = NewDb(config.Config.Database)
	defer d.Close()

	d.AddUser(&u)

	d.DeleteUser(u.Id)
	e := d.CheckUser(u.Id)
	if e {
		log.Fatalln(errors.New("failed to delete user"))
	}
}

func TestDeleteUsers(t *testing.T) {

}

func TestCreateTable(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	c := `
		create table test_table(
			id int not null primary key,
			name varchar(10)
		);
	`

	d.CreateTable(c)
	e := d.CheckTable("test_table")
	assert.True(t, e)

	d.DeleteTable("test_table")
}

func TestUpdateTable(t *testing.T) {

}

func TestDeleteTable(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	c := `
		create table test_table(
			id int not null primary key,
			name varchar(10)
		);
	`

	d.CreateTable(c)
	e := d.CheckTable("test_table")
	assert.True(t, e)

	d.DeleteTable("test_table")
	e = d.CheckTable("test_table")
	assert.False(t, e)
}

func TestCheckTable(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	e := d.CheckTable("check_table")
	assert.False(t, e)

	c := `
		create table check_table(
			id int not null primary key
		);
	`

	d.CreateTable(c)
	e = d.CheckTable("check_table")
	assert.True(t, e)

	d.DeleteTable("check_table")
	e = d.CheckTable("check_table")
	assert.False(t, e)
}
