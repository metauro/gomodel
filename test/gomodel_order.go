package test

import "strings"

type GomodelOrderBuilder struct {
	sb      strings.Builder
	noFirst bool
}

func newGomodelOrderBuilder() *GomodelOrderBuilder {
	b := &GomodelOrderBuilder{}
	b.sb.WriteString("ORDER BY")
	return b
}

func (b *GomodelOrderBuilder) TinyintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinyint` ASC")
	return b
}

func (b *GomodelOrderBuilder) TinyintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinyint` DESC")
	return b
}

func (b *GomodelOrderBuilder) SmallintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `smallint` ASC")
	return b
}

func (b *GomodelOrderBuilder) SmallintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `smallint` DESC")
	return b
}

func (b *GomodelOrderBuilder) MediumintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `mediumint` ASC")
	return b
}

func (b *GomodelOrderBuilder) MediumintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `mediumint` DESC")
	return b
}

func (b *GomodelOrderBuilder) IntASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `int` ASC")
	return b
}

func (b *GomodelOrderBuilder) IntDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `int` DESC")
	return b
}

func (b *GomodelOrderBuilder) BigintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `bigint` ASC")
	return b
}

func (b *GomodelOrderBuilder) BigintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `bigint` DESC")
	return b
}

func (b *GomodelOrderBuilder) FloatASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `float` ASC")
	return b
}

func (b *GomodelOrderBuilder) FloatDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `float` DESC")
	return b
}

func (b *GomodelOrderBuilder) DoubleASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `double` ASC")
	return b
}

func (b *GomodelOrderBuilder) DoubleDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `double` DESC")
	return b
}

func (b *GomodelOrderBuilder) DecimalASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `decimal` ASC")
	return b
}

func (b *GomodelOrderBuilder) DecimalDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `decimal` DESC")
	return b
}

func (b *GomodelOrderBuilder) UtinyintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `utinyint` ASC")
	return b
}

func (b *GomodelOrderBuilder) UtinyintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `utinyint` DESC")
	return b
}

func (b *GomodelOrderBuilder) UsmallintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `usmallint` ASC")
	return b
}

func (b *GomodelOrderBuilder) UsmallintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `usmallint` DESC")
	return b
}

func (b *GomodelOrderBuilder) UmediumintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `umediumint` ASC")
	return b
}

func (b *GomodelOrderBuilder) UmediumintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `umediumint` DESC")
	return b
}

func (b *GomodelOrderBuilder) UintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `uint` ASC")
	return b
}

func (b *GomodelOrderBuilder) UintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `uint` DESC")
	return b
}

func (b *GomodelOrderBuilder) UbigintASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `ubigint` ASC")
	return b
}

func (b *GomodelOrderBuilder) UbigintDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `ubigint` DESC")
	return b
}

func (b *GomodelOrderBuilder) UfloatASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `ufloat` ASC")
	return b
}

func (b *GomodelOrderBuilder) UfloatDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `ufloat` DESC")
	return b
}

func (b *GomodelOrderBuilder) UdoubleASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `udouble` ASC")
	return b
}

func (b *GomodelOrderBuilder) UdoubleDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `udouble` DESC")
	return b
}

func (b *GomodelOrderBuilder) UdecimalASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `udecimal` ASC")
	return b
}

func (b *GomodelOrderBuilder) UdecimalDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `udecimal` DESC")
	return b
}

func (b *GomodelOrderBuilder) DateASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `date` ASC")
	return b
}

func (b *GomodelOrderBuilder) DateDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `date` DESC")
	return b
}

func (b *GomodelOrderBuilder) DatetimeASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `datetime` ASC")
	return b
}

func (b *GomodelOrderBuilder) DatetimeDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `datetime` DESC")
	return b
}

func (b *GomodelOrderBuilder) TimestampASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `timestamp` ASC")
	return b
}

func (b *GomodelOrderBuilder) TimestampDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `timestamp` DESC")
	return b
}

func (b *GomodelOrderBuilder) TimeASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `time` ASC")
	return b
}

func (b *GomodelOrderBuilder) TimeDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `time` DESC")
	return b
}

func (b *GomodelOrderBuilder) YearASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `year` ASC")
	return b
}

func (b *GomodelOrderBuilder) YearDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `year` DESC")
	return b
}

func (b *GomodelOrderBuilder) CharASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `char` ASC")
	return b
}

func (b *GomodelOrderBuilder) CharDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `char` DESC")
	return b
}

func (b *GomodelOrderBuilder) VarcharASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `varchar` ASC")
	return b
}

func (b *GomodelOrderBuilder) VarcharDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `varchar` DESC")
	return b
}

func (b *GomodelOrderBuilder) BinaryASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `binary` ASC")
	return b
}

func (b *GomodelOrderBuilder) BinaryDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `binary` DESC")
	return b
}

func (b *GomodelOrderBuilder) VarbinaryASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `varbinary` ASC")
	return b
}

func (b *GomodelOrderBuilder) VarbinaryDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `varbinary` DESC")
	return b
}

func (b *GomodelOrderBuilder) TinyblobASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinyblob` ASC")
	return b
}

func (b *GomodelOrderBuilder) TinyblobDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinyblob` DESC")
	return b
}

func (b *GomodelOrderBuilder) TinytextASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinytext` ASC")
	return b
}

func (b *GomodelOrderBuilder) TinytextDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinytext` DESC")
	return b
}

func (b *GomodelOrderBuilder) BlobASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `blob` ASC")
	return b
}

func (b *GomodelOrderBuilder) BlobDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `blob` DESC")
	return b
}

func (b *GomodelOrderBuilder) TextASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `text` ASC")
	return b
}

func (b *GomodelOrderBuilder) TextDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `text` DESC")
	return b
}

func (b *GomodelOrderBuilder) MediumblobASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `mediumblob` ASC")
	return b
}

func (b *GomodelOrderBuilder) MediumblobDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `mediumblob` DESC")
	return b
}

func (b *GomodelOrderBuilder) MediumtextASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `mediumtext` ASC")
	return b
}

func (b *GomodelOrderBuilder) MediumtextDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `mediumtext` DESC")
	return b
}

func (b *GomodelOrderBuilder) LongblobASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `longblob` ASC")
	return b
}

func (b *GomodelOrderBuilder) LongblobDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `longblob` DESC")
	return b
}

func (b *GomodelOrderBuilder) LongtextASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `longtext` ASC")
	return b
}

func (b *GomodelOrderBuilder) LongtextDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `longtext` DESC")
	return b
}

func (b *GomodelOrderBuilder) EnumASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `enum` ASC")
	return b
}

func (b *GomodelOrderBuilder) EnumDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `enum` DESC")
	return b
}

func (b *GomodelOrderBuilder) SetASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `set` ASC")
	return b
}

func (b *GomodelOrderBuilder) SetDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `set` DESC")
	return b
}

func (b *GomodelOrderBuilder) JsonASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `json` ASC")
	return b
}

func (b *GomodelOrderBuilder) JsonDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `json` DESC")
	return b
}

func (b *GomodelOrderBuilder) TinyboolASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinybool` ASC")
	return b
}

func (b *GomodelOrderBuilder) TinyboolDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `tinybool` DESC")
	return b
}

func (b *GomodelOrderBuilder) BoolASC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `bool` ASC")
	return b
}

func (b *GomodelOrderBuilder) BoolDESC() *GomodelOrderBuilder {
	b.check()
	b.sb.WriteString(" `bool` DESC")
	return b
}

func (b *GomodelOrderBuilder) check() {
	if b.noFirst {
		b.sb.WriteString(",")
		return
	}

	b.noFirst = true
}

func (b *GomodelOrderBuilder) sql() string {
	return b.sb.String()
}
