package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/volatiletech/null/v9"
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
					Int: null.NewInt(1, true),
				}).
				SQL()
			So(sql, ShouldEqual, fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (?)", GomodelTable, GomodelFieldInt))
			So(args, ShouldHaveLength, 1)
			So(args[0], ShouldResemble, null.NewInt(1, true))

			sql, args = newGomodelInsertBuilder(nil).
				Fields(GomodelFieldBool, GomodelFieldUint).
				Values(&Gomodel{
					Bool: null.NewBool(false, false),
					Uint: null.NewUint(1, true),
				}).
				SQL()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf("INSERT INTO `%s` (%s,%s) VALUES (?,?)", GomodelTable, GomodelFieldBool, GomodelFieldUint),
			)
			So(args, ShouldHaveLength, 2)
			So(args[0], ShouldResemble, null.NewBool(false, false))
			So(args[1], ShouldResemble, null.NewUint(1, true))
		})

		Convey("Should build correct statement on insert multiple values", func() {
			sql, args := newGomodelInsertBuilder(nil).
				Fields(GomodelFieldChar, GomodelFieldVarchar).
				Values(&Gomodel{}, &Gomodel{}).
				SQL()
			So(
				sql,
				ShouldEqual,
				fmt.Sprintf(
					"INSERT INTO `%s` (%s,%s) VALUES (?,?),(?,?)",
					GomodelTable,
					GomodelFieldChar,
					GomodelFieldVarchar,
				),
			)
			So(args, ShouldHaveLength, 4)
		})
	})
}
