package test

import (
	"github.com/jmoiron/sqlx"
)

type GomodelDB struct {
	ext   sqlx.ExtContext
	hooks []Hook
}

func NewGomodelDB(db sqlx.ExtContext) *GomodelDB {
	return &GomodelDB{
		ext: db,
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

func (db *GomodelDB) runBeforeHooks(info *queryInfo) error {
	for _, hook := range db.hooks {
		err := hook.Before(info)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *GomodelDB) runAfterHooks(info *queryInfo) error {
	for _, hook := range db.hooks {
		err := hook.After(info)
		if err != nil {
			return err
		}
	}
	return nil
}
