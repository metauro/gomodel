package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGomodelWhereBuilder(t *testing.T) {
	Convey("WhereBuilder", t, func() {
		Convey("Should build correct statement on call once", func() {
			sql, args := newGomodelWhereBuilder().
				TinyintEQ(1).
				sql()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf("WHERE %s=?", GomodelFieldTinyint),
			)
			So(args, ShouldHaveLength, 1)
		})

		Convey("Should build correct statement on call more than once", func() {
			sql, args := newGomodelWhereBuilder().
				TinyintEQ(1).
				And().
				BigintGTE(10).
				Or().
				ConditionGroup(func(b *GomodelWhereBuilder) {
					b.
						TinyintIn(1, 2, 3).
						Or().
						TinyintNotIn(2, 4)
				}).
				And().
				VarcharRaw("=CONCAT(?,?)", "varchar", "").
				sql()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf(
					"WHERE %s=? AND %s>=? OR (%s IN (?,?,?) OR %s NOT IN (?,?)) AND %s=CONCAT(?,?)",
					GomodelFieldTinyint,
					GomodelFieldBigint,
					GomodelFieldTinyint,
					GomodelFieldTinyint,
					GomodelFieldVarchar,
				),
			)
			So(args, ShouldHaveLength, 9)
		})
	})
}
