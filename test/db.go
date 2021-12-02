package test

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	db      *sqlx.DB
	Gomodel *GomodelDB
}

type Tx struct {
	tx      *sqlx.Tx
	Gomodel *GomodelDB
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{
		Gomodel: NewGomodelDB(db),
	}
}

func NewTx(tx *sqlx.Tx) *Tx {
	return &Tx{
		Gomodel: NewGomodelDB(tx),
	}
}

func (db *DB) BeginFn(ctx context.Context, fn func(ctx context.Context, tx *Tx) error) error {
	return db.BeginTxFn(ctx, nil, fn)
}

func (db *DB) BeginTxFn(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context, tx *Tx) error) error {
	tx, err := db.db.BeginTxx(ctx, opts)
	if err != nil {
		return err
	}

	txDB := NewTx(tx)
	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
			panic(err)
		}
	}()

	err = fn(ctx, txDB)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil && err != sql.ErrTxDone {
		return err
	}

	return nil
}
