package msql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Username string
	Password string
	Type     string
	Host     string
	Port     int
	Database string
}

func NewDB(c *Config) *sqlx.DB {
	db, err := sqlx.Connect(c.Type, fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", c.Username, c.Password, c.Host,
		c.Port, c.Database))
	if err != nil {
		panic(err)
	}
	return db
}
