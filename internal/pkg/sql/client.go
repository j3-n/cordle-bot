package sql

import (
	"cordle/internal/config"
	"cordle/internal/pkg/util"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Clientx struct {
	Db *sqlx.DB
}

type Client struct {
	Db *sql.DB
}

func NewClientx(c config.SqlConfig) *Clientx {
	d, err := sqlx.Open("mysql", connStr(c))
	util.CheckErr(err)

	return &Clientx{
		Db: d,
	}
}

func NewClient(c config.SqlConfig) *Client {
	d, err := sql.Open("mysql", connStr(c))
	util.CheckErr(err)

	return &Client{
		Db: d,
	}
}

func (c *Clientx) Close() error {
	err := c.Db.Close()
	return err
}

func connStr(c config.SqlConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.Username,
		c.Password,
		c.Address,
		c.Port,
		c.Database)
}
