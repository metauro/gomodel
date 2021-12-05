package test

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestGomodelDB(t *testing.T) {
	gomodel := open().Gomodel
	ctx := context.Background()
	Convey("GomodelDB", t, func() {
		clean()
		Convey("CRUD", func() {
			id, err := gomodel.
				Insert().
				Fields(GomodelFieldTinyint).Values(&Gomodel{Tinyint: 1}).
				Exec(ctx)
			So(err, ShouldBeNil)
			So(id, ShouldEqual, 0)

			value, err := gomodel.
				Select().
				Fields(GomodelFieldTinyint).
				Get(ctx)
			So(err, ShouldBeNil)
			So(value, ShouldResemble, &Gomodel{Tinyint: 1})

			ra, err := gomodel.
				Update().
				IncrTinyint(10).
				Exec(ctx)
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1)

			values, err := gomodel.
				Select().
				Fields(GomodelFieldTinyint).
				List(ctx)
			So(err, ShouldBeNil)
			So(values, ShouldHaveLength, 1)
			So(values[0], ShouldResemble, &Gomodel{Tinyint: 11})

			ra, err = gomodel.
				Delete().
				Exec(ctx)
			So(err, ShouldBeNil)
			So(ra, ShouldEqual, 1)

			value, err = gomodel.
				Select().
				Get(ctx)
			So(err, ShouldEqual, ErrNoRows)
			So(value, ShouldBeNil)
		})

		Convey("Should trigger hooks", func() {
			gomodel := open().Gomodel
			insertBuilder := gomodel.Insert().Fields(GomodelFieldTinyint).Values(&Gomodel{})
			selectBuilder := gomodel.Select().Fields(GomodelFieldTinyint)
			updateBuilder := gomodel.Update().SetTinyint(1)
			deleteBuilder := gomodel.Delete()

			insertSQL, insertArgs := insertBuilder.SQL()
			selectSQL, selectArgs := selectBuilder.SQL()
			updateSQL, updateArgs := updateBuilder.SQL()
			deleteSQL, deleteArgs := deleteBuilder.SQL()

			hook := &testSetHook{
				beforeRun: make(map[Op]bool),
				afterRun:  make(map[Op]bool),
				sqlMap: map[Op]string{
					OpInsert: insertSQL,
					OpSelect: selectSQL,
					OpUpdate: updateSQL,
					OpDelete: deleteSQL,
				},
				argsMap: map[Op][]interface{}{
					OpInsert: insertArgs,
					OpSelect: selectArgs,
					OpUpdate: updateArgs,
					OpDelete: deleteArgs,
				},
			}
			gomodel.Use(hook)

			_, err := insertBuilder.Exec(ctx)
			So(err, ShouldBeNil)
			_, err = selectBuilder.List(ctx)
			So(err, ShouldBeNil)
			_, err = updateBuilder.Exec(ctx)
			So(err, ShouldBeNil)
			_, err = deleteBuilder.Exec(ctx)
			So(err, ShouldBeNil)
			So(hook.beforeRun, ShouldHaveLength, 4)
			So(hook.afterRun, ShouldHaveLength, 4)
		})
	})
}

type testSetHook struct {
	beforeRun map[Op]bool
	afterRun  map[Op]bool
	sqlMap    map[Op]string
	argsMap   map[Op][]interface{}
	ctx       context.Context
}

func (l *testSetHook) Before(e BeforeExecuteEvent) error {
	l.beforeRun[e.Op()] = true

	So(e.Table(), ShouldEqual, GomodelTable)
	So(e.Context(), ShouldNotBeNil)
	So(e.SQL(), ShouldEqual, l.sqlMap[e.Op()])
	So(e.Args(), ShouldResemble, l.argsMap[e.Op()])

	e.SetContext(context.WithValue(e.Context(), "op", e.Op()))
	e.SetTable("test")
	return nil
}

func (l *testSetHook) After(e AfterExecuteEvent) error {
	l.afterRun[e.Op()] = true

	So(e.Table(), ShouldEqual, "test")
	So(e.Context().Value("op"), ShouldEqual, e.Op())
	So(e.SQL(), ShouldEqual, strings.ReplaceAll(l.sqlMap[e.Op()], GomodelTable, "test"))
	So(e.Args(), ShouldResemble, l.argsMap[e.Op()])

	return nil
}
