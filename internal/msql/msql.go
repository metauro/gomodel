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

type DB struct {
	*sqlx.DB
}

type Column struct {
	Field      string  `db:"Field"`
	Type       string  `db:"Type"`
	Collation  *string `db:"Collation"`
	Null       string  `db:"Null"`
	Key        string  `db:"Key"`
	Default    *string `db:"Default"`
	Extra      string  `db:"Extra"`
	Privileges string  `db:"Privileges"`
	Comment    string  `db:"Comment"`
}

func Open(driveName, dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect(driveName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) GetTables() ([]string, error) {
	rows, err := db.Queryx("SHOW TABLES")
	if err != nil {
		return nil, err
	}

	tables := make([]string, 0)
	for rows.Next() {
		var str string
		if err := rows.Scan(&str); err != nil {
			panic(err)
		}
		tables = append(tables, str)
	}

	return tables, nil
}

func (db *DB) GetColumns(table string) ([]*Column, error) {
	rows, err := db.Queryx(fmt.Sprintf("SHOW FULL COLUMNS FROM `%s`", table))
	if err != nil {
		panic(err)
	}

	columns := make([]*Column, 0)
	for rows.Next() {
		c := &Column{}
		if err := rows.StructScan(c); err != nil {
			return nil, err
		}
		columns = append(columns, c)
	}

	return columns, nil
}
