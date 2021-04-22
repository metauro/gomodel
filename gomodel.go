package gomodel

import (
	"context"
	"database/sql"
)

type Execer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type DB struct {
	*sql.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) MustBegin() *sql.Tx {
	return db.MustBeginTx(context.Background(), nil)
}

func (db *DB) MustBeginTx(ctx context.Context, opts *sql.TxOptions) *sql.Tx {
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		panic(err)
	}
	return tx
}

func (db *DB) BeginFn(fn func(tx *sql.Tx) error) error {
	return db.BeginTxFn(context.Background(), nil, fn)
}

func (db *DB) BeginTxFn(ctx context.Context, opts *sql.TxOptions, fn func(tx *sql.Tx) error) (err error) {
	defer func() {
		if err == nil {
			err = recover().(error)
		}
	}()

	tx, err := db.BeginTx(ctx, opts)
	err = fn(tx)
	return
}

func (db *DB) MustBeginFn(fn func(tx *sql.Tx) error) {
	db.MustBeginTxFn(context.Background(), nil, fn)
}

func (db *DB) MustBeginTxFn(ctx context.Context, opts *sql.TxOptions, fn func(tx *sql.Tx) error) {
	if err := db.BeginTxFn(ctx, opts, fn); err != nil {
		panic(err)
	}
}
