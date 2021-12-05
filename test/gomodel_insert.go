package test

import (
	"context"
	"strings"
	"time"
)

var _ = time.Second

type GomodelInsertBuilder struct {
	db     *GomodelDB
	table  string
	values []*Gomodel
	fields []string
}

func newGomodelInsertBuilder(db *GomodelDB) *GomodelInsertBuilder {
	return &GomodelInsertBuilder{
		db:     db,
		table:  GomodelTable,
		fields: GomodelFields,
	}
}

// Fields 指定只插入结构体哪些字段，默认全部字段都插入
func (b *GomodelInsertBuilder) Fields(fields ...string) *GomodelInsertBuilder {
	b.fields = fields
	return b
}

// Values .
func (b *GomodelInsertBuilder) Values(gomodelList ...*Gomodel) *GomodelInsertBuilder {
	b.values = append(b.values, gomodelList...)
	return b
}

func (b *GomodelInsertBuilder) SQL() (string, []interface{}) {
	if len(b.values) == 0 {
		return "", nil
	}

	var sb strings.Builder
	args := make([]interface{}, 0, len(b.fields)*len(b.values))

	sb.WriteString("INSERT INTO ")
	sb.WriteString("`")
	sb.WriteString(b.table)
	sb.WriteString("` (")

	sb.WriteString(b.fields[0])
	for i, size := 1, len(b.fields); i < size; i++ {
		sb.WriteString(",")
		sb.WriteString(b.fields[i])
	}

	sb.WriteString(") VALUES (?")
	for i, size := 1, len(b.fields); i < size; i++ {
		sb.WriteString(",")
		sb.WriteString("?")
	}
	sb.WriteString(")")

	if len(b.values) > 1 {
		for i, size := 1, len(b.values); i < size; i++ {
			sb.WriteString(",(?")
			for j, fsize := 1, len(b.fields); j < fsize; j++ {
				sb.WriteString(",?")
			}
			sb.WriteString(")")
		}
	}

	for i, size := 0, len(b.values); i < size; i++ {
		for _, field := range b.fields {
			switch field {
			case GomodelFieldTinyint:
				args = append(args, b.values[i].Tinyint)
			case GomodelFieldSmallint:
				args = append(args, b.values[i].Smallint)
			case GomodelFieldMediumint:
				args = append(args, b.values[i].Mediumint)
			case GomodelFieldInt:
				args = append(args, b.values[i].Int)
			case GomodelFieldBigint:
				args = append(args, b.values[i].Bigint)
			case GomodelFieldFloat:
				args = append(args, b.values[i].Float)
			case GomodelFieldDouble:
				args = append(args, b.values[i].Double)
			case GomodelFieldDecimal:
				args = append(args, b.values[i].Decimal)
			case GomodelFieldUtinyint:
				args = append(args, b.values[i].Utinyint)
			case GomodelFieldUsmallint:
				args = append(args, b.values[i].Usmallint)
			case GomodelFieldUmediumint:
				args = append(args, b.values[i].Umediumint)
			case GomodelFieldUint:
				args = append(args, b.values[i].Uint)
			case GomodelFieldUbigint:
				args = append(args, b.values[i].Ubigint)
			case GomodelFieldUfloat:
				args = append(args, b.values[i].Ufloat)
			case GomodelFieldUdouble:
				args = append(args, b.values[i].Udouble)
			case GomodelFieldUdecimal:
				args = append(args, b.values[i].Udecimal)
			case GomodelFieldDate:
				args = append(args, b.values[i].Date)
			case GomodelFieldDatetime:
				args = append(args, b.values[i].Datetime)
			case GomodelFieldTimestamp:
				args = append(args, b.values[i].Timestamp)
			case GomodelFieldTime:
				args = append(args, b.values[i].Time)
			case GomodelFieldYear:
				args = append(args, b.values[i].Year)
			case GomodelFieldChar:
				args = append(args, b.values[i].Char)
			case GomodelFieldVarchar:
				args = append(args, b.values[i].Varchar)
			case GomodelFieldBinary:
				args = append(args, b.values[i].Binary)
			case GomodelFieldVarbinary:
				args = append(args, b.values[i].Varbinary)
			case GomodelFieldTinyblob:
				args = append(args, b.values[i].Tinyblob)
			case GomodelFieldTinytext:
				args = append(args, b.values[i].Tinytext)
			case GomodelFieldBlob:
				args = append(args, b.values[i].Blob)
			case GomodelFieldText:
				args = append(args, b.values[i].Text)
			case GomodelFieldMediumblob:
				args = append(args, b.values[i].Mediumblob)
			case GomodelFieldMediumtext:
				args = append(args, b.values[i].Mediumtext)
			case GomodelFieldLongblob:
				args = append(args, b.values[i].Longblob)
			case GomodelFieldLongtext:
				args = append(args, b.values[i].Longtext)
			case GomodelFieldEnum:
				args = append(args, b.values[i].Enum)
			case GomodelFieldSet:
				args = append(args, b.values[i].Set)
			case GomodelFieldJson:
				args = append(args, b.values[i].Json)
			case GomodelFieldTinybool:
				args = append(args, b.values[i].Tinybool)
			case GomodelFieldBool:
				args = append(args, b.values[i].Bool)
			}
		}
	}

	return sb.String(), args
}

func (b *GomodelInsertBuilder) Exec(ctx context.Context) (int64, error) {
	var id int64
	e := newGomodelInsertEvent(ctx, b)
	return id, b.db.exec(e, func(ctx context.Context, sql string, args ...interface{}) (interface{}, error) {
		res, err := b.db.ext.ExecContext(ctx, sql, args...)
		if err != nil {
			return 0, err
		}
		id, err = res.LastInsertId()
		return id, err
	})
}
