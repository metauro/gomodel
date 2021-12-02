package test

import (
	"context"
	"strings"
)

type GomodelDeleteBuilder struct {
	db    *GomodelDB
	table string
	where *GomodelWhereBuilder
}

func newGomodelDeleteBuilder(db *GomodelDB) *GomodelDeleteBuilder {
	return &GomodelDeleteBuilder{
		db:    db,
		table: GomodelTable,
	}
}

func (b *GomodelDeleteBuilder) Where(fn func(b *GomodelWhereBuilder)) *GomodelDeleteBuilder {
	if b.where == nil {
		b.where = newGomodelWhereBuilder()
	}
	fn(b.where)
	return b
}

func (b *GomodelDeleteBuilder) SQL() (string, []interface{}) {
	var sb strings.Builder
	sb.WriteString("DELETE FROM ")
	sb.WriteString("`")
	sb.WriteString(b.table)
	sb.WriteString("`")

	if b.where != nil {
		sb.WriteString(" ")
		whereSQL, whereArgs := b.where.sql()
		sb.WriteString(whereSQL)
		return sb.String(), whereArgs
	}

	return sb.String(), nil
}

func (b *GomodelDeleteBuilder) Exec(ctx context.Context) (int64, error) {
	sql, args := b.SQL()
	info := &queryInfo{
		ctx:   ctx,
		table: b.table,
		op:    OpDelete,
		query: sql,
		args:  args,
	}
	err := b.db.runBeforeHooks(info)
	if err != nil {
		return 0, err
	}
	ctx = info.ctx

	if info.modified {
		b.table = info.table
		sql, _ = b.SQL()
		args = info.args
	}

	res, err := b.db.ext.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	info.value = ra
	return ra, b.db.runAfterHooks(info)
}
