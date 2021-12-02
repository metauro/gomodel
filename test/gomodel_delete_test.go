package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGomodelDeleteBuilder(t *testing.T) {
	Convey("DeleteBuilder", t, func() {
		Convey("Should build correct statement", func() {
			sql, args := newGomodelDeleteBuilder(nil).SQL()
			So(sql, ShouldEqual, fmt.Sprintf("DELETE FROM `%s`", GomodelTable))
			So(args, ShouldHaveLength, 0)
		})

		Convey("Should where builder work fine", func() {
			sql, args := newGomodelDeleteBuilder(nil).
				Where(func(b *GomodelWhereBuilder) {
					b.TinyintGT(1)
				}).
				SQL()
			So(sql, ShouldEqual, fmt.Sprintf("DELETE FROM `%s` WHERE %s>?", GomodelTable, GomodelFieldTinyint))
			So(args, ShouldHaveLength, 1)
			So(args[0], ShouldEqual, 1)
		})
	})
}
