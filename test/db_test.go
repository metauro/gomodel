package test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *DB

func init() {
	sqlDB, err := sqlx.Open("mysql", "root:wpsepmysql@(localhost:3306)/dev?parseTime=true")
	if err != nil {
		panic(err)
	}
	db = NewDB(sqlDB)
}
