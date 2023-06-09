package sql

import (
	"cordle/internal/pkg/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	Db sqlx.DB
}

func NewClient(connStr string) *Client {
	d, err := sqlx.Open("mysql", connStr)
	util.CheckErr(err)

	return &Client{
		Db: *d,
	}
}

func (c *Client) Close() error {
	err := c.Db.Close()
	return err
}
