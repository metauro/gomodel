package test

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestGomodelSelectBuilder(t *testing.T) {
	Convey("SelectBuilder", t, func() {
		Convey("Should build correct statement", func() {
			sql, args := newGomodelSelectBuilder(nil).SQL()
			So(sql, ShouldEqual, fmt.Sprintf("SELECT %s FROM `%s`", strings.Join(GomodelFields, ","), GomodelTable))
			So(args, ShouldHaveLength, 0)
		})

		Convey("Should option work fine", func() {
			builder := newGomodelSelectBuilder(db.Gomodel).
				Distinct(true).
				Fields(GomodelFieldInt).
				Where(func(b *GomodelWhereBuilder) {
					b.IntEQ(1)
				}).
				Order(func(b *GomodelOrderBuilder) {
					b.IntDESC()
				}).
				Limit(2).
				Offset(3)
			sql, args := builder.SQL()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf(
					"SELECT DISTINCT %s FROM `%s` WHERE %s=? ORDER BY %s DESC LIMIT ? OFFSET ?",
					GomodelFieldInt,
					GomodelTable,
					GomodelFieldInt,
					GomodelFieldInt,
				),
			)
			So(args, ShouldHaveLength, 3)
			So(args[0], ShouldEqual, 1)
			So(args[1], ShouldEqual, 2)
			So(args[2], ShouldEqual, 3)

			list, err := builder.List(context.Background())
			So(err, ShouldBeNil)
			So(list, ShouldHaveLength, 0)
		})

		Convey("Should omit fields", func() {
			sql, args := newGomodelSelectBuilder(nil).
				OmitFields(GomodelFieldInt).
				SQL()
			fields := make([]string, 0, len(GomodelFields))
			for _, field := range GomodelFields {
				if field == GomodelFieldInt {
					continue
				}
				fields = append(fields, field)
			}
			So(sql, ShouldEqual, fmt.Sprintf("SELECT %s FROM `%s`", strings.Join(fields, ","), GomodelTable))
			So(args, ShouldHaveLength, 0)
		})
	})
}
