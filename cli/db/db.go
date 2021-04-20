package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var _db *sqlx.DB

func init() {
	db, err := sqlx.Connect("mysql", "root:root@(localhost:3306)/dev?parseTime=true")
	if err != nil {
		panic(err)
	}
	_db = db
}

func DB() *sqlx.DB {
	return _db
}
