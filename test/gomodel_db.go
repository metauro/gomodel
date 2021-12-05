package test

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type GomodelDB struct {
	ext         sqlx.ExtContext
	hooks       []Hook
	handleError ErrorHandler
}

func NewGomodelDB(db sqlx.ExtContext) *GomodelDB {
	return &GomodelDB{
		ext:         db,
		handleError: defaultErrorHandler,
	}
}

func (db *GomodelDB) Insert() *GomodelInsertBuilder {
	return newGomodelInsertBuilder(db)
}

func (db *GomodelDB) Update() *GomodelUpdateBuilder {
	return newGomodelUpdateBuilder(db)
}

func (db *GomodelDB) Select() *GomodelSelectBuilder {
	return newGomodelSelectBuilder(db)
}

func (db *GomodelDB) Delete() *GomodelDeleteBuilder {
	return newGomodelDeleteBuilder(db)
}

func (db *GomodelDB) Use(hooks ...Hook) {
	for _, hook := range hooks {
		db.hooks = append(db.hooks, hook)
	}
}

func (db *GomodelDB) SetErrorHandler(handler ErrorHandler) {
	db.handleError = handler
}

func (db *GomodelDB) exec(
	e event,
	execute func(ctx context.Context, sql string, args ...interface{}) (interface{}, error),
) error {
	for _, hook := range db.hooks {
		err := hook.Before(e)
		if err != nil {
			return err
		}
	}

	value, err := execute(e.Context(), e.SQL(), e.Args()...)
	if err != nil {
		return db.handleError(err)
	}

	e.SetValue(value)
	for _, hook := range db.hooks {
		err := hook.After(e)
		if err != nil {
			return err
		}
	}

	return nil
}
