package test

import (
	"context"
	"strings"
)

type GomodelSelectBuilder struct {
	db       *GomodelDB
	table    string
	fields   []string
	distinct bool
	offset   int
	limit    int
	where    *GomodelWhereBuilder
	order    *GomodelOrderBuilder
}

func newGomodelSelectBuilder(db *GomodelDB) *GomodelSelectBuilder {
	return &GomodelSelectBuilder{
		db:     db,
		table:  GomodelTable,
		fields: GomodelFields,
	}
}

// Fields 指定查询哪些字段
func (b *GomodelSelectBuilder) Fields(fields ...string) *GomodelSelectBuilder {
	if len(fields) == 0 {
		return b
	}

	b.fields = fields
	return b
}

// OmitFields 指定查询时忽略哪些字段
func (b *GomodelSelectBuilder) OmitFields(fields ...string) *GomodelSelectBuilder {
	if len(fields) == 0 {
		return b
	}

	b.fields = make([]string, 0, len(GomodelFields)-len(fields))
omit:
	for _, field := range GomodelFields {
		for _, omitField := range fields {
			if field == omitField {
				continue omit
			}
		}
		b.fields = append(b.fields, field)
	}
	return b
}

// Distinct 是否忽略相同数据
func (b *GomodelSelectBuilder) Distinct(distinct bool) *GomodelSelectBuilder {
	b.distinct = distinct
	return b
}

// Limit 指定查询数量
func (b *GomodelSelectBuilder) Limit(limit int) *GomodelSelectBuilder {
	b.limit = limit
	return b
}

// Offset 指定偏移量
func (b *GomodelSelectBuilder) Offset(offset int) *GomodelSelectBuilder {
	b.offset = offset
	return b
}

// Where 指定查询条件
func (b *GomodelSelectBuilder) Where(fn func(b *GomodelWhereBuilder)) *GomodelSelectBuilder {
	if b.where == nil {
		b.where = newGomodelWhereBuilder()
	}
	fn(b.where)
	return b
}

// Order 指定排序规则
func (b *GomodelSelectBuilder) Order(fn func(b *GomodelOrderBuilder)) *GomodelSelectBuilder {
	if b.order == nil {
		b.order = newGomodelOrderBuilder()
	}
	fn(b.order)
	return b
}

func (b *GomodelSelectBuilder) SQL() (string, []interface{}) {
	var sb strings.Builder

	sb.WriteString("SELECT ")
	if b.distinct {
		sb.WriteString("DISTINCT ")
	}
	sb.WriteString(b.fields[0])
	for i, size := 1, len(b.fields); i < size; i++ {
		sb.WriteString(",")
		sb.WriteString(b.fields[i])
	}

	sb.WriteString(" FROM ")
	sb.WriteString("`")
	sb.WriteString(b.table)
	sb.WriteString("`")

	args := make([]interface{}, 0, 2)
	if b.where != nil {
		whereSQL, whereArgs := b.where.sql()
		args = make([]interface{}, 0, len(whereArgs)+2)
		sb.WriteString(" ")
		sb.WriteString(whereSQL)
		args = append(args, whereArgs...)
	}

	if b.order != nil {
		orderSQL := b.order.sql()
		sb.WriteString(" ")
		sb.WriteString(orderSQL)
	}

	if b.limit > 0 {
		sb.WriteString(" LIMIT ?")
		args = append(args, b.limit)
	}

	if b.offset > 0 {
		sb.WriteString(" OFFSET ?")
		args = append(args, b.offset)
	}

	return sb.String(), args
}

// Get 获取单条数据
func (b *GomodelSelectBuilder) Get(ctx context.Context) (*Gomodel, error) {
	originLimit := b.limit
	defer func() {
		b.limit = originLimit
	}()

	b.Limit(1)
	list, err := b.List(ctx)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, ErrNoRows
	}

	return list[0], nil
}

// List 获取多条数据
func (b *GomodelSelectBuilder) List(ctx context.Context) ([]*Gomodel, error) {
	sql, args := b.SQL()
	info := &queryInfo{
		ctx:    ctx,
		table:  b.table,
		op:     OpSelect,
		query:  sql,
		fields: b.fields,
		args:   args,
	}

	err := b.db.runBeforeHooks(info)
	if err != nil {
		return nil, err
	}
	if info.modified {
		b.fields = info.fields
		b.table = info.table
		ctx = info.ctx
		sql, _ = b.SQL()
		args = info.args
	}

	rows, err := b.db.ext.QueryxContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	res := make([]*Gomodel, 0)
	scanners := make([]interface{}, len(b.fields))
	for rows.Next() {
		m := &Gomodel{}
		for i, field := range b.fields {
			switch field {
			case "`tinyint`":
				scanners[i] = &m.Tinyint
			case "`smallint`":
				scanners[i] = &m.Smallint
			case "`mediumint`":
				scanners[i] = &m.Mediumint
			case "`int`":
				scanners[i] = &m.Int
			case "`bigint`":
				scanners[i] = &m.Bigint
			case "`float`":
				scanners[i] = &m.Float
			case "`double`":
				scanners[i] = &m.Double
			case "`decimal`":
				scanners[i] = &m.Decimal
			case "`utinyint`":
				scanners[i] = &m.Utinyint
			case "`usmallint`":
				scanners[i] = &m.Usmallint
			case "`umediumint`":
				scanners[i] = &m.Umediumint
			case "`uint`":
				scanners[i] = &m.Uint
			case "`ubigint`":
				scanners[i] = &m.Ubigint
			case "`ufloat`":
				scanners[i] = &m.Ufloat
			case "`udouble`":
				scanners[i] = &m.Udouble
			case "`udecimal`":
				scanners[i] = &m.Udecimal
			case "`date`":
				scanners[i] = &m.Date
			case "`datetime`":
				scanners[i] = &m.Datetime
			case "`timestamp`":
				scanners[i] = &m.Timestamp
			case "`time`":
				scanners[i] = &m.Time
			case "`year`":
				scanners[i] = &m.Year
			case "`char`":
				scanners[i] = &m.Char
			case "`varchar`":
				scanners[i] = &m.Varchar
			case "`binary`":
				scanners[i] = &m.Binary
			case "`varbinary`":
				scanners[i] = &m.Varbinary
			case "`tinyblob`":
				scanners[i] = &m.Tinyblob
			case "`tinytext`":
				scanners[i] = &m.Tinytext
			case "`blob`":
				scanners[i] = &m.Blob
			case "`text`":
				scanners[i] = &m.Text
			case "`mediumblob`":
				scanners[i] = &m.Mediumblob
			case "`mediumtext`":
				scanners[i] = &m.Mediumtext
			case "`longblob`":
				scanners[i] = &m.Longblob
			case "`longtext`":
				scanners[i] = &m.Longtext
			case "`enum`":
				scanners[i] = &m.Enum
			case "`set`":
				scanners[i] = &m.Set
			case "`json`":
				scanners[i] = &m.Json
			case "`tinybool`":
				scanners[i] = &m.Tinybool
			case "`bool`":
				scanners[i] = &m.Bool
			}
		}
		if err := rows.Scan(scanners...); err != nil {
			return nil, err
		}
		res = append(res, m)
	}

	info.value = res
	return res, b.db.runAfterHooks(info)
}

func (b *GomodelSelectBuilder) Count(ctx context.Context) (int64, error) {
	originFields := b.fields
	defer func() {
		b.fields = originFields
	}()

	b.Fields("COUNT(1) AS `count`")
	sql, args := b.SQL()

	info := &queryInfo{
		ctx:   ctx,
		table: b.table,
		op:    OpSelect,
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

	row := b.db.ext.QueryRowxContext(ctx, sql, args...)
	if row.Err() != nil {
		return 0, row.Err()
	}

	var count int64
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	info.value = count
	return count, b.db.runAfterHooks(info)
}

func (b *GomodelSelectBuilder) Page(ctx context.Context, page, pageSize int) ([]*Gomodel, int64, error) {
	originOffset := b.offset
	originLimit := b.limit
	defer func() {
		b.offset = originOffset
		b.limit = originLimit
	}()
	b.offset = page * pageSize
	b.limit = pageSize

	list, err := b.List(ctx)
	if err != nil {
		return nil, 0, err
	}

	count, err := b.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}
