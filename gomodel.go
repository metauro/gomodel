package gomodel

import "context"

type Op uint

const (
	OpInsert Op = 1 << iota // insert data
	OpUpdate                // update data
	OpSelect                // select data
	OpDelete                // delete data
)

//go:generate stringer -type Op

type Hook interface {
	Before(info QueryInfo) error
	After(info QueryInfo) error
}

type QueryInfo interface {
	Context() context.Context
	SetContext(ctx context.Context)
	Table() string
	SetTable(table string)
	Op() Op
	// Query return sql string
	Query() string
	// Fields return select/update/insert fields
	Fields() []string
	AddField(field string)
	// Args return sql args
	Args() []interface{}
	AddArg(arg interface{})
	// Value return sql execute result
	Value() interface{}
	// Err return sql execute error
	Err() error
}
