package database

import (
	"cordle/internal/pkg/util"
	"fmt"
)

func (d *Db) CreateTable(s string) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	q, err := d.client.Db.Query(s)
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) UpdateTable(u string) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	q, err := d.client.Db.Query(u)
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) DeleteTable(t string) {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	q, err := d.client.Db.Query(fmt.Sprintf(
		"drop table %s;",
		t,
	))
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) CheckTable(t string) bool {
	d.clientMu.Lock()
	defer d.clientMu.Unlock()

	_, err := d.client.Db.Query(fmt.Sprintf(
		"select * from %s limit 0,1;",
		t,
	))
	e, err := util.CheckTable(err)
	util.CheckErr(err)

	return e
}
