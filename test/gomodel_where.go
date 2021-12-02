package test

import (
	"strings"
	"time"
)

var _ = time.Second

type GomodelWhereBuilder struct {
	sb   strings.Builder
	args []interface{}
}

func newGomodelWhereBuilder() *GomodelWhereBuilder {
	b := &GomodelWhereBuilder{}
	b.sb.WriteString("WHERE ")
	return b
}

func (b *GomodelWhereBuilder) And() *GomodelWhereBuilder {
	b.sb.WriteString(" AND ")
	return b
}

func (b *GomodelWhereBuilder) Or() *GomodelWhereBuilder {
	b.sb.WriteString(" OR ")
	return b
}

func (b *GomodelWhereBuilder) ConditionGroup(fn func(b *GomodelWhereBuilder)) *GomodelWhereBuilder {
	b.sb.WriteString("(")
	fn(b)
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyintEQ(tinyint int8) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`=?")
	b.args = append(b.args, tinyint)
	return b
}

func (b *GomodelWhereBuilder) TinyintNEQ(tinyint int8) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`<>?")
	b.args = append(b.args, tinyint)
	return b
}

func (b *GomodelWhereBuilder) TinyintIn(tinyintList ...int8) *GomodelWhereBuilder {
	if len(tinyintList) == 0 {
		return b
	}

	b.sb.WriteString("`tinyint` IN (?")
	b.args = append(b.args, tinyintList[0])
	for i, size := 1, len(tinyintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinyintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyintNotIn(tinyintList ...int8) *GomodelWhereBuilder {
	if len(tinyintList) == 0 {
		return b
	}

	b.sb.WriteString("`tinyint` NOT IN (?")
	b.args = append(b.args, tinyintList[0])
	for i, size := 1, len(tinyintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinyintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyintGT(tinyint int8) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`>?")
	b.args = append(b.args, tinyint)
	return b
}

func (b *GomodelWhereBuilder) TinyintGTE(tinyint int8) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`>=?")
	b.args = append(b.args, tinyint)
	return b
}

func (b *GomodelWhereBuilder) TinyintLT(tinyint int8) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`<?")
	b.args = append(b.args, tinyint)
	return b
}

func (b *GomodelWhereBuilder) TinyintLTE(tinyint int8) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`<=?")
	b.args = append(b.args, tinyint)
	return b
}

func (b *GomodelWhereBuilder) TinyintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) SmallintEQ(smallint int16) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`=?")
	b.args = append(b.args, smallint)
	return b
}

func (b *GomodelWhereBuilder) SmallintNEQ(smallint int16) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`<>?")
	b.args = append(b.args, smallint)
	return b
}

func (b *GomodelWhereBuilder) SmallintIn(smallintList ...int16) *GomodelWhereBuilder {
	if len(smallintList) == 0 {
		return b
	}

	b.sb.WriteString("`smallint` IN (?")
	b.args = append(b.args, smallintList[0])
	for i, size := 1, len(smallintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, smallintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) SmallintNotIn(smallintList ...int16) *GomodelWhereBuilder {
	if len(smallintList) == 0 {
		return b
	}

	b.sb.WriteString("`smallint` NOT IN (?")
	b.args = append(b.args, smallintList[0])
	for i, size := 1, len(smallintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, smallintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) SmallintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`smallint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) SmallintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`smallint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) SmallintGT(smallint int16) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`>?")
	b.args = append(b.args, smallint)
	return b
}

func (b *GomodelWhereBuilder) SmallintGTE(smallint int16) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`>=?")
	b.args = append(b.args, smallint)
	return b
}

func (b *GomodelWhereBuilder) SmallintLT(smallint int16) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`<?")
	b.args = append(b.args, smallint)
	return b
}

func (b *GomodelWhereBuilder) SmallintLTE(smallint int16) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`<=?")
	b.args = append(b.args, smallint)
	return b
}

func (b *GomodelWhereBuilder) SmallintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`smallint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) MediumintEQ(mediumint int32) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`=?")
	b.args = append(b.args, mediumint)
	return b
}

func (b *GomodelWhereBuilder) MediumintNEQ(mediumint int32) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`<>?")
	b.args = append(b.args, mediumint)
	return b
}

func (b *GomodelWhereBuilder) MediumintIn(mediumintList ...int32) *GomodelWhereBuilder {
	if len(mediumintList) == 0 {
		return b
	}

	b.sb.WriteString("`mediumint` IN (?")
	b.args = append(b.args, mediumintList[0])
	for i, size := 1, len(mediumintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, mediumintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) MediumintNotIn(mediumintList ...int32) *GomodelWhereBuilder {
	if len(mediumintList) == 0 {
		return b
	}

	b.sb.WriteString("`mediumint` NOT IN (?")
	b.args = append(b.args, mediumintList[0])
	for i, size := 1, len(mediumintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, mediumintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) MediumintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) MediumintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) MediumintGT(mediumint int32) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`>?")
	b.args = append(b.args, mediumint)
	return b
}

func (b *GomodelWhereBuilder) MediumintGTE(mediumint int32) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`>=?")
	b.args = append(b.args, mediumint)
	return b
}

func (b *GomodelWhereBuilder) MediumintLT(mediumint int32) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`<?")
	b.args = append(b.args, mediumint)
	return b
}

func (b *GomodelWhereBuilder) MediumintLTE(mediumint int32) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`<=?")
	b.args = append(b.args, mediumint)
	return b
}

func (b *GomodelWhereBuilder) MediumintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) IntEQ(int int) *GomodelWhereBuilder {
	b.sb.WriteString("`int`=?")
	b.args = append(b.args, int)
	return b
}

func (b *GomodelWhereBuilder) IntNEQ(int int) *GomodelWhereBuilder {
	b.sb.WriteString("`int`<>?")
	b.args = append(b.args, int)
	return b
}

func (b *GomodelWhereBuilder) IntIn(intList ...int) *GomodelWhereBuilder {
	if len(intList) == 0 {
		return b
	}

	b.sb.WriteString("`int` IN (?")
	b.args = append(b.args, intList[0])
	for i, size := 1, len(intList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, intList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) IntNotIn(intList ...int) *GomodelWhereBuilder {
	if len(intList) == 0 {
		return b
	}

	b.sb.WriteString("`int` NOT IN (?")
	b.args = append(b.args, intList[0])
	for i, size := 1, len(intList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, intList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) IntIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`int` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) IntIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`int` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) IntGT(int int) *GomodelWhereBuilder {
	b.sb.WriteString("`int`>?")
	b.args = append(b.args, int)
	return b
}

func (b *GomodelWhereBuilder) IntGTE(int int) *GomodelWhereBuilder {
	b.sb.WriteString("`int`>=?")
	b.args = append(b.args, int)
	return b
}

func (b *GomodelWhereBuilder) IntLT(int int) *GomodelWhereBuilder {
	b.sb.WriteString("`int`<?")
	b.args = append(b.args, int)
	return b
}

func (b *GomodelWhereBuilder) IntLTE(int int) *GomodelWhereBuilder {
	b.sb.WriteString("`int`<=?")
	b.args = append(b.args, int)
	return b
}

func (b *GomodelWhereBuilder) IntRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`int`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) BigintEQ(bigint int64) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`=?")
	b.args = append(b.args, bigint)
	return b
}

func (b *GomodelWhereBuilder) BigintNEQ(bigint int64) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`<>?")
	b.args = append(b.args, bigint)
	return b
}

func (b *GomodelWhereBuilder) BigintIn(bigintList ...int64) *GomodelWhereBuilder {
	if len(bigintList) == 0 {
		return b
	}

	b.sb.WriteString("`bigint` IN (?")
	b.args = append(b.args, bigintList[0])
	for i, size := 1, len(bigintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, bigintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BigintNotIn(bigintList ...int64) *GomodelWhereBuilder {
	if len(bigintList) == 0 {
		return b
	}

	b.sb.WriteString("`bigint` NOT IN (?")
	b.args = append(b.args, bigintList[0])
	for i, size := 1, len(bigintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, bigintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BigintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`bigint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) BigintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`bigint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) BigintGT(bigint int64) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`>?")
	b.args = append(b.args, bigint)
	return b
}

func (b *GomodelWhereBuilder) BigintGTE(bigint int64) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`>=?")
	b.args = append(b.args, bigint)
	return b
}

func (b *GomodelWhereBuilder) BigintLT(bigint int64) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`<?")
	b.args = append(b.args, bigint)
	return b
}

func (b *GomodelWhereBuilder) BigintLTE(bigint int64) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`<=?")
	b.args = append(b.args, bigint)
	return b
}

func (b *GomodelWhereBuilder) BigintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`bigint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) FloatEQ(float float32) *GomodelWhereBuilder {
	b.sb.WriteString("`float`=?")
	b.args = append(b.args, float)
	return b
}

func (b *GomodelWhereBuilder) FloatNEQ(float float32) *GomodelWhereBuilder {
	b.sb.WriteString("`float`<>?")
	b.args = append(b.args, float)
	return b
}

func (b *GomodelWhereBuilder) FloatIn(floatList ...float32) *GomodelWhereBuilder {
	if len(floatList) == 0 {
		return b
	}

	b.sb.WriteString("`float` IN (?")
	b.args = append(b.args, floatList[0])
	for i, size := 1, len(floatList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, floatList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) FloatNotIn(floatList ...float32) *GomodelWhereBuilder {
	if len(floatList) == 0 {
		return b
	}

	b.sb.WriteString("`float` NOT IN (?")
	b.args = append(b.args, floatList[0])
	for i, size := 1, len(floatList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, floatList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) FloatIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`float` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) FloatIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`float` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) FloatGT(float float32) *GomodelWhereBuilder {
	b.sb.WriteString("`float`>?")
	b.args = append(b.args, float)
	return b
}

func (b *GomodelWhereBuilder) FloatGTE(float float32) *GomodelWhereBuilder {
	b.sb.WriteString("`float`>=?")
	b.args = append(b.args, float)
	return b
}

func (b *GomodelWhereBuilder) FloatLT(float float32) *GomodelWhereBuilder {
	b.sb.WriteString("`float`<?")
	b.args = append(b.args, float)
	return b
}

func (b *GomodelWhereBuilder) FloatLTE(float float32) *GomodelWhereBuilder {
	b.sb.WriteString("`float`<=?")
	b.args = append(b.args, float)
	return b
}

func (b *GomodelWhereBuilder) FloatRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`float`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) DoubleEQ(double float64) *GomodelWhereBuilder {
	b.sb.WriteString("`double`=?")
	b.args = append(b.args, double)
	return b
}

func (b *GomodelWhereBuilder) DoubleNEQ(double float64) *GomodelWhereBuilder {
	b.sb.WriteString("`double`<>?")
	b.args = append(b.args, double)
	return b
}

func (b *GomodelWhereBuilder) DoubleIn(doubleList ...float64) *GomodelWhereBuilder {
	if len(doubleList) == 0 {
		return b
	}

	b.sb.WriteString("`double` IN (?")
	b.args = append(b.args, doubleList[0])
	for i, size := 1, len(doubleList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, doubleList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DoubleNotIn(doubleList ...float64) *GomodelWhereBuilder {
	if len(doubleList) == 0 {
		return b
	}

	b.sb.WriteString("`double` NOT IN (?")
	b.args = append(b.args, doubleList[0])
	for i, size := 1, len(doubleList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, doubleList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DoubleIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`double` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) DoubleIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`double` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) DoubleGT(double float64) *GomodelWhereBuilder {
	b.sb.WriteString("`double`>?")
	b.args = append(b.args, double)
	return b
}

func (b *GomodelWhereBuilder) DoubleGTE(double float64) *GomodelWhereBuilder {
	b.sb.WriteString("`double`>=?")
	b.args = append(b.args, double)
	return b
}

func (b *GomodelWhereBuilder) DoubleLT(double float64) *GomodelWhereBuilder {
	b.sb.WriteString("`double`<?")
	b.args = append(b.args, double)
	return b
}

func (b *GomodelWhereBuilder) DoubleLTE(double float64) *GomodelWhereBuilder {
	b.sb.WriteString("`double`<=?")
	b.args = append(b.args, double)
	return b
}

func (b *GomodelWhereBuilder) DoubleRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`double`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) DecimalEQ(decimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`=?")
	b.args = append(b.args, decimal)
	return b
}

func (b *GomodelWhereBuilder) DecimalNEQ(decimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`<>?")
	b.args = append(b.args, decimal)
	return b
}

func (b *GomodelWhereBuilder) DecimalIn(decimalList ...float64) *GomodelWhereBuilder {
	if len(decimalList) == 0 {
		return b
	}

	b.sb.WriteString("`decimal` IN (?")
	b.args = append(b.args, decimalList[0])
	for i, size := 1, len(decimalList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, decimalList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DecimalNotIn(decimalList ...float64) *GomodelWhereBuilder {
	if len(decimalList) == 0 {
		return b
	}

	b.sb.WriteString("`decimal` NOT IN (?")
	b.args = append(b.args, decimalList[0])
	for i, size := 1, len(decimalList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, decimalList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DecimalIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`decimal` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) DecimalIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`decimal` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) DecimalGT(decimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`>?")
	b.args = append(b.args, decimal)
	return b
}

func (b *GomodelWhereBuilder) DecimalGTE(decimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`>=?")
	b.args = append(b.args, decimal)
	return b
}

func (b *GomodelWhereBuilder) DecimalLT(decimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`<?")
	b.args = append(b.args, decimal)
	return b
}

func (b *GomodelWhereBuilder) DecimalLTE(decimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`<=?")
	b.args = append(b.args, decimal)
	return b
}

func (b *GomodelWhereBuilder) DecimalRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`decimal`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UtinyintEQ(utinyint uint8) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`=?")
	b.args = append(b.args, utinyint)
	return b
}

func (b *GomodelWhereBuilder) UtinyintNEQ(utinyint uint8) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`<>?")
	b.args = append(b.args, utinyint)
	return b
}

func (b *GomodelWhereBuilder) UtinyintIn(utinyintList ...uint8) *GomodelWhereBuilder {
	if len(utinyintList) == 0 {
		return b
	}

	b.sb.WriteString("`utinyint` IN (?")
	b.args = append(b.args, utinyintList[0])
	for i, size := 1, len(utinyintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, utinyintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UtinyintNotIn(utinyintList ...uint8) *GomodelWhereBuilder {
	if len(utinyintList) == 0 {
		return b
	}

	b.sb.WriteString("`utinyint` NOT IN (?")
	b.args = append(b.args, utinyintList[0])
	for i, size := 1, len(utinyintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, utinyintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UtinyintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UtinyintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UtinyintGT(utinyint uint8) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`>?")
	b.args = append(b.args, utinyint)
	return b
}

func (b *GomodelWhereBuilder) UtinyintGTE(utinyint uint8) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`>=?")
	b.args = append(b.args, utinyint)
	return b
}

func (b *GomodelWhereBuilder) UtinyintLT(utinyint uint8) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`<?")
	b.args = append(b.args, utinyint)
	return b
}

func (b *GomodelWhereBuilder) UtinyintLTE(utinyint uint8) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`<=?")
	b.args = append(b.args, utinyint)
	return b
}

func (b *GomodelWhereBuilder) UtinyintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`utinyint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UsmallintEQ(usmallint uint16) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`=?")
	b.args = append(b.args, usmallint)
	return b
}

func (b *GomodelWhereBuilder) UsmallintNEQ(usmallint uint16) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`<>?")
	b.args = append(b.args, usmallint)
	return b
}

func (b *GomodelWhereBuilder) UsmallintIn(usmallintList ...uint16) *GomodelWhereBuilder {
	if len(usmallintList) == 0 {
		return b
	}

	b.sb.WriteString("`usmallint` IN (?")
	b.args = append(b.args, usmallintList[0])
	for i, size := 1, len(usmallintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, usmallintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UsmallintNotIn(usmallintList ...uint16) *GomodelWhereBuilder {
	if len(usmallintList) == 0 {
		return b
	}

	b.sb.WriteString("`usmallint` NOT IN (?")
	b.args = append(b.args, usmallintList[0])
	for i, size := 1, len(usmallintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, usmallintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UsmallintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UsmallintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UsmallintGT(usmallint uint16) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`>?")
	b.args = append(b.args, usmallint)
	return b
}

func (b *GomodelWhereBuilder) UsmallintGTE(usmallint uint16) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`>=?")
	b.args = append(b.args, usmallint)
	return b
}

func (b *GomodelWhereBuilder) UsmallintLT(usmallint uint16) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`<?")
	b.args = append(b.args, usmallint)
	return b
}

func (b *GomodelWhereBuilder) UsmallintLTE(usmallint uint16) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`<=?")
	b.args = append(b.args, usmallint)
	return b
}

func (b *GomodelWhereBuilder) UsmallintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`usmallint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UmediumintEQ(umediumint uint32) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`=?")
	b.args = append(b.args, umediumint)
	return b
}

func (b *GomodelWhereBuilder) UmediumintNEQ(umediumint uint32) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`<>?")
	b.args = append(b.args, umediumint)
	return b
}

func (b *GomodelWhereBuilder) UmediumintIn(umediumintList ...uint32) *GomodelWhereBuilder {
	if len(umediumintList) == 0 {
		return b
	}

	b.sb.WriteString("`umediumint` IN (?")
	b.args = append(b.args, umediumintList[0])
	for i, size := 1, len(umediumintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, umediumintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UmediumintNotIn(umediumintList ...uint32) *GomodelWhereBuilder {
	if len(umediumintList) == 0 {
		return b
	}

	b.sb.WriteString("`umediumint` NOT IN (?")
	b.args = append(b.args, umediumintList[0])
	for i, size := 1, len(umediumintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, umediumintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UmediumintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UmediumintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UmediumintGT(umediumint uint32) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`>?")
	b.args = append(b.args, umediumint)
	return b
}

func (b *GomodelWhereBuilder) UmediumintGTE(umediumint uint32) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`>=?")
	b.args = append(b.args, umediumint)
	return b
}

func (b *GomodelWhereBuilder) UmediumintLT(umediumint uint32) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`<?")
	b.args = append(b.args, umediumint)
	return b
}

func (b *GomodelWhereBuilder) UmediumintLTE(umediumint uint32) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`<=?")
	b.args = append(b.args, umediumint)
	return b
}

func (b *GomodelWhereBuilder) UmediumintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`umediumint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UintEQ(uint uint) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`=?")
	b.args = append(b.args, uint)
	return b
}

func (b *GomodelWhereBuilder) UintNEQ(uint uint) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`<>?")
	b.args = append(b.args, uint)
	return b
}

func (b *GomodelWhereBuilder) UintIn(uintList ...uint) *GomodelWhereBuilder {
	if len(uintList) == 0 {
		return b
	}

	b.sb.WriteString("`uint` IN (?")
	b.args = append(b.args, uintList[0])
	for i, size := 1, len(uintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, uintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UintNotIn(uintList ...uint) *GomodelWhereBuilder {
	if len(uintList) == 0 {
		return b
	}

	b.sb.WriteString("`uint` NOT IN (?")
	b.args = append(b.args, uintList[0])
	for i, size := 1, len(uintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, uintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`uint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`uint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UintGT(uint uint) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`>?")
	b.args = append(b.args, uint)
	return b
}

func (b *GomodelWhereBuilder) UintGTE(uint uint) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`>=?")
	b.args = append(b.args, uint)
	return b
}

func (b *GomodelWhereBuilder) UintLT(uint uint) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`<?")
	b.args = append(b.args, uint)
	return b
}

func (b *GomodelWhereBuilder) UintLTE(uint uint) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`<=?")
	b.args = append(b.args, uint)
	return b
}

func (b *GomodelWhereBuilder) UintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`uint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UbigintEQ(ubigint uint64) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`=?")
	b.args = append(b.args, ubigint)
	return b
}

func (b *GomodelWhereBuilder) UbigintNEQ(ubigint uint64) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`<>?")
	b.args = append(b.args, ubigint)
	return b
}

func (b *GomodelWhereBuilder) UbigintIn(ubigintList ...uint64) *GomodelWhereBuilder {
	if len(ubigintList) == 0 {
		return b
	}

	b.sb.WriteString("`ubigint` IN (?")
	b.args = append(b.args, ubigintList[0])
	for i, size := 1, len(ubigintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, ubigintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UbigintNotIn(ubigintList ...uint64) *GomodelWhereBuilder {
	if len(ubigintList) == 0 {
		return b
	}

	b.sb.WriteString("`ubigint` NOT IN (?")
	b.args = append(b.args, ubigintList[0])
	for i, size := 1, len(ubigintList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, ubigintList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UbigintIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UbigintIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UbigintGT(ubigint uint64) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`>?")
	b.args = append(b.args, ubigint)
	return b
}

func (b *GomodelWhereBuilder) UbigintGTE(ubigint uint64) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`>=?")
	b.args = append(b.args, ubigint)
	return b
}

func (b *GomodelWhereBuilder) UbigintLT(ubigint uint64) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`<?")
	b.args = append(b.args, ubigint)
	return b
}

func (b *GomodelWhereBuilder) UbigintLTE(ubigint uint64) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`<=?")
	b.args = append(b.args, ubigint)
	return b
}

func (b *GomodelWhereBuilder) UbigintRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`ubigint`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UfloatEQ(ufloat float32) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`=?")
	b.args = append(b.args, ufloat)
	return b
}

func (b *GomodelWhereBuilder) UfloatNEQ(ufloat float32) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`<>?")
	b.args = append(b.args, ufloat)
	return b
}

func (b *GomodelWhereBuilder) UfloatIn(ufloatList ...float32) *GomodelWhereBuilder {
	if len(ufloatList) == 0 {
		return b
	}

	b.sb.WriteString("`ufloat` IN (?")
	b.args = append(b.args, ufloatList[0])
	for i, size := 1, len(ufloatList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, ufloatList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UfloatNotIn(ufloatList ...float32) *GomodelWhereBuilder {
	if len(ufloatList) == 0 {
		return b
	}

	b.sb.WriteString("`ufloat` NOT IN (?")
	b.args = append(b.args, ufloatList[0])
	for i, size := 1, len(ufloatList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, ufloatList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UfloatIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UfloatIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UfloatGT(ufloat float32) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`>?")
	b.args = append(b.args, ufloat)
	return b
}

func (b *GomodelWhereBuilder) UfloatGTE(ufloat float32) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`>=?")
	b.args = append(b.args, ufloat)
	return b
}

func (b *GomodelWhereBuilder) UfloatLT(ufloat float32) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`<?")
	b.args = append(b.args, ufloat)
	return b
}

func (b *GomodelWhereBuilder) UfloatLTE(ufloat float32) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`<=?")
	b.args = append(b.args, ufloat)
	return b
}

func (b *GomodelWhereBuilder) UfloatRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`ufloat`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UdoubleEQ(udouble float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`=?")
	b.args = append(b.args, udouble)
	return b
}

func (b *GomodelWhereBuilder) UdoubleNEQ(udouble float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`<>?")
	b.args = append(b.args, udouble)
	return b
}

func (b *GomodelWhereBuilder) UdoubleIn(udoubleList ...float64) *GomodelWhereBuilder {
	if len(udoubleList) == 0 {
		return b
	}

	b.sb.WriteString("`udouble` IN (?")
	b.args = append(b.args, udoubleList[0])
	for i, size := 1, len(udoubleList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, udoubleList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UdoubleNotIn(udoubleList ...float64) *GomodelWhereBuilder {
	if len(udoubleList) == 0 {
		return b
	}

	b.sb.WriteString("`udouble` NOT IN (?")
	b.args = append(b.args, udoubleList[0])
	for i, size := 1, len(udoubleList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, udoubleList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UdoubleIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`udouble` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UdoubleIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`udouble` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UdoubleGT(udouble float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`>?")
	b.args = append(b.args, udouble)
	return b
}

func (b *GomodelWhereBuilder) UdoubleGTE(udouble float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`>=?")
	b.args = append(b.args, udouble)
	return b
}

func (b *GomodelWhereBuilder) UdoubleLT(udouble float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`<?")
	b.args = append(b.args, udouble)
	return b
}

func (b *GomodelWhereBuilder) UdoubleLTE(udouble float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`<=?")
	b.args = append(b.args, udouble)
	return b
}

func (b *GomodelWhereBuilder) UdoubleRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`udouble`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) UdecimalEQ(udecimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`=?")
	b.args = append(b.args, udecimal)
	return b
}

func (b *GomodelWhereBuilder) UdecimalNEQ(udecimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`<>?")
	b.args = append(b.args, udecimal)
	return b
}

func (b *GomodelWhereBuilder) UdecimalIn(udecimalList ...float64) *GomodelWhereBuilder {
	if len(udecimalList) == 0 {
		return b
	}

	b.sb.WriteString("`udecimal` IN (?")
	b.args = append(b.args, udecimalList[0])
	for i, size := 1, len(udecimalList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, udecimalList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UdecimalNotIn(udecimalList ...float64) *GomodelWhereBuilder {
	if len(udecimalList) == 0 {
		return b
	}

	b.sb.WriteString("`udecimal` NOT IN (?")
	b.args = append(b.args, udecimalList[0])
	for i, size := 1, len(udecimalList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, udecimalList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) UdecimalIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) UdecimalIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) UdecimalGT(udecimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`>?")
	b.args = append(b.args, udecimal)
	return b
}

func (b *GomodelWhereBuilder) UdecimalGTE(udecimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`>=?")
	b.args = append(b.args, udecimal)
	return b
}

func (b *GomodelWhereBuilder) UdecimalLT(udecimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`<?")
	b.args = append(b.args, udecimal)
	return b
}

func (b *GomodelWhereBuilder) UdecimalLTE(udecimal float64) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`<=?")
	b.args = append(b.args, udecimal)
	return b
}

func (b *GomodelWhereBuilder) UdecimalRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`udecimal`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) DateEQ(date time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`date`=?")
	b.args = append(b.args, date)
	return b
}

func (b *GomodelWhereBuilder) DateNEQ(date time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`date`<>?")
	b.args = append(b.args, date)
	return b
}

func (b *GomodelWhereBuilder) DateIn(dateList ...time.Time) *GomodelWhereBuilder {
	if len(dateList) == 0 {
		return b
	}

	b.sb.WriteString("`date` IN (?")
	b.args = append(b.args, dateList[0])
	for i, size := 1, len(dateList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, dateList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DateNotIn(dateList ...time.Time) *GomodelWhereBuilder {
	if len(dateList) == 0 {
		return b
	}

	b.sb.WriteString("`date` NOT IN (?")
	b.args = append(b.args, dateList[0])
	for i, size := 1, len(dateList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, dateList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DateIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`date` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) DateIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`date` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) DateGT(date time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`date`>?")
	b.args = append(b.args, date)
	return b
}

func (b *GomodelWhereBuilder) DateGTE(date time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`date`>=?")
	b.args = append(b.args, date)
	return b
}

func (b *GomodelWhereBuilder) DateLT(date time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`date`<?")
	b.args = append(b.args, date)
	return b
}

func (b *GomodelWhereBuilder) DateLTE(date time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`date`<=?")
	b.args = append(b.args, date)
	return b
}

func (b *GomodelWhereBuilder) DateRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`date`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) DatetimeEQ(datetime time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`=?")
	b.args = append(b.args, datetime)
	return b
}

func (b *GomodelWhereBuilder) DatetimeNEQ(datetime time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`<>?")
	b.args = append(b.args, datetime)
	return b
}

func (b *GomodelWhereBuilder) DatetimeIn(datetimeList ...time.Time) *GomodelWhereBuilder {
	if len(datetimeList) == 0 {
		return b
	}

	b.sb.WriteString("`datetime` IN (?")
	b.args = append(b.args, datetimeList[0])
	for i, size := 1, len(datetimeList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, datetimeList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DatetimeNotIn(datetimeList ...time.Time) *GomodelWhereBuilder {
	if len(datetimeList) == 0 {
		return b
	}

	b.sb.WriteString("`datetime` NOT IN (?")
	b.args = append(b.args, datetimeList[0])
	for i, size := 1, len(datetimeList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, datetimeList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) DatetimeIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`datetime` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) DatetimeIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`datetime` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) DatetimeGT(datetime time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`>?")
	b.args = append(b.args, datetime)
	return b
}

func (b *GomodelWhereBuilder) DatetimeGTE(datetime time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`>=?")
	b.args = append(b.args, datetime)
	return b
}

func (b *GomodelWhereBuilder) DatetimeLT(datetime time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`<?")
	b.args = append(b.args, datetime)
	return b
}

func (b *GomodelWhereBuilder) DatetimeLTE(datetime time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`<=?")
	b.args = append(b.args, datetime)
	return b
}

func (b *GomodelWhereBuilder) DatetimeRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`datetime`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) TimestampEQ(timestamp time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`=?")
	b.args = append(b.args, timestamp)
	return b
}

func (b *GomodelWhereBuilder) TimestampNEQ(timestamp time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`<>?")
	b.args = append(b.args, timestamp)
	return b
}

func (b *GomodelWhereBuilder) TimestampIn(timestampList ...time.Time) *GomodelWhereBuilder {
	if len(timestampList) == 0 {
		return b
	}

	b.sb.WriteString("`timestamp` IN (?")
	b.args = append(b.args, timestampList[0])
	for i, size := 1, len(timestampList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, timestampList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TimestampNotIn(timestampList ...time.Time) *GomodelWhereBuilder {
	if len(timestampList) == 0 {
		return b
	}

	b.sb.WriteString("`timestamp` NOT IN (?")
	b.args = append(b.args, timestampList[0])
	for i, size := 1, len(timestampList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, timestampList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TimestampGT(timestamp time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`>?")
	b.args = append(b.args, timestamp)
	return b
}

func (b *GomodelWhereBuilder) TimestampGTE(timestamp time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`>=?")
	b.args = append(b.args, timestamp)
	return b
}

func (b *GomodelWhereBuilder) TimestampLT(timestamp time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`<?")
	b.args = append(b.args, timestamp)
	return b
}

func (b *GomodelWhereBuilder) TimestampLTE(timestamp time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`<=?")
	b.args = append(b.args, timestamp)
	return b
}

func (b *GomodelWhereBuilder) TimestampRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`timestamp`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) TimeEQ(time time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`time`=?")
	b.args = append(b.args, time)
	return b
}

func (b *GomodelWhereBuilder) TimeNEQ(time time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`time`<>?")
	b.args = append(b.args, time)
	return b
}

func (b *GomodelWhereBuilder) TimeIn(timeList ...time.Time) *GomodelWhereBuilder {
	if len(timeList) == 0 {
		return b
	}

	b.sb.WriteString("`time` IN (?")
	b.args = append(b.args, timeList[0])
	for i, size := 1, len(timeList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, timeList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TimeNotIn(timeList ...time.Time) *GomodelWhereBuilder {
	if len(timeList) == 0 {
		return b
	}

	b.sb.WriteString("`time` NOT IN (?")
	b.args = append(b.args, timeList[0])
	for i, size := 1, len(timeList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, timeList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TimeIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`time` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) TimeIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`time` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) TimeGT(time time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`time`>?")
	b.args = append(b.args, time)
	return b
}

func (b *GomodelWhereBuilder) TimeGTE(time time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`time`>=?")
	b.args = append(b.args, time)
	return b
}

func (b *GomodelWhereBuilder) TimeLT(time time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`time`<?")
	b.args = append(b.args, time)
	return b
}

func (b *GomodelWhereBuilder) TimeLTE(time time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`time`<=?")
	b.args = append(b.args, time)
	return b
}

func (b *GomodelWhereBuilder) TimeRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`time`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) YearEQ(year time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`year`=?")
	b.args = append(b.args, year)
	return b
}

func (b *GomodelWhereBuilder) YearNEQ(year time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`year`<>?")
	b.args = append(b.args, year)
	return b
}

func (b *GomodelWhereBuilder) YearIn(yearList ...time.Time) *GomodelWhereBuilder {
	if len(yearList) == 0 {
		return b
	}

	b.sb.WriteString("`year` IN (?")
	b.args = append(b.args, yearList[0])
	for i, size := 1, len(yearList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, yearList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) YearNotIn(yearList ...time.Time) *GomodelWhereBuilder {
	if len(yearList) == 0 {
		return b
	}

	b.sb.WriteString("`year` NOT IN (?")
	b.args = append(b.args, yearList[0])
	for i, size := 1, len(yearList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, yearList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) YearIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`year` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) YearIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`year` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) YearGT(year time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`year`>?")
	b.args = append(b.args, year)
	return b
}

func (b *GomodelWhereBuilder) YearGTE(year time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`year`>=?")
	b.args = append(b.args, year)
	return b
}

func (b *GomodelWhereBuilder) YearLT(year time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`year`<?")
	b.args = append(b.args, year)
	return b
}

func (b *GomodelWhereBuilder) YearLTE(year time.Time) *GomodelWhereBuilder {
	b.sb.WriteString("`year`<=?")
	b.args = append(b.args, year)
	return b
}

func (b *GomodelWhereBuilder) YearRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`year`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) CharEQ(char string) *GomodelWhereBuilder {
	b.sb.WriteString("`char`=?")
	b.args = append(b.args, char)
	return b
}

func (b *GomodelWhereBuilder) CharNEQ(char string) *GomodelWhereBuilder {
	b.sb.WriteString("`char`<>?")
	b.args = append(b.args, char)
	return b
}

func (b *GomodelWhereBuilder) CharIn(charList ...string) *GomodelWhereBuilder {
	if len(charList) == 0 {
		return b
	}

	b.sb.WriteString("`char` IN (?")
	b.args = append(b.args, charList[0])
	for i, size := 1, len(charList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, charList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) CharNotIn(charList ...string) *GomodelWhereBuilder {
	if len(charList) == 0 {
		return b
	}

	b.sb.WriteString("`char` NOT IN (?")
	b.args = append(b.args, charList[0])
	for i, size := 1, len(charList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, charList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) CharIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`char` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) CharIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`char` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) CharLike(char string) *GomodelWhereBuilder {
	b.sb.WriteString("`char` LIKE ?")
	b.args = append(b.args, char)
	return b
}

func (b *GomodelWhereBuilder) CharNotLike(char string) *GomodelWhereBuilder {
	b.sb.WriteString("`char` NOT LIKE ?")
	b.args = append(b.args, char)
	return b
}

func (b *GomodelWhereBuilder) CharRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`char`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) VarcharEQ(varchar string) *GomodelWhereBuilder {
	b.sb.WriteString("`varchar`=?")
	b.args = append(b.args, varchar)
	return b
}

func (b *GomodelWhereBuilder) VarcharNEQ(varchar string) *GomodelWhereBuilder {
	b.sb.WriteString("`varchar`<>?")
	b.args = append(b.args, varchar)
	return b
}

func (b *GomodelWhereBuilder) VarcharIn(varcharList ...string) *GomodelWhereBuilder {
	if len(varcharList) == 0 {
		return b
	}

	b.sb.WriteString("`varchar` IN (?")
	b.args = append(b.args, varcharList[0])
	for i, size := 1, len(varcharList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, varcharList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) VarcharNotIn(varcharList ...string) *GomodelWhereBuilder {
	if len(varcharList) == 0 {
		return b
	}

	b.sb.WriteString("`varchar` NOT IN (?")
	b.args = append(b.args, varcharList[0])
	for i, size := 1, len(varcharList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, varcharList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) VarcharIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`varchar` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) VarcharIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`varchar` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) VarcharLike(varchar string) *GomodelWhereBuilder {
	b.sb.WriteString("`varchar` LIKE ?")
	b.args = append(b.args, varchar)
	return b
}

func (b *GomodelWhereBuilder) VarcharNotLike(varchar string) *GomodelWhereBuilder {
	b.sb.WriteString("`varchar` NOT LIKE ?")
	b.args = append(b.args, varchar)
	return b
}

func (b *GomodelWhereBuilder) VarcharRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`varchar`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) BinaryEQ(binary string) *GomodelWhereBuilder {
	b.sb.WriteString("`binary`=?")
	b.args = append(b.args, binary)
	return b
}

func (b *GomodelWhereBuilder) BinaryNEQ(binary string) *GomodelWhereBuilder {
	b.sb.WriteString("`binary`<>?")
	b.args = append(b.args, binary)
	return b
}

func (b *GomodelWhereBuilder) BinaryIn(binaryList ...string) *GomodelWhereBuilder {
	if len(binaryList) == 0 {
		return b
	}

	b.sb.WriteString("`binary` IN (?")
	b.args = append(b.args, binaryList[0])
	for i, size := 1, len(binaryList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, binaryList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BinaryNotIn(binaryList ...string) *GomodelWhereBuilder {
	if len(binaryList) == 0 {
		return b
	}

	b.sb.WriteString("`binary` NOT IN (?")
	b.args = append(b.args, binaryList[0])
	for i, size := 1, len(binaryList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, binaryList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BinaryIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`binary` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) BinaryIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`binary` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) BinaryLike(binary string) *GomodelWhereBuilder {
	b.sb.WriteString("`binary` LIKE ?")
	b.args = append(b.args, binary)
	return b
}

func (b *GomodelWhereBuilder) BinaryNotLike(binary string) *GomodelWhereBuilder {
	b.sb.WriteString("`binary` NOT LIKE ?")
	b.args = append(b.args, binary)
	return b
}

func (b *GomodelWhereBuilder) BinaryRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`binary`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) VarbinaryEQ(varbinary string) *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary`=?")
	b.args = append(b.args, varbinary)
	return b
}

func (b *GomodelWhereBuilder) VarbinaryNEQ(varbinary string) *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary`<>?")
	b.args = append(b.args, varbinary)
	return b
}

func (b *GomodelWhereBuilder) VarbinaryIn(varbinaryList ...string) *GomodelWhereBuilder {
	if len(varbinaryList) == 0 {
		return b
	}

	b.sb.WriteString("`varbinary` IN (?")
	b.args = append(b.args, varbinaryList[0])
	for i, size := 1, len(varbinaryList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, varbinaryList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) VarbinaryNotIn(varbinaryList ...string) *GomodelWhereBuilder {
	if len(varbinaryList) == 0 {
		return b
	}

	b.sb.WriteString("`varbinary` NOT IN (?")
	b.args = append(b.args, varbinaryList[0])
	for i, size := 1, len(varbinaryList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, varbinaryList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) VarbinaryIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) VarbinaryIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) VarbinaryLike(varbinary string) *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary` LIKE ?")
	b.args = append(b.args, varbinary)
	return b
}

func (b *GomodelWhereBuilder) VarbinaryNotLike(varbinary string) *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary` NOT LIKE ?")
	b.args = append(b.args, varbinary)
	return b
}

func (b *GomodelWhereBuilder) VarbinaryRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`varbinary`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) TinyblobEQ(tinyblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob`=?")
	b.args = append(b.args, tinyblob)
	return b
}

func (b *GomodelWhereBuilder) TinyblobNEQ(tinyblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob`<>?")
	b.args = append(b.args, tinyblob)
	return b
}

func (b *GomodelWhereBuilder) TinyblobIn(tinyblobList ...string) *GomodelWhereBuilder {
	if len(tinyblobList) == 0 {
		return b
	}

	b.sb.WriteString("`tinyblob` IN (?")
	b.args = append(b.args, tinyblobList[0])
	for i, size := 1, len(tinyblobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinyblobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyblobNotIn(tinyblobList ...string) *GomodelWhereBuilder {
	if len(tinyblobList) == 0 {
		return b
	}

	b.sb.WriteString("`tinyblob` NOT IN (?")
	b.args = append(b.args, tinyblobList[0])
	for i, size := 1, len(tinyblobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinyblobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyblobIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) TinyblobIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) TinyblobLike(tinyblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob` LIKE ?")
	b.args = append(b.args, tinyblob)
	return b
}

func (b *GomodelWhereBuilder) TinyblobNotLike(tinyblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob` NOT LIKE ?")
	b.args = append(b.args, tinyblob)
	return b
}

func (b *GomodelWhereBuilder) TinyblobRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`tinyblob`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) TinytextEQ(tinytext string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext`=?")
	b.args = append(b.args, tinytext)
	return b
}

func (b *GomodelWhereBuilder) TinytextNEQ(tinytext string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext`<>?")
	b.args = append(b.args, tinytext)
	return b
}

func (b *GomodelWhereBuilder) TinytextIn(tinytextList ...string) *GomodelWhereBuilder {
	if len(tinytextList) == 0 {
		return b
	}

	b.sb.WriteString("`tinytext` IN (?")
	b.args = append(b.args, tinytextList[0])
	for i, size := 1, len(tinytextList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinytextList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinytextNotIn(tinytextList ...string) *GomodelWhereBuilder {
	if len(tinytextList) == 0 {
		return b
	}

	b.sb.WriteString("`tinytext` NOT IN (?")
	b.args = append(b.args, tinytextList[0])
	for i, size := 1, len(tinytextList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinytextList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinytextIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) TinytextIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) TinytextLike(tinytext string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext` LIKE ?")
	b.args = append(b.args, tinytext)
	return b
}

func (b *GomodelWhereBuilder) TinytextNotLike(tinytext string) *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext` NOT LIKE ?")
	b.args = append(b.args, tinytext)
	return b
}

func (b *GomodelWhereBuilder) TinytextRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`tinytext`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) BlobEQ(blob string) *GomodelWhereBuilder {
	b.sb.WriteString("`blob`=?")
	b.args = append(b.args, blob)
	return b
}

func (b *GomodelWhereBuilder) BlobNEQ(blob string) *GomodelWhereBuilder {
	b.sb.WriteString("`blob`<>?")
	b.args = append(b.args, blob)
	return b
}

func (b *GomodelWhereBuilder) BlobIn(blobList ...string) *GomodelWhereBuilder {
	if len(blobList) == 0 {
		return b
	}

	b.sb.WriteString("`blob` IN (?")
	b.args = append(b.args, blobList[0])
	for i, size := 1, len(blobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, blobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BlobNotIn(blobList ...string) *GomodelWhereBuilder {
	if len(blobList) == 0 {
		return b
	}

	b.sb.WriteString("`blob` NOT IN (?")
	b.args = append(b.args, blobList[0])
	for i, size := 1, len(blobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, blobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BlobIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`blob` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) BlobIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`blob` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) BlobLike(blob string) *GomodelWhereBuilder {
	b.sb.WriteString("`blob` LIKE ?")
	b.args = append(b.args, blob)
	return b
}

func (b *GomodelWhereBuilder) BlobNotLike(blob string) *GomodelWhereBuilder {
	b.sb.WriteString("`blob` NOT LIKE ?")
	b.args = append(b.args, blob)
	return b
}

func (b *GomodelWhereBuilder) BlobRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`blob`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) TextEQ(text string) *GomodelWhereBuilder {
	b.sb.WriteString("`text`=?")
	b.args = append(b.args, text)
	return b
}

func (b *GomodelWhereBuilder) TextNEQ(text string) *GomodelWhereBuilder {
	b.sb.WriteString("`text`<>?")
	b.args = append(b.args, text)
	return b
}

func (b *GomodelWhereBuilder) TextIn(textList ...string) *GomodelWhereBuilder {
	if len(textList) == 0 {
		return b
	}

	b.sb.WriteString("`text` IN (?")
	b.args = append(b.args, textList[0])
	for i, size := 1, len(textList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, textList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TextNotIn(textList ...string) *GomodelWhereBuilder {
	if len(textList) == 0 {
		return b
	}

	b.sb.WriteString("`text` NOT IN (?")
	b.args = append(b.args, textList[0])
	for i, size := 1, len(textList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, textList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TextIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`text` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) TextIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`text` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) TextLike(text string) *GomodelWhereBuilder {
	b.sb.WriteString("`text` LIKE ?")
	b.args = append(b.args, text)
	return b
}

func (b *GomodelWhereBuilder) TextNotLike(text string) *GomodelWhereBuilder {
	b.sb.WriteString("`text` NOT LIKE ?")
	b.args = append(b.args, text)
	return b
}

func (b *GomodelWhereBuilder) TextRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`text`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) MediumblobEQ(mediumblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob`=?")
	b.args = append(b.args, mediumblob)
	return b
}

func (b *GomodelWhereBuilder) MediumblobNEQ(mediumblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob`<>?")
	b.args = append(b.args, mediumblob)
	return b
}

func (b *GomodelWhereBuilder) MediumblobIn(mediumblobList ...string) *GomodelWhereBuilder {
	if len(mediumblobList) == 0 {
		return b
	}

	b.sb.WriteString("`mediumblob` IN (?")
	b.args = append(b.args, mediumblobList[0])
	for i, size := 1, len(mediumblobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, mediumblobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) MediumblobNotIn(mediumblobList ...string) *GomodelWhereBuilder {
	if len(mediumblobList) == 0 {
		return b
	}

	b.sb.WriteString("`mediumblob` NOT IN (?")
	b.args = append(b.args, mediumblobList[0])
	for i, size := 1, len(mediumblobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, mediumblobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) MediumblobIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) MediumblobIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) MediumblobLike(mediumblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob` LIKE ?")
	b.args = append(b.args, mediumblob)
	return b
}

func (b *GomodelWhereBuilder) MediumblobNotLike(mediumblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob` NOT LIKE ?")
	b.args = append(b.args, mediumblob)
	return b
}

func (b *GomodelWhereBuilder) MediumblobRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumblob`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) MediumtextEQ(mediumtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext`=?")
	b.args = append(b.args, mediumtext)
	return b
}

func (b *GomodelWhereBuilder) MediumtextNEQ(mediumtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext`<>?")
	b.args = append(b.args, mediumtext)
	return b
}

func (b *GomodelWhereBuilder) MediumtextIn(mediumtextList ...string) *GomodelWhereBuilder {
	if len(mediumtextList) == 0 {
		return b
	}

	b.sb.WriteString("`mediumtext` IN (?")
	b.args = append(b.args, mediumtextList[0])
	for i, size := 1, len(mediumtextList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, mediumtextList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) MediumtextNotIn(mediumtextList ...string) *GomodelWhereBuilder {
	if len(mediumtextList) == 0 {
		return b
	}

	b.sb.WriteString("`mediumtext` NOT IN (?")
	b.args = append(b.args, mediumtextList[0])
	for i, size := 1, len(mediumtextList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, mediumtextList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) MediumtextIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) MediumtextIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) MediumtextLike(mediumtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext` LIKE ?")
	b.args = append(b.args, mediumtext)
	return b
}

func (b *GomodelWhereBuilder) MediumtextNotLike(mediumtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext` NOT LIKE ?")
	b.args = append(b.args, mediumtext)
	return b
}

func (b *GomodelWhereBuilder) MediumtextRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`mediumtext`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) LongblobEQ(longblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`longblob`=?")
	b.args = append(b.args, longblob)
	return b
}

func (b *GomodelWhereBuilder) LongblobNEQ(longblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`longblob`<>?")
	b.args = append(b.args, longblob)
	return b
}

func (b *GomodelWhereBuilder) LongblobIn(longblobList ...string) *GomodelWhereBuilder {
	if len(longblobList) == 0 {
		return b
	}

	b.sb.WriteString("`longblob` IN (?")
	b.args = append(b.args, longblobList[0])
	for i, size := 1, len(longblobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, longblobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) LongblobNotIn(longblobList ...string) *GomodelWhereBuilder {
	if len(longblobList) == 0 {
		return b
	}

	b.sb.WriteString("`longblob` NOT IN (?")
	b.args = append(b.args, longblobList[0])
	for i, size := 1, len(longblobList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, longblobList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) LongblobIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`longblob` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) LongblobIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`longblob` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) LongblobLike(longblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`longblob` LIKE ?")
	b.args = append(b.args, longblob)
	return b
}

func (b *GomodelWhereBuilder) LongblobNotLike(longblob string) *GomodelWhereBuilder {
	b.sb.WriteString("`longblob` NOT LIKE ?")
	b.args = append(b.args, longblob)
	return b
}

func (b *GomodelWhereBuilder) LongblobRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`longblob`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) LongtextEQ(longtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`longtext`=?")
	b.args = append(b.args, longtext)
	return b
}

func (b *GomodelWhereBuilder) LongtextNEQ(longtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`longtext`<>?")
	b.args = append(b.args, longtext)
	return b
}

func (b *GomodelWhereBuilder) LongtextIn(longtextList ...string) *GomodelWhereBuilder {
	if len(longtextList) == 0 {
		return b
	}

	b.sb.WriteString("`longtext` IN (?")
	b.args = append(b.args, longtextList[0])
	for i, size := 1, len(longtextList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, longtextList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) LongtextNotIn(longtextList ...string) *GomodelWhereBuilder {
	if len(longtextList) == 0 {
		return b
	}

	b.sb.WriteString("`longtext` NOT IN (?")
	b.args = append(b.args, longtextList[0])
	for i, size := 1, len(longtextList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, longtextList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) LongtextIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`longtext` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) LongtextIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`longtext` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) LongtextLike(longtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`longtext` LIKE ?")
	b.args = append(b.args, longtext)
	return b
}

func (b *GomodelWhereBuilder) LongtextNotLike(longtext string) *GomodelWhereBuilder {
	b.sb.WriteString("`longtext` NOT LIKE ?")
	b.args = append(b.args, longtext)
	return b
}

func (b *GomodelWhereBuilder) LongtextRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`longtext`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) EnumEQ(enum string) *GomodelWhereBuilder {
	b.sb.WriteString("`enum`=?")
	b.args = append(b.args, enum)
	return b
}

func (b *GomodelWhereBuilder) EnumNEQ(enum string) *GomodelWhereBuilder {
	b.sb.WriteString("`enum`<>?")
	b.args = append(b.args, enum)
	return b
}

func (b *GomodelWhereBuilder) EnumIn(enumList ...string) *GomodelWhereBuilder {
	if len(enumList) == 0 {
		return b
	}

	b.sb.WriteString("`enum` IN (?")
	b.args = append(b.args, enumList[0])
	for i, size := 1, len(enumList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, enumList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) EnumNotIn(enumList ...string) *GomodelWhereBuilder {
	if len(enumList) == 0 {
		return b
	}

	b.sb.WriteString("`enum` NOT IN (?")
	b.args = append(b.args, enumList[0])
	for i, size := 1, len(enumList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, enumList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) EnumIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`enum` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) EnumIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`enum` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) EnumLike(enum string) *GomodelWhereBuilder {
	b.sb.WriteString("`enum` LIKE ?")
	b.args = append(b.args, enum)
	return b
}

func (b *GomodelWhereBuilder) EnumNotLike(enum string) *GomodelWhereBuilder {
	b.sb.WriteString("`enum` NOT LIKE ?")
	b.args = append(b.args, enum)
	return b
}

func (b *GomodelWhereBuilder) EnumRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`enum`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) SetEQ(set string) *GomodelWhereBuilder {
	b.sb.WriteString("`set`=?")
	b.args = append(b.args, set)
	return b
}

func (b *GomodelWhereBuilder) SetNEQ(set string) *GomodelWhereBuilder {
	b.sb.WriteString("`set`<>?")
	b.args = append(b.args, set)
	return b
}

func (b *GomodelWhereBuilder) SetIn(setList ...string) *GomodelWhereBuilder {
	if len(setList) == 0 {
		return b
	}

	b.sb.WriteString("`set` IN (?")
	b.args = append(b.args, setList[0])
	for i, size := 1, len(setList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, setList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) SetNotIn(setList ...string) *GomodelWhereBuilder {
	if len(setList) == 0 {
		return b
	}

	b.sb.WriteString("`set` NOT IN (?")
	b.args = append(b.args, setList[0])
	for i, size := 1, len(setList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, setList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) SetIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`set` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) SetIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`set` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) SetLike(set string) *GomodelWhereBuilder {
	b.sb.WriteString("`set` LIKE ?")
	b.args = append(b.args, set)
	return b
}

func (b *GomodelWhereBuilder) SetNotLike(set string) *GomodelWhereBuilder {
	b.sb.WriteString("`set` NOT LIKE ?")
	b.args = append(b.args, set)
	return b
}

func (b *GomodelWhereBuilder) SetRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`set`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) JsonEQ(json string) *GomodelWhereBuilder {
	b.sb.WriteString("`json`=?")
	b.args = append(b.args, json)
	return b
}

func (b *GomodelWhereBuilder) JsonNEQ(json string) *GomodelWhereBuilder {
	b.sb.WriteString("`json`<>?")
	b.args = append(b.args, json)
	return b
}

func (b *GomodelWhereBuilder) JsonIn(jsonList ...string) *GomodelWhereBuilder {
	if len(jsonList) == 0 {
		return b
	}

	b.sb.WriteString("`json` IN (?")
	b.args = append(b.args, jsonList[0])
	for i, size := 1, len(jsonList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, jsonList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) JsonNotIn(jsonList ...string) *GomodelWhereBuilder {
	if len(jsonList) == 0 {
		return b
	}

	b.sb.WriteString("`json` NOT IN (?")
	b.args = append(b.args, jsonList[0])
	for i, size := 1, len(jsonList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, jsonList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) JsonIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`json` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) JsonIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`json` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) JsonLike(json string) *GomodelWhereBuilder {
	b.sb.WriteString("`json` LIKE ?")
	b.args = append(b.args, json)
	return b
}

func (b *GomodelWhereBuilder) JsonNotLike(json string) *GomodelWhereBuilder {
	b.sb.WriteString("`json` NOT LIKE ?")
	b.args = append(b.args, json)
	return b
}

func (b *GomodelWhereBuilder) JsonRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`json`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) TinyboolEQ(tinybool bool) *GomodelWhereBuilder {
	b.sb.WriteString("`tinybool`=?")
	b.args = append(b.args, tinybool)
	return b
}

func (b *GomodelWhereBuilder) TinyboolNEQ(tinybool bool) *GomodelWhereBuilder {
	b.sb.WriteString("`tinybool`<>?")
	b.args = append(b.args, tinybool)
	return b
}

func (b *GomodelWhereBuilder) TinyboolIn(tinyboolList ...bool) *GomodelWhereBuilder {
	if len(tinyboolList) == 0 {
		return b
	}

	b.sb.WriteString("`tinybool` IN (?")
	b.args = append(b.args, tinyboolList[0])
	for i, size := 1, len(tinyboolList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinyboolList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyboolNotIn(tinyboolList ...bool) *GomodelWhereBuilder {
	if len(tinyboolList) == 0 {
		return b
	}

	b.sb.WriteString("`tinybool` NOT IN (?")
	b.args = append(b.args, tinyboolList[0])
	for i, size := 1, len(tinyboolList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, tinyboolList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) TinyboolIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`tinybool` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) TinyboolIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`tinybool` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) TinyboolRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`tinybool`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}
func (b *GomodelWhereBuilder) BoolEQ(bool bool) *GomodelWhereBuilder {
	b.sb.WriteString("`bool`=?")
	b.args = append(b.args, bool)
	return b
}

func (b *GomodelWhereBuilder) BoolNEQ(bool bool) *GomodelWhereBuilder {
	b.sb.WriteString("`bool`<>?")
	b.args = append(b.args, bool)
	return b
}

func (b *GomodelWhereBuilder) BoolIn(boolList ...bool) *GomodelWhereBuilder {
	if len(boolList) == 0 {
		return b
	}

	b.sb.WriteString("`bool` IN (?")
	b.args = append(b.args, boolList[0])
	for i, size := 1, len(boolList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, boolList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BoolNotIn(boolList ...bool) *GomodelWhereBuilder {
	if len(boolList) == 0 {
		return b
	}

	b.sb.WriteString("`bool` NOT IN (?")
	b.args = append(b.args, boolList[0])
	for i, size := 1, len(boolList); i < size; i++ {
		b.sb.WriteString(",?")
		b.args = append(b.args, boolList[i])
	}
	b.sb.WriteString(")")
	return b
}

func (b *GomodelWhereBuilder) BoolIsNull() *GomodelWhereBuilder {
	b.sb.WriteString("`bool` IS NULL")
	return b
}

func (b *GomodelWhereBuilder) BoolIsNotNull() *GomodelWhereBuilder {
	b.sb.WriteString("`bool` IS NOT NULL")
	return b
}

func (b *GomodelWhereBuilder) BoolRaw(raw string, args ...interface{}) *GomodelWhereBuilder {
	b.sb.WriteString("`bool`")
	b.sb.WriteString(raw)
	b.args = append(b.args, args...)
	return b
}

func (b *GomodelWhereBuilder) sql() (string, []interface{}) {
	return b.sb.String(), b.args
}
