package test

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var open = func() func() *DB {
	db, err := sqlx.Open("mysql", "root:root@(localhost:3306)/dev?parseTime=true")
	if err != nil {
		panic(err)
	}
	return func() *DB {
		return NewDB(db)
	}
}()

func clean() {
	_, _ = open().db.Exec(fmt.Sprintf("TRUNCATE TABLE `%s`", GomodelTable))
}

func TestDB(t *testing.T) {
	db := open()
	ctx := context.Background()

	Convey("DB", t, func() {
		clean()
		Convey("Should commit on transaction end", func() {
			err := db.BeginFn(ctx, func(ctx context.Context, tx *Tx) error {
				_, err := tx.Gomodel.Insert().Fields(GomodelFieldTinyint).Values(&Gomodel{Tinyint: 1}).Exec(ctx)
				So(err, ShouldBeNil)
				return nil
			})
			So(err, ShouldBeNil)
			val, err := db.Gomodel.Select().Get(ctx)
			So(err, ShouldBeNil)
			So(val.Tinyint, ShouldEqual, 1)
		})

		Convey("Should rollback if transaction return err", func() {
			err := db.BeginFn(ctx, func(ctx context.Context, tx *Tx) error {
				_, err := tx.Gomodel.Insert().Fields(GomodelFieldTinyint).Values(&Gomodel{Tinyint: 1}).Exec(ctx)
				So(err, ShouldBeNil)
				return errors.New("")
			})
			So(err, ShouldNotBeNil)
			_, err = db.Gomodel.Select().Get(ctx)
			So(err, ShouldEqual, ErrNoRows)
		})

		Convey("Should rollback if transaction panic", func() {
			defer func() {
				err := recover()
				So(err, ShouldNotBeNil)
				_, err = db.Gomodel.Select().Get(ctx)
				So(err, ShouldEqual, ErrNoRows)
			}()

			_ = db.BeginFn(ctx, func(ctx context.Context, tx *Tx) error {
				_, _ = tx.Gomodel.Insert().Fields(GomodelFieldTinyint).Values(&Gomodel{Tinyint: 1}).Exec(ctx)
				panic("")
			})
		})
	})
}
