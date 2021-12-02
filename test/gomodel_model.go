package test

import "time"

var _ = time.Second

type Gomodel struct {
	Tinyint    int8      `json:"tinyint" db:"tinyint"`
	Smallint   int16     `json:"smallint" db:"smallint"`
	Mediumint  int32     `json:"mediumint" db:"mediumint"`
	Int        int32     `json:"int" db:"int"`
	Bigint     int64     `json:"bigint" db:"bigint"`
	Float      float32   `json:"float" db:"float"`
	Double     float64   `json:"double" db:"double"`
	Decimal    float64   `json:"decimal" db:"decimal"`
	Utinyint   uint8     `json:"utinyint" db:"utinyint"`
	Usmallint  uint16    `json:"usmallint" db:"usmallint"`
	Umediumint uint32    `json:"umediumint" db:"umediumint"`
	Uint       uint32    `json:"uint" db:"uint"`
	Ubigint    uint64    `json:"ubigint" db:"ubigint"`
	Ufloat     float32   `json:"ufloat" db:"ufloat"`
	Udouble    float64   `json:"udouble" db:"udouble"`
	Udecimal   float64   `json:"udecimal" db:"udecimal"`
	Date       time.Time `json:"date" db:"date"`
	Datetime   time.Time `json:"datetime" db:"datetime"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
	Time       time.Time `json:"time" db:"time"`
	Year       time.Time `json:"year" db:"year"`
	Char       string    `json:"char" db:"char"`
	Varchar    string    `json:"varchar" db:"varchar"`
	Binary     string    `json:"binary" db:"binary"`
	Varbinary  string    `json:"varbinary" db:"varbinary"`
	Tinyblob   string    `json:"tinyblob" db:"tinyblob"`
	Tinytext   string    `json:"tinytext" db:"tinytext"`
	Blob       string    `json:"blob" db:"blob"`
	Text       string    `json:"text" db:"text"`
	Mediumblob string    `json:"mediumblob" db:"mediumblob"`
	Mediumtext string    `json:"mediumtext" db:"mediumtext"`
	Longblob   string    `json:"longblob" db:"longblob"`
	Longtext   string    `json:"longtext" db:"longtext"`
	Enum       string    `json:"enum" db:"enum"`
	Set        string    `json:"set" db:"set"`
	Json       string    `json:"json" db:"json"`
	Tinybool   bool      `json:"tinybool" db:"tinybool"`
	Bool       bool      `json:"bool" db:"bool"`
}
