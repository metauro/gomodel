package test

import (
	"database/sql"
)

type Op string

const (
	OpInsert Op = "insert"
	OpUpdate Op = "update"
	OpSelect Op = "select"
	OpDelete Op = "delete"
)

type set struct {
	sql string
	arg interface{}
}

type ErrorHandler func(err error) error

var ErrNoRows = sql.ErrNoRows

func defaultErrorHandler(err error) error {
	return err
}

type _Builder interface {
	SQL() (string, []interface{})
}
