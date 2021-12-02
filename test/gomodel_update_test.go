package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGomodelUpdateBuilder(t *testing.T) {
	//getIndex := func(sql string, field string) int {
	//	sql = sql[0:strings.Index(sql, field)]
	//	return strings.Count(sql, ",")
	//}
	Convey("UpdateBuilder", t, func() {
		Convey("Should update all non zero value", func() {
			sql, args := newGomodelUpdateBuilder(nil).
				Set(&Gomodel{
					Tinyint: 1,
				}).
				SQL()
			So(sql, ShouldEqual, fmt.Sprintf("UPDATE `%s` SET %s=?", GomodelTable, GomodelFieldTinyint))
			So(args, ShouldHaveLength, 1)
			So(args[0], ShouldEqual, 1)
		})

		Convey("Should control update fields", func() {
			sql, args := newGomodelUpdateBuilder(nil).
				SetTinyint(1).
				IncrInt(2).
				DecrSmallint(3).
				SetVarcharZero().
				SetCharNil().
				SetBigintRaw("=?+?", 4, 5).
				SQL()
			So(args, ShouldHaveLength, 6)
			//So(args[0], ShouldEqual, 1)
			fmt.Println(sql)
		})
	})
}
