package test

import "context"

type gomodelInsertEvent struct {
	*GomodelInsertBuilder
	ctx   context.Context
	sql   string
	args  []interface{}
	value interface{}
}

func newGomodelInsertEvent(ctx context.Context, builder *GomodelInsertBuilder) *gomodelInsertEvent {
	sql, args := builder.SQL()
	return &gomodelInsertEvent{
		GomodelInsertBuilder: builder,
		ctx:                  ctx,
		sql:                  sql,
		args:                 args,
	}
}

func (e *gomodelInsertEvent) Context() context.Context {
	return e.ctx
}

func (e *gomodelInsertEvent) SetContext(ctx context.Context) {
	e.ctx = ctx
}

func (e *gomodelInsertEvent) Table() string {
	return e.table
}

func (e *gomodelInsertEvent) SetTable(table string) {
	e.table = table
	e.sql, _ = e.GomodelInsertBuilder.SQL()
}

func (e *gomodelInsertEvent) Op() Op {
	return OpInsert
}

func (e *gomodelInsertEvent) SQL() string {
	return e.sql
}

func (e *gomodelInsertEvent) Args() []interface{} {
	return e.args
}

func (e *gomodelInsertEvent) Value() interface{} {
	return e.value
}

func (e *gomodelInsertEvent) SetValue(value interface{}) {
	e.value = value
}

type gomodelSelectEvent struct {
	*GomodelSelectBuilder
	ctx   context.Context
	sql   string
	args  []interface{}
	value interface{}
}

func newGomodelSelectEvent(ctx context.Context, builder *GomodelSelectBuilder) *gomodelSelectEvent {
	sql, args := builder.SQL()
	return &gomodelSelectEvent{
		GomodelSelectBuilder: builder,
		ctx:                  ctx,
		sql:                  sql,
		args:                 args,
	}
}

func (e *gomodelSelectEvent) Context() context.Context {
	return e.ctx
}

func (e *gomodelSelectEvent) SetContext(ctx context.Context) {
	e.ctx = ctx
}

func (e *gomodelSelectEvent) Table() string {
	return e.table
}

func (e *gomodelSelectEvent) SetTable(table string) {
	e.table = table
	e.sql, _ = e.GomodelSelectBuilder.SQL()
}

func (e *gomodelSelectEvent) Op() Op {
	return OpSelect
}

func (e *gomodelSelectEvent) SQL() string {
	return e.sql
}

func (e *gomodelSelectEvent) Args() []interface{} {
	return e.args
}

func (e *gomodelSelectEvent) Value() interface{} {
	return e.value
}

func (e *gomodelSelectEvent) SetValue(value interface{}) {
	e.value = value
}

type gomodelUpdateEvent struct {
	*GomodelUpdateBuilder
	ctx   context.Context
	sql   string
	args  []interface{}
	value interface{}
}

func newGomodelUpdateEvent(ctx context.Context, builder *GomodelUpdateBuilder) *gomodelUpdateEvent {
	sql, args := builder.SQL()
	return &gomodelUpdateEvent{
		GomodelUpdateBuilder: builder,
		ctx:                  ctx,
		sql:                  sql,
		args:                 args,
	}
}

func (e *gomodelUpdateEvent) Context() context.Context {
	return e.ctx
}

func (e *gomodelUpdateEvent) SetContext(ctx context.Context) {
	e.ctx = ctx
}

func (e *gomodelUpdateEvent) Table() string {
	return e.table
}

func (e *gomodelUpdateEvent) SetTable(table string) {
	e.table = table
	e.sql, _ = e.GomodelUpdateBuilder.SQL()
}

func (e *gomodelUpdateEvent) Op() Op {
	return OpUpdate
}

func (e *gomodelUpdateEvent) SQL() string {
	return e.sql
}

func (e *gomodelUpdateEvent) Args() []interface{} {
	return e.args
}

func (e *gomodelUpdateEvent) Value() interface{} {
	return e.value
}

func (e *gomodelUpdateEvent) SetValue(value interface{}) {
	e.value = value
}

type gomodeldeleteEvent struct {
	*GomodelDeleteBuilder
	ctx   context.Context
	sql   string
	args  []interface{}
	value interface{}
}

func newGomodeldeleteEvent(ctx context.Context, builder *GomodelDeleteBuilder) *gomodeldeleteEvent {
	sql, args := builder.SQL()
	return &gomodeldeleteEvent{
		GomodelDeleteBuilder: builder,
		ctx:                  ctx,
		sql:                  sql,
		args:                 args,
	}
}

func (e *gomodeldeleteEvent) Context() context.Context {
	return e.ctx
}

func (e *gomodeldeleteEvent) SetContext(ctx context.Context) {
	e.ctx = ctx
}

func (e *gomodeldeleteEvent) Table() string {
	return e.table
}

func (e *gomodeldeleteEvent) SetTable(table string) {
	e.table = table
	e.sql, _ = e.GomodelDeleteBuilder.SQL()
}

func (e *gomodeldeleteEvent) Op() Op {
	return OpDelete
}

func (e *gomodeldeleteEvent) SQL() string {
	return e.sql
}

func (e *gomodeldeleteEvent) Args() []interface{} {
	return e.args
}

func (e *gomodeldeleteEvent) Value() interface{} {
	return e.value
}

func (e *gomodeldeleteEvent) SetValue(value interface{}) {
	e.value = value
}

func (e *gomodelInsertEvent) Fields() []string {
	return nil
}

func (e *gomodelSelectEvent) Fields() []string {
	return e.fields
}

func (e *gomodelUpdateEvent) Fields() []string {
	return nil
}

func (e *gomodeldeleteEvent) Fields() []string {
	return nil
}
