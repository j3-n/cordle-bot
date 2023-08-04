package database

import (
	"cordle/internal/pkg/util"
	"cordle/internal/users"
	"fmt"
)

func (d *Db) AddUser(id string) error {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	insert, err := d.client.Db.Query(fmt.Sprintf(
		`insert into users(id)
		values(%s);`,
		id,
	))

	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}

func (d *Db) AddUsers(users *[]users.User) error {
	for _, user := range *users {
		err := d.AddUser(user.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Db) UpdateUser(user *users.User) error {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	q := fmt.Sprintf(
		"id='%s'",
		user.Id,
	)

	update, err := d.client.Db.Query(fmt.Sprintf(
		`update users
		 set %s
		 where %s;`,
		user.ToSqlUpdate(),
		q,
	))

	if err != nil {
		return err
	}

	defer update.Close()
	return nil
}

func (d *Db) UpdateUsers(users *[]users.User) error {
	for _, user := range *users {
		err := d.UpdateUser(&user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Db) ReadUser(id string) (*users.User, error) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	result, err := d.client.Db.Queryx(fmt.Sprintf(
		"select * from users where id='%s';",
		id,
	))

	if err != nil {
		return &users.User{}, err
	}
	defer result.Close()

	var user users.User
	result.Next()
	err = result.StructScan(&user)

	if err != nil {
		return &users.User{}, err
	}

	return &user, nil
}

func (d *Db) ReadAllUsers() ([]*users.User, error) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	result, err := d.client.Db.Queryx("select * from users;")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	us := make([]*users.User, 0)
	var u users.User
	for i := 0; result.Next(); i++ {
		err := result.StructScan(&u)
		if err != nil {
			return nil, err
		}
		us = append(us, &u)
	}

	return us, nil
}

func (d *Db) ReadTop() ([]*users.User, error) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	results, err := d.client.Db.Queryx("select * from users order by elo, id asc limit 0,10;")
	if err != nil {
		return nil, err
	}
	defer results.Close()

	tt := make([]*users.User, 0)
	var u users.User
	for i := 0; results.Next(); i++ {
		err := results.StructScan(&u)
		if err != nil {
			return nil, err
		}
		tt = append(tt, &u)
	}

	return tt, nil
}

func (d *Db) ReadStats(id string) (*Stats, error) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	result, err := d.client.Db.Queryx(fmt.Sprintf(
		"select wins, losses, draws, games, elo, level from users where id='%s';",
		id))

	if err != nil {
		return &Stats{}, err
	}
	defer result.Close()

	var stats Stats
	result.Next()
	err = result.StructScan(&stats)
	if err != nil {
		return &Stats{}, err
	}

	return &stats, nil
}

func (d *Db) CheckUser(id string) (bool, error) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	err := d.client.Db.QueryRow(fmt.Sprintf(
		"select id from users where id='%s'",
		id)).Scan(&id)

	exists, err := util.CheckRow(err)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (d *Db) DeleteUser(id string) error {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	query := fmt.Sprintf(
		"id='%s'",
		id,
	)

	delete, err := d.client.Db.Query(fmt.Sprintf(
		"delete from users where %s;",
		query,
	))

	if err != nil {
		return err
	}

	defer delete.Close()
	return nil
}

func (d *Db) DeleteUsers(ids []string) error {
	for _, id := range ids {
		err := d.DeleteUser(id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Db) Ping() error {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	return d.client.Db.Ping()
}
