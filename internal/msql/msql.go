package msql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"regexp"
	"strconv"
	"strings"
)

type DB struct {
	*sqlx.DB
}

type Column struct {
	Field string
	// Type eg: "varchar" | "bigint"
	Type       string
	Unsigned   bool
	Len        int
	Nullable   bool
	Default    *string
	Privileges string
	Comment    string
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

	type rawColumn struct {
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

	columns := make([]*Column, 0)
	for rows.Next() {
		c := &rawColumn{}
		if err := rows.StructScan(c); err != nil {
			return nil, err
		}

		typ := strings.ToLower(c.Type)
		if bracketIdx := strings.Index(typ, "("); bracketIdx != -1 {
			typ = typ[0:bracketIdx]
		}
		if spaceIdx := strings.Index(typ, " "); spaceIdx != -1 {
			typ = typ[0:spaceIdx]
		}

		var l int
		reg := regexp.MustCompile(".*?\\((\\d+)\\).*")
		submatch := reg.FindStringSubmatch(c.Type)
		if len(submatch) > 1 {
			l, _ = strconv.Atoi(submatch[1])
		}

		columns = append(columns, &Column{
			Field:      c.Field,
			Type:       typ,
			Unsigned:   strings.Contains(c.Type, "unsigned"),
			Len:        l,
			Nullable:   c.Null == "YES",
			Default:    c.Default,
			Privileges: c.Privileges,
			Comment:    c.Comment,
		})
	}

	return columns, nil
}
