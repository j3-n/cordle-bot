package database

import (
	"cordle/internal/pkg/util"
	"fmt"
)

func (d *Db) CreateTable(s string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	q, err := d.Client.Db.Query(s)
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) UpdateTable(u string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	q, err := d.Client.Db.Query(u)
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) DeleteTable(t string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	q, err := d.Client.Db.Query(fmt.Sprintf(
		"drop table %s;",
		t,
	))
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) CheckTable(t string) bool {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()

	_, err := d.Client.Db.Query(fmt.Sprintf(
		"select * from %s limit 0,1;",
		t,
	))
	e, err := util.CheckTable(err)
	util.CheckErr(err)

	return e
}
