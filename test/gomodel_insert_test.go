package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestGomodelInsertBuilder(t *testing.T) {
	Convey("InsertBuilder", t, func() {
		Convey("Should insert all fields by default", func() {
			sql, args := newGomodelInsertBuilder(nil).
				Values(&Gomodel{}).
				SQL()
			questions := make([]string, 0, len(GomodelFields))
			for range GomodelFields {
				questions = append(questions, "?")
			}
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf(
					"INSERT INTO `%s` (%s) VALUES (%s)",
					GomodelTable,
					strings.Join(GomodelFields, ","),
					strings.Join(questions, ","),
				),
			)
			So(args, ShouldHaveLength, len(GomodelFields))
		})

		Convey("Should control insert fields", func() {
			sql, args := newGomodelInsertBuilder(nil).
				Fields(GomodelFieldInt).
				Values(&Gomodel{
					Int: 1,
				}).
				SQL()
			So(sql, ShouldEqual, fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (?)", GomodelTable, GomodelFieldInt))
			So(args, ShouldHaveLength, 1)
			So(args[0], ShouldEqual, 1)

			sql, args = newGomodelInsertBuilder(nil).
				Fields(GomodelFieldBool, GomodelFieldUint).
				Values(&Gomodel{
					Bool: false,
					Uint: 1,
				}).
				SQL()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf("INSERT INTO `%s` (%s,%s) VALUES (?,?)", GomodelTable, GomodelFieldBool, GomodelFieldUint),
			)
			So(args, ShouldHaveLength, 2)
			So(args[0], ShouldEqual, false)
			So(args[1], ShouldEqual, 1)
		})
	})
}

func BenchmarkGomodelInsertBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newGomodelInsertBuilder(nil).
			Fields(GomodelFieldInt, GomodelFieldBigint, GomodelFieldFloat).
			Values(&Gomodel{}, &Gomodel{}).
			SQL()
	}
}

func BenchmarkInsertRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("INSERT INTO %s (%s,%s,%s) VALUES (?,?,?),(?,?,?)", GomodelTable, GomodelFieldInt,
			GomodelFieldBigint, GomodelFieldFloat)
	}
}
