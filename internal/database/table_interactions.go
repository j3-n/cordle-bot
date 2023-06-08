package database

import (
	"cordle/internal/pkg/util"
	"fmt"
)

func (d *Db) CreateTable(s string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	q, err := db.Query(s)
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) UpdateTable(u string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	q, err := db.Query(u)
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) DeleteTable(t string) {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	q, err := db.Query(fmt.Sprintf(
		"drop table %s", 
		t,
	))
	util.CheckErr(err)
	defer q.Close()
}

func (d *Db) CheckTable(t string) bool {
	d.ClientMu.Lock()
	defer d.ClientMu.Unlock()
	db := d.Client.Db

	q, err := db.Query(fmt.Sprintf(
		"drop table %s", 
		t,
	))
	e, err := util.CheckRow(err)
	util.CheckErr(err)
	defer q.Close()

	return e
}
