package sql

import (
	"cordle/internal/pkg/util"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Clientx struct {
	Db sqlx.DB
}

type Client struct {
	Db sql.DB
}

func NewClientx(connStr string) *Clientx {
	d, err := sqlx.Open("mysql", connStr)
	util.CheckErr(err)

	return &Clientx{
		Db: *d,
	}
}

func NewClient(connStr string) *Client {
	d, err := sql.Open("mysql", connStr)
	util.CheckErr(err)

	return &Client{
		Db: *d,
	}
}

func (c *Clientx) Close() error {
	err := c.Db.Close()
	return err
}
