package test

import (
	"context"
	"strings"
	"time"
)

var _ = time.Second

type GomodelUpdateBuilder struct {
	db      *GomodelDB
	table   string
	changes map[string]*set
	where   *GomodelWhereBuilder
}

func newGomodelUpdateBuilder(db *GomodelDB) *GomodelUpdateBuilder {
	return &GomodelUpdateBuilder{
		db:      db,
		table:   GomodelTable,
		changes: make(map[string]*set),
	}
}

// Set 将数据更新为 gomodel 内的值,零值会被忽略
func (b *GomodelUpdateBuilder) Set(gomodel *Gomodel) *GomodelUpdateBuilder {
	if gomodel.Tinyint != 0 {
		b.set("Tinyint", "`tinyint`=?", gomodel.Tinyint)
	}

	if gomodel.Smallint.IsValid() {
		b.set("Smallint", "`smallint`=?", gomodel.Smallint.Int16)
	}

	if gomodel.Mediumint.IsValid() {
		b.set("Mediumint", "`mediumint`=?", gomodel.Mediumint.Int32)
	}

	if gomodel.Int.IsValid() {
		b.set("Int", "`int`=?", gomodel.Int.Int)
	}

	if gomodel.Bigint.IsValid() {
		b.set("Bigint", "`bigint`=?", gomodel.Bigint.Int64)
	}

	if gomodel.Float.IsValid() {
		b.set("Float", "`float`=?", gomodel.Float.Float32)
	}

	if gomodel.Double.IsValid() {
		b.set("Double", "`double`=?", gomodel.Double.Float64)
	}

	if gomodel.Decimal.IsValid() {
		b.set("Decimal", "`decimal`=?", gomodel.Decimal.Float64)
	}

	if gomodel.Utinyint.IsValid() {
		b.set("Utinyint", "`utinyint`=?", gomodel.Utinyint.Uint8)
	}

	if gomodel.Usmallint.IsValid() {
		b.set("Usmallint", "`usmallint`=?", gomodel.Usmallint.Uint16)
	}

	if gomodel.Umediumint.IsValid() {
		b.set("Umediumint", "`umediumint`=?", gomodel.Umediumint.Uint32)
	}

	if gomodel.Uint.IsValid() {
		b.set("Uint", "`uint`=?", gomodel.Uint.Uint)
	}

	if gomodel.Ubigint.IsValid() {
		b.set("Ubigint", "`ubigint`=?", gomodel.Ubigint.Uint64)
	}

	if gomodel.Ufloat.IsValid() {
		b.set("Ufloat", "`ufloat`=?", gomodel.Ufloat.Float32)
	}

	if gomodel.Udouble.IsValid() {
		b.set("Udouble", "`udouble`=?", gomodel.Udouble.Float64)
	}

	if gomodel.Udecimal.IsValid() {
		b.set("Udecimal", "`udecimal`=?", gomodel.Udecimal.Float64)
	}

	if gomodel.Date.IsValid() {
		b.set("Date", "`date`=?", gomodel.Date.Time)
	}

	if gomodel.Datetime.IsValid() {
		b.set("Datetime", "`datetime`=?", gomodel.Datetime.Time)
	}

	if !gomodel.Timestamp.IsZero() {
		b.set("Timestamp", "`timestamp`=?", gomodel.Timestamp)
	}

	if gomodel.Time.IsValid() {
		b.set("Time", "`time`=?", gomodel.Time.Time)
	}

	if gomodel.Year.IsValid() {
		b.set("Year", "`year`=?", gomodel.Year.Time)
	}

	if gomodel.Char.IsValid() {
		b.set("Char", "`char`=?", gomodel.Char.String)
	}

	if gomodel.Varchar.IsValid() {
		b.set("Varchar", "`varchar`=?", gomodel.Varchar.String)
	}

	if gomodel.Binary.IsValid() {
		b.set("Binary", "`binary`=?", gomodel.Binary.String)
	}

	if gomodel.Varbinary.IsValid() {
		b.set("Varbinary", "`varbinary`=?", gomodel.Varbinary.String)
	}

	if gomodel.Tinyblob.IsValid() {
		b.set("Tinyblob", "`tinyblob`=?", gomodel.Tinyblob.String)
	}

	if gomodel.Tinytext.IsValid() {
		b.set("Tinytext", "`tinytext`=?", gomodel.Tinytext.String)
	}

	if gomodel.Blob.IsValid() {
		b.set("Blob", "`blob`=?", gomodel.Blob.String)
	}

	if gomodel.Text.IsValid() {
		b.set("Text", "`text`=?", gomodel.Text.String)
	}

	if gomodel.Mediumblob.IsValid() {
		b.set("Mediumblob", "`mediumblob`=?", gomodel.Mediumblob.String)
	}

	if gomodel.Mediumtext.IsValid() {
		b.set("Mediumtext", "`mediumtext`=?", gomodel.Mediumtext.String)
	}

	if gomodel.Longblob.IsValid() {
		b.set("Longblob", "`longblob`=?", gomodel.Longblob.String)
	}

	if gomodel.Longtext.IsValid() {
		b.set("Longtext", "`longtext`=?", gomodel.Longtext.String)
	}

	if gomodel.Enum.IsValid() {
		b.set("Enum", "`enum`=?", gomodel.Enum.String)
	}

	if gomodel.Set.IsValid() {
		b.set("Set", "`set`=?", gomodel.Set.String)
	}

	if gomodel.Json.IsValid() {
		b.set("Json", "`json`=?", gomodel.Json.String)
	}

	if gomodel.Tinybool.IsValid() {
		b.set("Tinybool", "`tinybool`=?", gomodel.Tinybool.Bool)
	}

	if gomodel.Bool.IsValid() {
		b.set("Bool", "`bool`=?", gomodel.Bool.Bool)
	}

	return b
}

// SetTinyint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetTinyint(tinyint int8) *GomodelUpdateBuilder {
	b.set("`tinyint`", "`tinyint`=?", tinyint)
	return b
}

func (b *GomodelUpdateBuilder) IncrTinyint(val int8) *GomodelUpdateBuilder {
	b.set("`tinyint`", "`tinyint`=`tinyint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrTinyint(val int8) *GomodelUpdateBuilder {
	b.set("`tinyint`", "`tinyint`=`tinyint`-?", val)
	return b
}

// SetTinyintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTinyintZero() *GomodelUpdateBuilder {
	b.set("`tinyint`", "`tinyint`=?", 0)
	return b
}

// SetTinyintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTinyintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`tinyint`", "`tinyint` "+sql, args)
	return b
}

// SetSmallint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetSmallint(smallint int16) *GomodelUpdateBuilder {
	b.set("`smallint`", "`smallint`=?", smallint)
	return b
}

func (b *GomodelUpdateBuilder) IncrSmallint(val int16) *GomodelUpdateBuilder {
	b.set("`smallint`", "`smallint`=`smallint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrSmallint(val int16) *GomodelUpdateBuilder {
	b.set("`smallint`", "`smallint`=`smallint`-?", val)
	return b
}

// SetSmallintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetSmallintZero() *GomodelUpdateBuilder {
	b.set("`smallint`", "`smallint`=?", 0)
	return b
}

// SetSmallintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetSmallintNull() *GomodelUpdateBuilder {
	b.set("`smallint`", "`smallint`=NULL", nil)
	return b
}

// SetSmallintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetSmallintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`smallint`", "`smallint` "+sql, args)
	return b
}

// SetMediumint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetMediumint(mediumint int32) *GomodelUpdateBuilder {
	b.set("`mediumint`", "`mediumint`=?", mediumint)
	return b
}

func (b *GomodelUpdateBuilder) IncrMediumint(val int32) *GomodelUpdateBuilder {
	b.set("`mediumint`", "`mediumint`=`mediumint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrMediumint(val int32) *GomodelUpdateBuilder {
	b.set("`mediumint`", "`mediumint`=`mediumint`-?", val)
	return b
}

// SetMediumintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetMediumintZero() *GomodelUpdateBuilder {
	b.set("`mediumint`", "`mediumint`=?", 0)
	return b
}

// SetMediumintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetMediumintNull() *GomodelUpdateBuilder {
	b.set("`mediumint`", "`mediumint`=NULL", nil)
	return b
}

// SetMediumintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetMediumintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`mediumint`", "`mediumint` "+sql, args)
	return b
}

// SetInt 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetInt(int int) *GomodelUpdateBuilder {
	b.set("`int`", "`int`=?", int)
	return b
}

func (b *GomodelUpdateBuilder) IncrInt(val int) *GomodelUpdateBuilder {
	b.set("`int`", "`int`=`int`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrInt(val int) *GomodelUpdateBuilder {
	b.set("`int`", "`int`=`int`-?", val)
	return b
}

// SetIntZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetIntZero() *GomodelUpdateBuilder {
	b.set("`int`", "`int`=?", 0)
	return b
}

// SetIntNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetIntNull() *GomodelUpdateBuilder {
	b.set("`int`", "`int`=NULL", nil)
	return b
}

// SetIntRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetIntRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`int`", "`int` "+sql, args)
	return b
}

// SetBigint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetBigint(bigint int64) *GomodelUpdateBuilder {
	b.set("`bigint`", "`bigint`=?", bigint)
	return b
}

func (b *GomodelUpdateBuilder) IncrBigint(val int64) *GomodelUpdateBuilder {
	b.set("`bigint`", "`bigint`=`bigint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrBigint(val int64) *GomodelUpdateBuilder {
	b.set("`bigint`", "`bigint`=`bigint`-?", val)
	return b
}

// SetBigintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetBigintZero() *GomodelUpdateBuilder {
	b.set("`bigint`", "`bigint`=?", 0)
	return b
}

// SetBigintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetBigintNull() *GomodelUpdateBuilder {
	b.set("`bigint`", "`bigint`=NULL", nil)
	return b
}

// SetBigintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetBigintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`bigint`", "`bigint` "+sql, args)
	return b
}

// SetFloat 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetFloat(float float32) *GomodelUpdateBuilder {
	b.set("`float`", "`float`=?", float)
	return b
}

func (b *GomodelUpdateBuilder) IncrFloat(val float32) *GomodelUpdateBuilder {
	b.set("`float`", "`float`=`float`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrFloat(val float32) *GomodelUpdateBuilder {
	b.set("`float`", "`float`=`float`-?", val)
	return b
}

// SetFloatZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetFloatZero() *GomodelUpdateBuilder {
	b.set("`float`", "`float`=?", 0)
	return b
}

// SetFloatNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetFloatNull() *GomodelUpdateBuilder {
	b.set("`float`", "`float`=NULL", nil)
	return b
}

// SetFloatRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetFloatRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`float`", "`float` "+sql, args)
	return b
}

// SetDouble 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetDouble(double float64) *GomodelUpdateBuilder {
	b.set("`double`", "`double`=?", double)
	return b
}

func (b *GomodelUpdateBuilder) IncrDouble(val float64) *GomodelUpdateBuilder {
	b.set("`double`", "`double`=`double`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrDouble(val float64) *GomodelUpdateBuilder {
	b.set("`double`", "`double`=`double`-?", val)
	return b
}

// SetDoubleZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetDoubleZero() *GomodelUpdateBuilder {
	b.set("`double`", "`double`=?", 0)
	return b
}

// SetDoubleNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetDoubleNull() *GomodelUpdateBuilder {
	b.set("`double`", "`double`=NULL", nil)
	return b
}

// SetDoubleRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetDoubleRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`double`", "`double` "+sql, args)
	return b
}

// SetDecimal 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetDecimal(decimal float64) *GomodelUpdateBuilder {
	b.set("`decimal`", "`decimal`=?", decimal)
	return b
}

func (b *GomodelUpdateBuilder) IncrDecimal(val float64) *GomodelUpdateBuilder {
	b.set("`decimal`", "`decimal`=`decimal`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrDecimal(val float64) *GomodelUpdateBuilder {
	b.set("`decimal`", "`decimal`=`decimal`-?", val)
	return b
}

// SetDecimalZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetDecimalZero() *GomodelUpdateBuilder {
	b.set("`decimal`", "`decimal`=?", 0)
	return b
}

// SetDecimalNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetDecimalNull() *GomodelUpdateBuilder {
	b.set("`decimal`", "`decimal`=NULL", nil)
	return b
}

// SetDecimalRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetDecimalRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`decimal`", "`decimal` "+sql, args)
	return b
}

// SetUtinyint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUtinyint(utinyint uint8) *GomodelUpdateBuilder {
	b.set("`utinyint`", "`utinyint`=?", utinyint)
	return b
}

func (b *GomodelUpdateBuilder) IncrUtinyint(val uint8) *GomodelUpdateBuilder {
	b.set("`utinyint`", "`utinyint`=`utinyint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUtinyint(val uint8) *GomodelUpdateBuilder {
	b.set("`utinyint`", "`utinyint`=`utinyint`-?", val)
	return b
}

// SetUtinyintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUtinyintZero() *GomodelUpdateBuilder {
	b.set("`utinyint`", "`utinyint`=?", 0)
	return b
}

// SetUtinyintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUtinyintNull() *GomodelUpdateBuilder {
	b.set("`utinyint`", "`utinyint`=NULL", nil)
	return b
}

// SetUtinyintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUtinyintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`utinyint`", "`utinyint` "+sql, args)
	return b
}

// SetUsmallint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUsmallint(usmallint uint16) *GomodelUpdateBuilder {
	b.set("`usmallint`", "`usmallint`=?", usmallint)
	return b
}

func (b *GomodelUpdateBuilder) IncrUsmallint(val uint16) *GomodelUpdateBuilder {
	b.set("`usmallint`", "`usmallint`=`usmallint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUsmallint(val uint16) *GomodelUpdateBuilder {
	b.set("`usmallint`", "`usmallint`=`usmallint`-?", val)
	return b
}

// SetUsmallintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUsmallintZero() *GomodelUpdateBuilder {
	b.set("`usmallint`", "`usmallint`=?", 0)
	return b
}

// SetUsmallintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUsmallintNull() *GomodelUpdateBuilder {
	b.set("`usmallint`", "`usmallint`=NULL", nil)
	return b
}

// SetUsmallintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUsmallintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`usmallint`", "`usmallint` "+sql, args)
	return b
}

// SetUmediumint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUmediumint(umediumint uint32) *GomodelUpdateBuilder {
	b.set("`umediumint`", "`umediumint`=?", umediumint)
	return b
}

func (b *GomodelUpdateBuilder) IncrUmediumint(val uint32) *GomodelUpdateBuilder {
	b.set("`umediumint`", "`umediumint`=`umediumint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUmediumint(val uint32) *GomodelUpdateBuilder {
	b.set("`umediumint`", "`umediumint`=`umediumint`-?", val)
	return b
}

// SetUmediumintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUmediumintZero() *GomodelUpdateBuilder {
	b.set("`umediumint`", "`umediumint`=?", 0)
	return b
}

// SetUmediumintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUmediumintNull() *GomodelUpdateBuilder {
	b.set("`umediumint`", "`umediumint`=NULL", nil)
	return b
}

// SetUmediumintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUmediumintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`umediumint`", "`umediumint` "+sql, args)
	return b
}

// SetUint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUint(uint uint) *GomodelUpdateBuilder {
	b.set("`uint`", "`uint`=?", uint)
	return b
}

func (b *GomodelUpdateBuilder) IncrUint(val uint) *GomodelUpdateBuilder {
	b.set("`uint`", "`uint`=`uint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUint(val uint) *GomodelUpdateBuilder {
	b.set("`uint`", "`uint`=`uint`-?", val)
	return b
}

// SetUintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUintZero() *GomodelUpdateBuilder {
	b.set("`uint`", "`uint`=?", 0)
	return b
}

// SetUintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUintNull() *GomodelUpdateBuilder {
	b.set("`uint`", "`uint`=NULL", nil)
	return b
}

// SetUintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`uint`", "`uint` "+sql, args)
	return b
}

// SetUbigint 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUbigint(ubigint uint64) *GomodelUpdateBuilder {
	b.set("`ubigint`", "`ubigint`=?", ubigint)
	return b
}

func (b *GomodelUpdateBuilder) IncrUbigint(val uint64) *GomodelUpdateBuilder {
	b.set("`ubigint`", "`ubigint`=`ubigint`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUbigint(val uint64) *GomodelUpdateBuilder {
	b.set("`ubigint`", "`ubigint`=`ubigint`-?", val)
	return b
}

// SetUbigintZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUbigintZero() *GomodelUpdateBuilder {
	b.set("`ubigint`", "`ubigint`=?", 0)
	return b
}

// SetUbigintNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUbigintNull() *GomodelUpdateBuilder {
	b.set("`ubigint`", "`ubigint`=NULL", nil)
	return b
}

// SetUbigintRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUbigintRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`ubigint`", "`ubigint` "+sql, args)
	return b
}

// SetUfloat 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUfloat(ufloat float32) *GomodelUpdateBuilder {
	b.set("`ufloat`", "`ufloat`=?", ufloat)
	return b
}

func (b *GomodelUpdateBuilder) IncrUfloat(val float32) *GomodelUpdateBuilder {
	b.set("`ufloat`", "`ufloat`=`ufloat`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUfloat(val float32) *GomodelUpdateBuilder {
	b.set("`ufloat`", "`ufloat`=`ufloat`-?", val)
	return b
}

// SetUfloatZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUfloatZero() *GomodelUpdateBuilder {
	b.set("`ufloat`", "`ufloat`=?", 0)
	return b
}

// SetUfloatNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUfloatNull() *GomodelUpdateBuilder {
	b.set("`ufloat`", "`ufloat`=NULL", nil)
	return b
}

// SetUfloatRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUfloatRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`ufloat`", "`ufloat` "+sql, args)
	return b
}

// SetUdouble 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUdouble(udouble float64) *GomodelUpdateBuilder {
	b.set("`udouble`", "`udouble`=?", udouble)
	return b
}

func (b *GomodelUpdateBuilder) IncrUdouble(val float64) *GomodelUpdateBuilder {
	b.set("`udouble`", "`udouble`=`udouble`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUdouble(val float64) *GomodelUpdateBuilder {
	b.set("`udouble`", "`udouble`=`udouble`-?", val)
	return b
}

// SetUdoubleZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUdoubleZero() *GomodelUpdateBuilder {
	b.set("`udouble`", "`udouble`=?", 0)
	return b
}

// SetUdoubleNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUdoubleNull() *GomodelUpdateBuilder {
	b.set("`udouble`", "`udouble`=NULL", nil)
	return b
}

// SetUdoubleRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUdoubleRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`udouble`", "`udouble` "+sql, args)
	return b
}

// SetUdecimal 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetUdecimal(udecimal float64) *GomodelUpdateBuilder {
	b.set("`udecimal`", "`udecimal`=?", udecimal)
	return b
}

func (b *GomodelUpdateBuilder) IncrUdecimal(val float64) *GomodelUpdateBuilder {
	b.set("`udecimal`", "`udecimal`=`udecimal`+?", val)
	return b
}

func (b *GomodelUpdateBuilder) DecrUdecimal(val float64) *GomodelUpdateBuilder {
	b.set("`udecimal`", "`udecimal`=`udecimal`-?", val)
	return b
}

// SetUdecimalZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetUdecimalZero() *GomodelUpdateBuilder {
	b.set("`udecimal`", "`udecimal`=?", 0)
	return b
}

// SetUdecimalNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetUdecimalNull() *GomodelUpdateBuilder {
	b.set("`udecimal`", "`udecimal`=NULL", nil)
	return b
}

// SetUdecimalRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetUdecimalRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`udecimal`", "`udecimal` "+sql, args)
	return b
}

// SetDate 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetDate(date time.Time) *GomodelUpdateBuilder {
	b.set("`date`", "`date`=?", date)
	return b
}

// SetDateZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetDateZero() *GomodelUpdateBuilder {
	b.set("`date`", "`date`=?", time.Time{})
	return b
}

// SetDateNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetDateNull() *GomodelUpdateBuilder {
	b.set("`date`", "`date`=NULL", nil)
	return b
}

// SetDateRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetDateRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`date`", "`date` "+sql, args)
	return b
}

// SetDatetime 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetDatetime(datetime time.Time) *GomodelUpdateBuilder {
	b.set("`datetime`", "`datetime`=?", datetime)
	return b
}

// SetDatetimeZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetDatetimeZero() *GomodelUpdateBuilder {
	b.set("`datetime`", "`datetime`=?", time.Time{})
	return b
}

// SetDatetimeNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetDatetimeNull() *GomodelUpdateBuilder {
	b.set("`datetime`", "`datetime`=NULL", nil)
	return b
}

// SetDatetimeRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetDatetimeRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`datetime`", "`datetime` "+sql, args)
	return b
}

// SetTimestamp 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetTimestamp(timestamp time.Time) *GomodelUpdateBuilder {
	b.set("`timestamp`", "`timestamp`=?", timestamp)
	return b
}

// SetTimestampZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTimestampZero() *GomodelUpdateBuilder {
	b.set("`timestamp`", "`timestamp`=?", time.Time{})
	return b
}

// SetTimestampRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTimestampRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`timestamp`", "`timestamp` "+sql, args)
	return b
}

// SetTime 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetTime(time time.Time) *GomodelUpdateBuilder {
	b.set("`time`", "`time`=?", time)
	return b
}

// SetTimeZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTimeZero() *GomodelUpdateBuilder {
	b.set("`time`", "`time`=?", time.Time{})
	return b
}

// SetTimeNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetTimeNull() *GomodelUpdateBuilder {
	b.set("`time`", "`time`=NULL", nil)
	return b
}

// SetTimeRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTimeRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`time`", "`time` "+sql, args)
	return b
}

// SetYear 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetYear(year time.Time) *GomodelUpdateBuilder {
	b.set("`year`", "`year`=?", year)
	return b
}

// SetYearZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetYearZero() *GomodelUpdateBuilder {
	b.set("`year`", "`year`=?", time.Time{})
	return b
}

// SetYearNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetYearNull() *GomodelUpdateBuilder {
	b.set("`year`", "`year`=NULL", nil)
	return b
}

// SetYearRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetYearRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`year`", "`year` "+sql, args)
	return b
}

// SetChar 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetChar(char string) *GomodelUpdateBuilder {
	b.set("`char`", "`char`=?", char)
	return b
}

// SetCharZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetCharZero() *GomodelUpdateBuilder {
	b.set("`char`", "`char`=?", "")
	return b
}

// SetCharNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetCharNull() *GomodelUpdateBuilder {
	b.set("`char`", "`char`=NULL", nil)
	return b
}

// SetCharRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetCharRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`char`", "`char` "+sql, args)
	return b
}

// SetVarchar 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetVarchar(varchar string) *GomodelUpdateBuilder {
	b.set("`varchar`", "`varchar`=?", varchar)
	return b
}

// SetVarcharZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetVarcharZero() *GomodelUpdateBuilder {
	b.set("`varchar`", "`varchar`=?", "")
	return b
}

// SetVarcharNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetVarcharNull() *GomodelUpdateBuilder {
	b.set("`varchar`", "`varchar`=NULL", nil)
	return b
}

// SetVarcharRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetVarcharRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`varchar`", "`varchar` "+sql, args)
	return b
}

// SetBinary 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetBinary(binary string) *GomodelUpdateBuilder {
	b.set("`binary`", "`binary`=?", binary)
	return b
}

// SetBinaryZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetBinaryZero() *GomodelUpdateBuilder {
	b.set("`binary`", "`binary`=?", "")
	return b
}

// SetBinaryNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetBinaryNull() *GomodelUpdateBuilder {
	b.set("`binary`", "`binary`=NULL", nil)
	return b
}

// SetBinaryRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetBinaryRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`binary`", "`binary` "+sql, args)
	return b
}

// SetVarbinary 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetVarbinary(varbinary string) *GomodelUpdateBuilder {
	b.set("`varbinary`", "`varbinary`=?", varbinary)
	return b
}

// SetVarbinaryZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetVarbinaryZero() *GomodelUpdateBuilder {
	b.set("`varbinary`", "`varbinary`=?", "")
	return b
}

// SetVarbinaryNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetVarbinaryNull() *GomodelUpdateBuilder {
	b.set("`varbinary`", "`varbinary`=NULL", nil)
	return b
}

// SetVarbinaryRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetVarbinaryRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`varbinary`", "`varbinary` "+sql, args)
	return b
}

// SetTinyblob 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetTinyblob(tinyblob string) *GomodelUpdateBuilder {
	b.set("`tinyblob`", "`tinyblob`=?", tinyblob)
	return b
}

// SetTinyblobZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTinyblobZero() *GomodelUpdateBuilder {
	b.set("`tinyblob`", "`tinyblob`=?", "")
	return b
}

// SetTinyblobNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetTinyblobNull() *GomodelUpdateBuilder {
	b.set("`tinyblob`", "`tinyblob`=NULL", nil)
	return b
}

// SetTinyblobRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTinyblobRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`tinyblob`", "`tinyblob` "+sql, args)
	return b
}

// SetTinytext 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetTinytext(tinytext string) *GomodelUpdateBuilder {
	b.set("`tinytext`", "`tinytext`=?", tinytext)
	return b
}

// SetTinytextZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTinytextZero() *GomodelUpdateBuilder {
	b.set("`tinytext`", "`tinytext`=?", "")
	return b
}

// SetTinytextNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetTinytextNull() *GomodelUpdateBuilder {
	b.set("`tinytext`", "`tinytext`=NULL", nil)
	return b
}

// SetTinytextRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTinytextRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`tinytext`", "`tinytext` "+sql, args)
	return b
}

// SetBlob 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetBlob(blob string) *GomodelUpdateBuilder {
	b.set("`blob`", "`blob`=?", blob)
	return b
}

// SetBlobZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetBlobZero() *GomodelUpdateBuilder {
	b.set("`blob`", "`blob`=?", "")
	return b
}

// SetBlobNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetBlobNull() *GomodelUpdateBuilder {
	b.set("`blob`", "`blob`=NULL", nil)
	return b
}

// SetBlobRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetBlobRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`blob`", "`blob` "+sql, args)
	return b
}

// SetText 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetText(text string) *GomodelUpdateBuilder {
	b.set("`text`", "`text`=?", text)
	return b
}

// SetTextZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTextZero() *GomodelUpdateBuilder {
	b.set("`text`", "`text`=?", "")
	return b
}

// SetTextNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetTextNull() *GomodelUpdateBuilder {
	b.set("`text`", "`text`=NULL", nil)
	return b
}

// SetTextRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTextRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`text`", "`text` "+sql, args)
	return b
}

// SetMediumblob 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetMediumblob(mediumblob string) *GomodelUpdateBuilder {
	b.set("`mediumblob`", "`mediumblob`=?", mediumblob)
	return b
}

// SetMediumblobZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetMediumblobZero() *GomodelUpdateBuilder {
	b.set("`mediumblob`", "`mediumblob`=?", "")
	return b
}

// SetMediumblobNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetMediumblobNull() *GomodelUpdateBuilder {
	b.set("`mediumblob`", "`mediumblob`=NULL", nil)
	return b
}

// SetMediumblobRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetMediumblobRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`mediumblob`", "`mediumblob` "+sql, args)
	return b
}

// SetMediumtext 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetMediumtext(mediumtext string) *GomodelUpdateBuilder {
	b.set("`mediumtext`", "`mediumtext`=?", mediumtext)
	return b
}

// SetMediumtextZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetMediumtextZero() *GomodelUpdateBuilder {
	b.set("`mediumtext`", "`mediumtext`=?", "")
	return b
}

// SetMediumtextNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetMediumtextNull() *GomodelUpdateBuilder {
	b.set("`mediumtext`", "`mediumtext`=NULL", nil)
	return b
}

// SetMediumtextRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetMediumtextRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`mediumtext`", "`mediumtext` "+sql, args)
	return b
}

// SetLongblob 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetLongblob(longblob string) *GomodelUpdateBuilder {
	b.set("`longblob`", "`longblob`=?", longblob)
	return b
}

// SetLongblobZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetLongblobZero() *GomodelUpdateBuilder {
	b.set("`longblob`", "`longblob`=?", "")
	return b
}

// SetLongblobNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetLongblobNull() *GomodelUpdateBuilder {
	b.set("`longblob`", "`longblob`=NULL", nil)
	return b
}

// SetLongblobRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetLongblobRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`longblob`", "`longblob` "+sql, args)
	return b
}

// SetLongtext 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetLongtext(longtext string) *GomodelUpdateBuilder {
	b.set("`longtext`", "`longtext`=?", longtext)
	return b
}

// SetLongtextZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetLongtextZero() *GomodelUpdateBuilder {
	b.set("`longtext`", "`longtext`=?", "")
	return b
}

// SetLongtextNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetLongtextNull() *GomodelUpdateBuilder {
	b.set("`longtext`", "`longtext`=NULL", nil)
	return b
}

// SetLongtextRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetLongtextRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`longtext`", "`longtext` "+sql, args)
	return b
}

// SetEnum 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetEnum(enum string) *GomodelUpdateBuilder {
	b.set("`enum`", "`enum`=?", enum)
	return b
}

// SetEnumZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetEnumZero() *GomodelUpdateBuilder {
	b.set("`enum`", "`enum`=?", "")
	return b
}

// SetEnumNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetEnumNull() *GomodelUpdateBuilder {
	b.set("`enum`", "`enum`=NULL", nil)
	return b
}

// SetEnumRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetEnumRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`enum`", "`enum` "+sql, args)
	return b
}

// SetSet 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetSet(set string) *GomodelUpdateBuilder {
	b.set("`set`", "`set`=?", set)
	return b
}

// SetSetZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetSetZero() *GomodelUpdateBuilder {
	b.set("`set`", "`set`=?", "")
	return b
}

// SetSetNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetSetNull() *GomodelUpdateBuilder {
	b.set("`set`", "`set`=NULL", nil)
	return b
}

// SetSetRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetSetRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`set`", "`set` "+sql, args)
	return b
}

// SetJson 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetJson(json string) *GomodelUpdateBuilder {
	b.set("`json`", "`json`=?", json)
	return b
}

// SetJsonZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetJsonZero() *GomodelUpdateBuilder {
	b.set("`json`", "`json`=?", "")
	return b
}

// SetJsonNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetJsonNull() *GomodelUpdateBuilder {
	b.set("`json`", "`json`=NULL", nil)
	return b
}

// SetJsonRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetJsonRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`json`", "`json` "+sql, args)
	return b
}

// SetTinybool 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetTinybool(tinybool bool) *GomodelUpdateBuilder {
	b.set("`tinybool`", "`tinybool`=?", tinybool)
	return b
}

// SetTinyboolZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetTinyboolZero() *GomodelUpdateBuilder {
	b.set("`tinybool`", "`tinybool`=?", false)
	return b
}

// SetTinyboolNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetTinyboolNull() *GomodelUpdateBuilder {
	b.set("`tinybool`", "`tinybool`=NULL", nil)
	return b
}

// SetTinyboolRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetTinyboolRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`tinybool`", "`tinybool` "+sql, args)
	return b
}

// SetBool 将字段更新为指定值
func (b *GomodelUpdateBuilder) SetBool(bool bool) *GomodelUpdateBuilder {
	b.set("`bool`", "`bool`=?", bool)
	return b
}

// SetBoolZero 将字段更新为零值
func (b *GomodelUpdateBuilder) SetBoolZero() *GomodelUpdateBuilder {
	b.set("`bool`", "`bool`=?", false)
	return b
}

// SetBoolNull 将字段更新为 Null
func (b *GomodelUpdateBuilder) SetBoolNull() *GomodelUpdateBuilder {
	b.set("`bool`", "`bool`=NULL", nil)
	return b
}

// SetBoolRaw 自定义更新语句
func (b *GomodelUpdateBuilder) SetBoolRaw(sql string, args ...interface{}) *GomodelUpdateBuilder {
	b.set("`bool`", "`bool` "+sql, args)
	return b
}

func (b *GomodelUpdateBuilder) Where(fn func(b *GomodelWhereBuilder)) *GomodelUpdateBuilder {
	if b.where == nil {
		b.where = newGomodelWhereBuilder()
	}
	fn(b.where)
	return b
}

func (b *GomodelUpdateBuilder) SQL() (string, []interface{}) {
	var sb strings.Builder
	args := make([]interface{}, 0, len(b.changes))
	sb.WriteString("UPDATE ")
	sb.WriteString("`")
	sb.WriteString(b.table)
	sb.WriteString("` SET ")

	comma := false
	for _, s := range b.changes {
		if comma {
			sb.WriteString(",")
		}
		sb.WriteString(s.sql)
		comma = true

		if val, ok := s.arg.([]interface{}); ok {
			args = append(args, val...)
		} else if s.arg != nil {
			args = append(args, s.arg)
		}
	}

	if b.where != nil {
		whereSQL, whereArgs := b.where.sql()
		sb.WriteString(" ")
		sb.WriteString(whereSQL)
		args = append(args, whereArgs...)
	}
	return sb.String(), args
}

func (b *GomodelUpdateBuilder) Exec(ctx context.Context) (int64, error) {
	sql, args := b.SQL()
	info := &queryInfo{
		ctx:   ctx,
		table: b.table,
		op:    OpUpdate,
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

func (b *GomodelUpdateBuilder) set(field string, sql string, arg interface{}) {
	s, exists := b.changes[field]
	if exists {
		s.sql = sql
		s.arg = arg
		return
	}

	b.changes[field] = &set{
		sql: sql,
		arg: arg,
	}
}
