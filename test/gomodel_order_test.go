package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGomodelOrderBuilder(t *testing.T) {
	Convey("OrderBuilder", t, func() {
		Convey("Should build correct statement on call once", func() {
			sql := newGomodelOrderBuilder().
				IntASC().
				sql()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf("ORDER BY %s ASC", GomodelFieldInt),
			)
		})

		Convey("Should build correct statement on call more than once", func() {
			sql := newGomodelOrderBuilder().
				IntASC().
				TinyintDESC().
				CharASC().
				VarcharDESC().
				sql()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf(
					"ORDER BY %s ASC, %s DESC, %s ASC, %s DESC",
					GomodelFieldInt,
					GomodelFieldTinyint,
					GomodelFieldChar,
					GomodelFieldVarchar,
				),
			)
		})
	})
}
