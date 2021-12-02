package test

import (
	"github.com/volatiletech/null/v9"
	"time"
)

var (
	_ time.Time
	_ null.Bool
)

type Gomodel struct {
	Tinyint    int8         `json:"tinyint" db:"tinyint"`
	Smallint   null.Int16   `json:"smallint" db:"smallint"`
	Mediumint  null.Int32   `json:"mediumint" db:"mediumint"`
	Int        null.Int     `json:"int" db:"int"`
	Bigint     null.Int64   `json:"bigint" db:"bigint"`
	Float      null.Float32 `json:"float" db:"float"`
	Double     null.Float64 `json:"double" db:"double"`
	Decimal    null.Float64 `json:"decimal" db:"decimal"`
	Utinyint   null.Uint8   `json:"utinyint" db:"utinyint"`
	Usmallint  null.Uint16  `json:"usmallint" db:"usmallint"`
	Umediumint null.Uint32  `json:"umediumint" db:"umediumint"`
	Uint       null.Uint    `json:"uint" db:"uint"`
	Ubigint    null.Uint64  `json:"ubigint" db:"ubigint"`
	Ufloat     null.Float32 `json:"ufloat" db:"ufloat"`
	Udouble    null.Float64 `json:"udouble" db:"udouble"`
	Udecimal   null.Float64 `json:"udecimal" db:"udecimal"`
	Date       null.Time    `json:"date" db:"date"`
	Datetime   null.Time    `json:"datetime" db:"datetime"`
	Timestamp  time.Time    `json:"timestamp" db:"timestamp"`
	Time       null.Time    `json:"time" db:"time"`
	Year       null.Time    `json:"year" db:"year"`
	Char       null.String  `json:"char" db:"char"`
	Varchar    null.String  `json:"varchar" db:"varchar"`
	Binary     null.String  `json:"binary" db:"binary"`
	Varbinary  null.String  `json:"varbinary" db:"varbinary"`
	Tinyblob   null.String  `json:"tinyblob" db:"tinyblob"`
	Tinytext   null.String  `json:"tinytext" db:"tinytext"`
	Blob       null.String  `json:"blob" db:"blob"`
	Text       null.String  `json:"text" db:"text"`
	Mediumblob null.String  `json:"mediumblob" db:"mediumblob"`
	Mediumtext null.String  `json:"mediumtext" db:"mediumtext"`
	Longblob   null.String  `json:"longblob" db:"longblob"`
	Longtext   null.String  `json:"longtext" db:"longtext"`
	Enum       null.String  `json:"enum" db:"enum"`
	Set        null.String  `json:"set" db:"set"`
	Json       null.String  `json:"json" db:"json"`
	Tinybool   null.Bool    `json:"tinybool" db:"tinybool"`
	Bool       null.Bool    `json:"bool" db:"bool"`
}
