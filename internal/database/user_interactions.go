package database

import (
	"cordle/internal/pkg/util"
	"cordle/internal/users"
	"fmt"
)

func (d *Db) AddUser(user *users.User) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	insert, err := d.Client.Db.Query(fmt.Sprintf(
		`insert into users(id, wins, losses, draws, elo) 
		values(%s);`,
		user.ToSqlAdd(),
	))

	util.CheckErr(err)
	defer insert.Close()
}

func (d *Db) AddUsers(users *[]users.User) {
	for _, user := range *users {
		d.AddUser(&user)
	}
}

func (d *Db) UpdateUser(user *users.User) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	q := fmt.Sprintf(
		"id='%s'",
		user.Id,
	)

	update, err := d.Client.Db.Query(fmt.Sprintf(
		`update users
		 set %s
		 where %s;`,
		user.ToSqlUpdate(),
		q,
	))

	util.CheckErr(err)
	defer update.Close()
}

func (d *Db) UpdateUsers(users *[]users.User) {
	for _, user := range *users {
		d.UpdateUser(&user)
	}
}

func (d *Db) ReadUser(id string) users.User {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	result, err := d.Client.Db.Queryx(fmt.Sprintf(
		"select * from users where id='%s';",
		id,
	))

	util.CheckErr(err)
	defer result.Close()

	var user users.User
	result.Next()
	err = result.StructScan(&user)

	util.CheckErr(err)
	return user
}

func (d *Db) ReadAllUsers() []users.User {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	result, err := d.Client.Db.Queryx("select * from users;")
	util.CheckErr(err)
	defer result.Close()

	us := make([]users.User, 0)
	var u users.User
	for i := 0; result.Next(); i++ {
		err := result.StructScan(&u)
		if err != nil {
			panic(err.Error())
		}
		us = append(us, u)
	}

	return us
}

func (d *Db) ReadTop() []users.User {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	results, err := d.Client.Db.Queryx("select * from users order by elo, id asc limit 0,10;")
	util.CheckErr(err)
	defer results.Close()

	tt := make([]users.User, 0)
	var u users.User
	for i := 0; results.Next(); i++ {
		err := results.StructScan(&u)
		util.CheckErr(err)
		tt = append(tt, u)
	}

	return tt
}

func (d *Db) ReadStats(id string) Stats {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	result, err := d.Client.Db.Queryx(fmt.Sprintf(
		"select wins, losses, draws, games, elo, level from users where id='%s';",
		id))

	util.CheckErr(err)
	defer result.Close()

	var stats Stats
	result.Next()
	err = result.StructScan(&stats)
	util.CheckErr(err)

	return stats
}

func (d *Db) CheckUser(id string) bool {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	err := d.Client.Db.QueryRow(fmt.Sprintf(
		"select id from users where id='%s'",
		id)).Scan(&id)

	exists, err := util.CheckRow(err)
	util.CheckErr(err)

	return exists
}

func (d *Db) DeleteUser(id string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	query := fmt.Sprintf(
		"id='%s'",
		id,
	)

	delete, err := d.Client.Db.Query(fmt.Sprintf(
		"delete from users where %s;",
		query,
	))

	util.CheckErr(err)
	defer delete.Close()
}

func (d *Db) DeleteUsers(ids []string) {
	for _, id := range ids {
		d.DeleteUser(id)
	}
}
