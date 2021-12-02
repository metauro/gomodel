package test

import (
	"context"
	"database/sql"
	"github.com/metauro/gomodel"
)

type (
	Op   = gomodel.Op
	Hook = gomodel.Hook
)

const (
	OpInsert = gomodel.OpInsert
	OpUpdate = gomodel.OpUpdate
	OpSelect = gomodel.OpSelect
	OpDelete = gomodel.OpDelete
)

type set struct {
	sql string
	arg interface{}
}

var ErrNoRows = sql.ErrNoRows

type queryInfo struct {
	ctx      context.Context
	table    string
	fields   []string
	op       Op
	query    string
	args     []interface{}
	value    interface{}
	err      error
	modified bool
}

func (qi *queryInfo) Fields() []string {
	return qi.fields
}

func (qi *queryInfo) AddField(field string) {
	qi.modified = true
	qi.fields = append(qi.fields, field)
}

func (qi *queryInfo) AddArg(arg interface{}) {
	qi.modified = true
	qi.args = append(qi.args, arg)
}

func (qi *queryInfo) Context() context.Context {
	return qi.ctx
}

func (qi *queryInfo) SetContext(ctx context.Context) {
	qi.modified = true
	qi.ctx = ctx
}

func (qi *queryInfo) Table() string {
	return qi.table
}

func (qi *queryInfo) SetTable(table string) {
	qi.modified = true
	qi.table = table
}

func (qi *queryInfo) Op() Op {
	return qi.op
}

func (qi *queryInfo) Query() string {
	return qi.query
}

func (qi *queryInfo) Args() []interface{} {
	return qi.args
}

func (qi *queryInfo) Value() interface{} {
	return qi.value
}

func (qi *queryInfo) Err() error {
	return qi.err
}
