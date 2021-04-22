package model

import (
	"context"
	"database/sql"
	"github.com/metauro/gomodel"
	"strings"
	"time"
)

var _ = time.Time{}

const TestFieldId = "`id`"
const TestFieldCreateAt = "`create_at`"
const TestFieldUpdateAt = "`update_at`"

var TestFields = []string{
	TestFieldId,
	TestFieldCreateAt,
	TestFieldUpdateAt,
}

type Test struct {
	Id       int       `db:"id"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
}

type testRepo struct {
	db gomodel.Execer
}

func NewTestRepo(db gomodel.Execer) *testRepo {
	return &testRepo{
		db: db,
	}
}

type insertTestRepo struct {
	db         gomodel.Execer
	sqlBuilder strings.Builder
	bindings   []interface{}
}

type duplicateTestRepo struct {
	db         gomodel.Execer
	first      bool
	sqlBuilder strings.Builder
	bindings   []interface{}
}

type updateTestRepo struct {
	db         gomodel.Execer
	hasOrderBy bool
	hasWhere   bool
	whereCond  string
	first      bool
	sqlBuilder strings.Builder
	bindings   []interface{}
}

type selectTestRepo struct {
	db         gomodel.Execer
	hasOrderBy bool
	hasWhere   bool
	whereCond  string
	fields     []string
	sqlBuilder strings.Builder
	bindings   []interface{}
}

type deleteTestRepo struct {
	db         gomodel.Execer
	hasOrderBy bool
	hasWhere   bool
	whereCond  string
	sqlBuilder strings.Builder
	bindings   []interface{}
}

func (r *testRepo) Select() *selectTestRepo {
	res := &selectTestRepo{
		db:     r.db,
		fields: TestFields,
	}
	res.sqlBuilder.WriteString("SELECT `id`,`create_at`,`update_at` FROM `test` ")
	return res
}

func (r *testRepo) SelectDistinct() *selectTestRepo {
	res := &selectTestRepo{
		db:     r.db,
		fields: TestFields,
	}
	res.sqlBuilder.WriteString("SELECT DISTINCT `id`,`create_at`,`update_at` FROM `test` ")
	return res
}

func (r *testRepo) SelectPick(fields ...string) *selectTestRepo {
	res := &selectTestRepo{
		db:     r.db,
		fields: fields,
	}
	res.sqlBuilder.WriteString("SELECT ")
	for i, l := 0, len(fields)-1; i < l; i++ {
		res.sqlBuilder.WriteString(fields[i])
		res.sqlBuilder.WriteString(",")
	}
	res.sqlBuilder.WriteString(fields[len(fields)-1])
	res.sqlBuilder.WriteString(" FROM `test` ")
	return res
}

func (r *testRepo) SelectDistinctPick(fields ...string) *selectTestRepo {
	res := &selectTestRepo{
		db:     r.db,
		fields: fields,
	}
	res.sqlBuilder.WriteString("SELECT DISTINCT ")
	for i, l := 0, len(fields)-1; i < l; i++ {
		res.sqlBuilder.WriteString(fields[i])
		res.sqlBuilder.WriteString(",")
	}
	res.sqlBuilder.WriteString(fields[len(fields)-1])
	res.sqlBuilder.WriteString(" FROM `test` ")
	return res
}

func (r *testRepo) SelectOmit(fields ...string) *selectTestRepo {
	size := len(fields)
omit:
	for _, field := range TestFields {
		for _, omitField := range fields {
			if field == omitField {
				continue omit
			}
		}
		fields = append(fields, field)
	}
	fields = fields[size:]
	res := &selectTestRepo{
		db:     r.db,
		fields: fields,
	}
	res.sqlBuilder.WriteString("SELECT ")
	for i, l := 0, len(fields)-1; i < l; i++ {
		res.sqlBuilder.WriteString(fields[i])
		res.sqlBuilder.WriteString(",")
	}
	res.sqlBuilder.WriteString(fields[len(fields)-1])
	res.sqlBuilder.WriteString(" FROM `test` ")
	return res
}

func (r *testRepo) SelectDistinctOmit(fields ...string) *selectTestRepo {
	size := len(fields)
omit:
	for _, field := range TestFields {
		for _, omitField := range fields {
			if field == omitField {
				continue omit
			}
		}
		fields = append(fields, field)
	}
	fields = fields[size:]
	res := &selectTestRepo{
		db:     r.db,
		fields: fields,
	}
	res.sqlBuilder.WriteString("SELECT DISTINCT ")
	for i, l := 0, len(fields)-1; i < l; i++ {
		res.sqlBuilder.WriteString(fields[i])
		res.sqlBuilder.WriteString(",")
	}
	res.sqlBuilder.WriteString(fields[len(fields)-1])
	res.sqlBuilder.WriteString(" FROM `test` ")
	return res
}

func (r *testRepo) Insert(testList ...*Test) *insertTestRepo {
	res := &insertTestRepo{
		db: r.db,
	}
	res.sqlBuilder.WriteString("INSERT INTO `test` (`id`,`create_at`,`update_at`) VALUES ")
	for i, m := range testList {
		if i > 0 {
			res.sqlBuilder.WriteString(",")
		}
		res.sqlBuilder.WriteString("(")
		if m.Id == 0 {
			res.sqlBuilder.WriteString("DEFAULT")
		} else {
			res.sqlBuilder.WriteString("?")
			res.bindings = append(res.bindings, m.Id)
		}
		if m.CreateAt.IsZero() {
			res.sqlBuilder.WriteString(",DEFAULT")
		} else {
			res.sqlBuilder.WriteString(",?")
			res.bindings = append(res.bindings, m.CreateAt)
		}
		if m.UpdateAt.IsZero() {
			res.sqlBuilder.WriteString(",DEFAULT")
		} else {
			res.sqlBuilder.WriteString(",?")
			res.bindings = append(res.bindings, m.UpdateAt)
		}

		res.sqlBuilder.WriteString(")")
	}
	return res
}

func (r *testRepo) Update() *updateTestRepo {
	res := &updateTestRepo{
		db:    r.db,
		first: true,
	}
	res.sqlBuilder.WriteString("UPDATE `test` SET ")
	return res
}

func (r *testRepo) Delete() *deleteTestRepo {
	res := &deleteTestRepo{
		db: r.db,
	}
	res.sqlBuilder.WriteString("DELETE FROM `test` ")
	return res
}

func (r *selectTestRepo) And() *selectTestRepo {
	r.whereCond = "AND"
	return r
}

func (r *selectTestRepo) Or() *selectTestRepo {
	r.whereCond = "OR"
	return r
}

func (r *selectTestRepo) whereCheck() {
	if r.hasWhere {
		if r.whereCond == "" {
			r.And()
		}

		r.sqlBuilder.WriteString(" ")
		r.sqlBuilder.WriteString(r.whereCond)
		return
	}

	r.hasWhere = true
	r.sqlBuilder.WriteString(" WHERE")
}

func (r *selectTestRepo) WhereIdEqual(id int) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *selectTestRepo) WhereIdNotEqual(id int) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<>?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *selectTestRepo) WhereIdIn(idList ...int) *selectTestRepo {
	size := len(idList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, idList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, idList[size-1])
	return r
}

func (r *selectTestRepo) WhereIdNotIn(idList ...int) *selectTestRepo {
	size := len(idList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, idList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, idList[size-1])
	return r
}

func (r *selectTestRepo) WhereIdIsNil() *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IS NULL")
	return r
}

func (r *selectTestRepo) WhereIdIsNotNil() *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IS NOT NULL")
	return r
}

func (r *selectTestRepo) WhereIdRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) WhereIdGreatThan(id int) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`>?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *selectTestRepo) WhereIdGreatThanEqual(id int) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`>=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *selectTestRepo) WhereIdLessThan(id int) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *selectTestRepo) WhereIdLessThenEqual(id int) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *selectTestRepo) WhereCreateAtEqual(createAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *selectTestRepo) WhereCreateAtNotEqual(createAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<>?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *selectTestRepo) WhereCreateAtIn(createAtList ...time.Time) *selectTestRepo {
	size := len(createAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, createAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, createAtList[size-1])
	return r
}

func (r *selectTestRepo) WhereCreateAtNotIn(createAtList ...time.Time) *selectTestRepo {
	size := len(createAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, createAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, createAtList[size-1])
	return r
}

func (r *selectTestRepo) WhereCreateAtIsNil() *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IS NULL")
	return r
}

func (r *selectTestRepo) WhereCreateAtIsNotNil() *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IS NOT NULL")
	return r
}

func (r *selectTestRepo) WhereCreateAtRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) WhereCreateAtGreatThan(createAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`>?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *selectTestRepo) WhereCreateAtGreatThanEqual(createAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`>=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *selectTestRepo) WhereCreateAtLessThan(createAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *selectTestRepo) WhereCreateAtLessThenEqual(createAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *selectTestRepo) WhereUpdateAtEqual(updateAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *selectTestRepo) WhereUpdateAtNotEqual(updateAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<>?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *selectTestRepo) WhereUpdateAtIn(updateAtList ...time.Time) *selectTestRepo {
	size := len(updateAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, updateAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, updateAtList[size-1])
	return r
}

func (r *selectTestRepo) WhereUpdateAtNotIn(updateAtList ...time.Time) *selectTestRepo {
	size := len(updateAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, updateAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, updateAtList[size-1])
	return r
}

func (r *selectTestRepo) WhereUpdateAtIsNil() *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IS NULL")
	return r
}

func (r *selectTestRepo) WhereUpdateAtIsNotNil() *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IS NOT NULL")
	return r
}

func (r *selectTestRepo) WhereUpdateAtRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) WhereUpdateAtGreatThan(updateAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`>?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *selectTestRepo) WhereUpdateAtGreatThanEqual(updateAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`>=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *selectTestRepo) WhereUpdateAtLessThan(updateAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *selectTestRepo) WhereUpdateAtLessThenEqual(updateAt time.Time) *selectTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *selectTestRepo) orderByCheck() {
	if r.hasOrderBy {
		r.sqlBuilder.WriteString(",")
		return
	}

	r.hasOrderBy = true
	r.sqlBuilder.WriteString(" ORDER BY")
}

func (r *selectTestRepo) OrderByRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) OrderByIdAsc() *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` ASC")
	return r
}

func (r *selectTestRepo) OrderByIdDesc() *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` DESC")
	return r
}

func (r *selectTestRepo) OrderByIdRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) OrderByCreateAtAsc() *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` ASC")
	return r
}

func (r *selectTestRepo) OrderByCreateAtDesc() *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` DESC")
	return r
}

func (r *selectTestRepo) OrderByCreateAtRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) OrderByUpdateAtAsc() *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` ASC")
	return r
}

func (r *selectTestRepo) OrderByUpdateAtDesc() *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` DESC")
	return r
}

func (r *selectTestRepo) OrderByUpdateAtRaw(raw string, bindings ...interface{}) *selectTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) Limit(limit int) *selectTestRepo {
	r.sqlBuilder.WriteString(" LIMIT ?")
	r.bindings = append(r.bindings, limit)
	return r
}
func (r *updateTestRepo) And() *updateTestRepo {
	r.whereCond = "AND"
	return r
}

func (r *updateTestRepo) Or() *updateTestRepo {
	r.whereCond = "OR"
	return r
}

func (r *updateTestRepo) whereCheck() {
	if r.hasWhere {
		if r.whereCond == "" {
			r.And()
		}

		r.sqlBuilder.WriteString(" ")
		r.sqlBuilder.WriteString(r.whereCond)
		return
	}

	r.hasWhere = true
	r.sqlBuilder.WriteString(" WHERE")
}

func (r *updateTestRepo) WhereIdEqual(id int) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *updateTestRepo) WhereIdNotEqual(id int) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<>?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *updateTestRepo) WhereIdIn(idList ...int) *updateTestRepo {
	size := len(idList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, idList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, idList[size-1])
	return r
}

func (r *updateTestRepo) WhereIdNotIn(idList ...int) *updateTestRepo {
	size := len(idList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, idList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, idList[size-1])
	return r
}

func (r *updateTestRepo) WhereIdIsNil() *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IS NULL")
	return r
}

func (r *updateTestRepo) WhereIdIsNotNil() *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IS NOT NULL")
	return r
}

func (r *updateTestRepo) WhereIdRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) WhereIdGreatThan(id int) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`>?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *updateTestRepo) WhereIdGreatThanEqual(id int) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`>=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *updateTestRepo) WhereIdLessThan(id int) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *updateTestRepo) WhereIdLessThenEqual(id int) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *updateTestRepo) WhereCreateAtEqual(createAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *updateTestRepo) WhereCreateAtNotEqual(createAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<>?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *updateTestRepo) WhereCreateAtIn(createAtList ...time.Time) *updateTestRepo {
	size := len(createAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, createAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, createAtList[size-1])
	return r
}

func (r *updateTestRepo) WhereCreateAtNotIn(createAtList ...time.Time) *updateTestRepo {
	size := len(createAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, createAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, createAtList[size-1])
	return r
}

func (r *updateTestRepo) WhereCreateAtIsNil() *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IS NULL")
	return r
}

func (r *updateTestRepo) WhereCreateAtIsNotNil() *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IS NOT NULL")
	return r
}

func (r *updateTestRepo) WhereCreateAtRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) WhereCreateAtGreatThan(createAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`>?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *updateTestRepo) WhereCreateAtGreatThanEqual(createAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`>=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *updateTestRepo) WhereCreateAtLessThan(createAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *updateTestRepo) WhereCreateAtLessThenEqual(createAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *updateTestRepo) WhereUpdateAtEqual(updateAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *updateTestRepo) WhereUpdateAtNotEqual(updateAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<>?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *updateTestRepo) WhereUpdateAtIn(updateAtList ...time.Time) *updateTestRepo {
	size := len(updateAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, updateAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, updateAtList[size-1])
	return r
}

func (r *updateTestRepo) WhereUpdateAtNotIn(updateAtList ...time.Time) *updateTestRepo {
	size := len(updateAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, updateAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, updateAtList[size-1])
	return r
}

func (r *updateTestRepo) WhereUpdateAtIsNil() *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IS NULL")
	return r
}

func (r *updateTestRepo) WhereUpdateAtIsNotNil() *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IS NOT NULL")
	return r
}

func (r *updateTestRepo) WhereUpdateAtRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) WhereUpdateAtGreatThan(updateAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`>?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *updateTestRepo) WhereUpdateAtGreatThanEqual(updateAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`>=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *updateTestRepo) WhereUpdateAtLessThan(updateAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *updateTestRepo) WhereUpdateAtLessThenEqual(updateAt time.Time) *updateTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *updateTestRepo) orderByCheck() {
	if r.hasOrderBy {
		r.sqlBuilder.WriteString(",")
		return
	}

	r.hasOrderBy = true
	r.sqlBuilder.WriteString(" ORDER BY")
}

func (r *updateTestRepo) OrderByRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) OrderByIdAsc() *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` ASC")
	return r
}

func (r *updateTestRepo) OrderByIdDesc() *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` DESC")
	return r
}

func (r *updateTestRepo) OrderByIdRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) OrderByCreateAtAsc() *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` ASC")
	return r
}

func (r *updateTestRepo) OrderByCreateAtDesc() *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` DESC")
	return r
}

func (r *updateTestRepo) OrderByCreateAtRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) OrderByUpdateAtAsc() *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` ASC")
	return r
}

func (r *updateTestRepo) OrderByUpdateAtDesc() *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` DESC")
	return r
}

func (r *updateTestRepo) OrderByUpdateAtRaw(raw string, bindings ...interface{}) *updateTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *updateTestRepo) Limit(limit int) *updateTestRepo {
	r.sqlBuilder.WriteString(" LIMIT ?")
	r.bindings = append(r.bindings, limit)
	return r
}
func (r *deleteTestRepo) And() *deleteTestRepo {
	r.whereCond = "AND"
	return r
}

func (r *deleteTestRepo) Or() *deleteTestRepo {
	r.whereCond = "OR"
	return r
}

func (r *deleteTestRepo) whereCheck() {
	if r.hasWhere {
		if r.whereCond == "" {
			r.And()
		}

		r.sqlBuilder.WriteString(" ")
		r.sqlBuilder.WriteString(r.whereCond)
		return
	}

	r.hasWhere = true
	r.sqlBuilder.WriteString(" WHERE")
}

func (r *deleteTestRepo) WhereIdEqual(id int) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *deleteTestRepo) WhereIdNotEqual(id int) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<>?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *deleteTestRepo) WhereIdIn(idList ...int) *deleteTestRepo {
	size := len(idList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, idList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, idList[size-1])
	return r
}

func (r *deleteTestRepo) WhereIdNotIn(idList ...int) *deleteTestRepo {
	size := len(idList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, idList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, idList[size-1])
	return r
}

func (r *deleteTestRepo) WhereIdIsNil() *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IS NULL")
	return r
}

func (r *deleteTestRepo) WhereIdIsNotNil() *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` IS NOT NULL")
	return r
}

func (r *deleteTestRepo) WhereIdRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) WhereIdGreatThan(id int) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`>?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *deleteTestRepo) WhereIdGreatThanEqual(id int) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`>=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *deleteTestRepo) WhereIdLessThan(id int) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *deleteTestRepo) WhereIdLessThenEqual(id int) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `id`<=?")
	r.bindings = append(r.bindings, id)
	return r
}

func (r *deleteTestRepo) WhereCreateAtEqual(createAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *deleteTestRepo) WhereCreateAtNotEqual(createAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<>?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *deleteTestRepo) WhereCreateAtIn(createAtList ...time.Time) *deleteTestRepo {
	size := len(createAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, createAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, createAtList[size-1])
	return r
}

func (r *deleteTestRepo) WhereCreateAtNotIn(createAtList ...time.Time) *deleteTestRepo {
	size := len(createAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, createAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, createAtList[size-1])
	return r
}

func (r *deleteTestRepo) WhereCreateAtIsNil() *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IS NULL")
	return r
}

func (r *deleteTestRepo) WhereCreateAtIsNotNil() *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` IS NOT NULL")
	return r
}

func (r *deleteTestRepo) WhereCreateAtRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) WhereCreateAtGreatThan(createAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`>?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *deleteTestRepo) WhereCreateAtGreatThanEqual(createAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`>=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *deleteTestRepo) WhereCreateAtLessThan(createAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *deleteTestRepo) WhereCreateAtLessThenEqual(createAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `create_at`<=?")
	r.bindings = append(r.bindings, createAt)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtEqual(updateAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtNotEqual(updateAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<>?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtIn(updateAtList ...time.Time) *deleteTestRepo {
	size := len(updateAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, updateAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, updateAtList[size-1])
	return r
}

func (r *deleteTestRepo) WhereUpdateAtNotIn(updateAtList ...time.Time) *deleteTestRepo {
	size := len(updateAtList)
	if size == 0 {
		return r
	}

	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` NOT IN(")

	for i, l := 0, size-1; i < l; i++ {
		r.sqlBuilder.WriteString("?,")
		r.bindings = append(r.bindings, updateAtList[i])
	}
	r.sqlBuilder.WriteString("?)")
	r.bindings = append(r.bindings, updateAtList[size-1])
	return r
}

func (r *deleteTestRepo) WhereUpdateAtIsNil() *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IS NULL")
	return r
}

func (r *deleteTestRepo) WhereUpdateAtIsNotNil() *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` IS NOT NULL")
	return r
}

func (r *deleteTestRepo) WhereUpdateAtRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtGreatThan(updateAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`>?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtGreatThanEqual(updateAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`>=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtLessThan(updateAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *deleteTestRepo) WhereUpdateAtLessThenEqual(updateAt time.Time) *deleteTestRepo {
	r.whereCheck()
	r.sqlBuilder.WriteString(" `update_at`<=?")
	r.bindings = append(r.bindings, updateAt)
	return r
}

func (r *deleteTestRepo) orderByCheck() {
	if r.hasOrderBy {
		r.sqlBuilder.WriteString(",")
		return
	}

	r.hasOrderBy = true
	r.sqlBuilder.WriteString(" ORDER BY")
}

func (r *deleteTestRepo) OrderByRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) OrderByIdAsc() *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` ASC")
	return r
}

func (r *deleteTestRepo) OrderByIdDesc() *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` DESC")
	return r
}

func (r *deleteTestRepo) OrderByIdRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) OrderByCreateAtAsc() *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` ASC")
	return r
}

func (r *deleteTestRepo) OrderByCreateAtDesc() *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` DESC")
	return r
}

func (r *deleteTestRepo) OrderByCreateAtRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) OrderByUpdateAtAsc() *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` ASC")
	return r
}

func (r *deleteTestRepo) OrderByUpdateAtDesc() *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` DESC")
	return r
}

func (r *deleteTestRepo) OrderByUpdateAtRaw(raw string, bindings ...interface{}) *deleteTestRepo {
	r.orderByCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(raw)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *deleteTestRepo) Limit(limit int) *deleteTestRepo {
	r.sqlBuilder.WriteString(" LIMIT ?")
	r.bindings = append(r.bindings, limit)
	return r
}

func (r *insertTestRepo) ExecContext(ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx, r.sqlBuilder.String(), r.bindings...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *insertTestRepo) MustExecContext(ctx context.Context) int64 {
	res, err := r.ExecContext(ctx)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *insertTestRepo) Exec() (int64, error) {
	return r.ExecContext(context.Background())
}

func (r *insertTestRepo) MustExec() int64 {
	return r.MustExecContext(context.Background())
}

func (r *duplicateTestRepo) ExecContext(ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx, r.sqlBuilder.String(), r.bindings...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *duplicateTestRepo) MustExecContext(ctx context.Context) int64 {
	res, err := r.ExecContext(ctx)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *duplicateTestRepo) Exec() (int64, error) {
	return r.ExecContext(context.Background())
}

func (r *duplicateTestRepo) MustExec() int64 {
	return r.MustExecContext(context.Background())
}

func (r *updateTestRepo) ExecContext(ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx, r.sqlBuilder.String(), r.bindings...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *updateTestRepo) MustExecContext(ctx context.Context) int64 {
	res, err := r.ExecContext(ctx)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *updateTestRepo) Exec() (int64, error) {
	return r.ExecContext(context.Background())
}

func (r *updateTestRepo) MustExec() int64 {
	return r.MustExecContext(context.Background())
}

func (r *deleteTestRepo) ExecContext(ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx, r.sqlBuilder.String(), r.bindings...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (r *deleteTestRepo) MustExecContext(ctx context.Context) int64 {
	res, err := r.ExecContext(ctx)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *deleteTestRepo) Exec() (int64, error) {
	return r.ExecContext(context.Background())
}

func (r *deleteTestRepo) MustExec() int64 {
	return r.MustExecContext(context.Background())
}

// GetContext 获取单条数据
func (r *selectTestRepo) GetContext(ctx context.Context) (*Test, error) {
	row := r.db.QueryRowContext(ctx, r.sqlBuilder.String(), r.bindings...)
	m := &Test{}
	scanners := make([]interface{}, len(r.fields))
	for i, field := range r.fields {
		switch field {
		case "`id`":
			scanners[i] = &m.Id
		case "`create_at`":
			scanners[i] = &m.CreateAt
		case "`update_at`":
			scanners[i] = &m.UpdateAt
		}
	}
	return m, row.Scan(scanners...)
}

// MustGetContext 获取单条数据,如果返回 sql.ErrNoRows 错误则返回 nil, 其他错误则 panic
func (r *selectTestRepo) MustGetContext(ctx context.Context) *Test {
	res, err := r.GetContext(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return res
}

// MustGetOrFailContext 必须返回数据,即使返回 sql.ErrNoRows 也会 panic
func (r *selectTestRepo) MustGetOrFailContext(ctx context.Context) *Test {
	res, err := r.GetContext(ctx)
	if err != nil {
		panic(err)
	}
	return res
}

// Get 获取单条数据
func (r *selectTestRepo) Get() (*Test, error) {
	return r.GetContext(context.Background())
}

// MustGet 获取单条数据,如果返回 sql.ErrNoRows 错误则返回 nil, 其他错误则 panic
func (r *selectTestRepo) MustGet() *Test {
	return r.MustGetContext(context.Background())
}

// MustGetOrFail 必须返回数据,即使返回 sql.ErrNoRows 也会 panic
func (r *selectTestRepo) MustGetOrFail() *Test {
	return r.MustGetOrFailContext(context.Background())
}

func (r *selectTestRepo) SelectContext(ctx context.Context) ([]*Test, error) {
	var err error
	rows, err := r.db.QueryContext(ctx, r.sqlBuilder.String(), r.bindings...)
	if err != nil {
		return nil, err
	}
	res := make([]*Test, 0)
	scanners := make([]interface{}, len(r.fields))
	defer func() {
		cerr := rows.Close()
		if err == nil {
			err = cerr
		}
	}()
	for rows.Next() {
		m := &Test{}
		for i, field := range r.fields {
			switch field {
			case `id`:
				scanners[i] = &m.Id
			case `create_at`:
				scanners[i] = &m.CreateAt
			case `update_at`:
				scanners[i] = &m.UpdateAt
			}
		}
		if err := rows.Scan(scanners...); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}

func (r *selectTestRepo) MustSelectContext(ctx context.Context) []*Test {
	res, err := r.SelectContext(ctx)
	if err != nil {
		panic(err)
	}
	return res
}

func (r *selectTestRepo) Select() ([]*Test, error) {
	return r.SelectContext(context.Background())
}

func (r *selectTestRepo) MustSelect() []*Test {
	return r.MustSelectContext(context.Background())
}

func (r *insertTestRepo) Duplicate() *duplicateTestRepo {
	res := &duplicateTestRepo{
		db:       r.db,
		bindings: r.bindings,
		first:    true,
	}
	res.sqlBuilder.WriteString(r.sqlBuilder.String() + " ON DUPLICATE KEY UPDATE")
	return res
}

func (r *updateTestRepo) setCheck() {
	if r.first {
		return
	}

	r.sqlBuilder.WriteString(",")
	r.first = false
}

func (r *updateTestRepo) Set(test *Test) *updateTestRepo {
	id := test.Id
	if id != 0 {
		r.setCheck()
		r.sqlBuilder.WriteString(" `id`=?")
		r.bindings = append(r.bindings, id)
	}

	createAt := test.CreateAt
	if !createAt.IsZero() {
		r.setCheck()
		r.sqlBuilder.WriteString(" `create_at`=?")
		r.bindings = append(r.bindings, createAt)
	}

	updateAt := test.UpdateAt
	if !updateAt.IsZero() {
		r.setCheck()
		r.sqlBuilder.WriteString(" `update_at`=?")
		r.bindings = append(r.bindings, updateAt)
	}

	return r
}

func (r *updateTestRepo) SetIdEmpty() *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `id`=?")
	r.bindings = append(r.bindings, 0)
	return r
}

func (r *updateTestRepo) SetIdNil() *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `id`=NULL")
	return r
}

func (r *updateTestRepo) SetIdRaw(sql string, bindings ...interface{}) *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(sql)
	r.bindings = append(r.bindings, bindings...)
	return r
}
func (r *updateTestRepo) SetCreateAtEmpty() *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `create_at`=?")
	r.bindings = append(r.bindings, time.Time{})
	return r
}

func (r *updateTestRepo) SetCreateAtNil() *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `create_at`=NULL")
	return r
}

func (r *updateTestRepo) SetCreateAtRaw(sql string, bindings ...interface{}) *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(sql)
	r.bindings = append(r.bindings, bindings...)
	return r
}
func (r *updateTestRepo) SetUpdateAtEmpty() *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `update_at`=?")
	r.bindings = append(r.bindings, time.Time{})
	return r
}

func (r *updateTestRepo) SetUpdateAtNil() *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `update_at`=NULL")
	return r
}

func (r *updateTestRepo) SetUpdateAtRaw(sql string, bindings ...interface{}) *updateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(sql)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *duplicateTestRepo) setCheck() {
	if r.first {
		return
	}

	r.sqlBuilder.WriteString(",")
	r.first = false
}

func (r *duplicateTestRepo) Set(test *Test) *duplicateTestRepo {
	id := test.Id
	if id != 0 {
		r.setCheck()
		r.sqlBuilder.WriteString(" `id`=?")
		r.bindings = append(r.bindings, id)
	}

	createAt := test.CreateAt
	if !createAt.IsZero() {
		r.setCheck()
		r.sqlBuilder.WriteString(" `create_at`=?")
		r.bindings = append(r.bindings, createAt)
	}

	updateAt := test.UpdateAt
	if !updateAt.IsZero() {
		r.setCheck()
		r.sqlBuilder.WriteString(" `update_at`=?")
		r.bindings = append(r.bindings, updateAt)
	}

	return r
}

func (r *duplicateTestRepo) SetIdEmpty() *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `id`=?")
	r.bindings = append(r.bindings, 0)
	return r
}

func (r *duplicateTestRepo) SetIdNil() *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `id`=NULL")
	return r
}

func (r *duplicateTestRepo) SetIdRaw(sql string, bindings ...interface{}) *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `id` ")
	r.sqlBuilder.WriteString(sql)
	r.bindings = append(r.bindings, bindings...)
	return r
}
func (r *duplicateTestRepo) SetCreateAtEmpty() *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `create_at`=?")
	r.bindings = append(r.bindings, time.Time{})
	return r
}

func (r *duplicateTestRepo) SetCreateAtNil() *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `create_at`=NULL")
	return r
}

func (r *duplicateTestRepo) SetCreateAtRaw(sql string, bindings ...interface{}) *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `create_at` ")
	r.sqlBuilder.WriteString(sql)
	r.bindings = append(r.bindings, bindings...)
	return r
}
func (r *duplicateTestRepo) SetUpdateAtEmpty() *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `update_at`=?")
	r.bindings = append(r.bindings, time.Time{})
	return r
}

func (r *duplicateTestRepo) SetUpdateAtNil() *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `update_at`=NULL")
	return r
}

func (r *duplicateTestRepo) SetUpdateAtRaw(sql string, bindings ...interface{}) *duplicateTestRepo {
	r.setCheck()
	r.sqlBuilder.WriteString(" `update_at` ")
	r.sqlBuilder.WriteString(sql)
	r.bindings = append(r.bindings, bindings...)
	return r
}

func (r *selectTestRepo) SQL() (string, []interface{}) {
	return r.sqlBuilder.String(), r.bindings
}

func (r *insertTestRepo) SQL() (string, []interface{}) {
	return r.sqlBuilder.String(), r.bindings
}

func (r *duplicateTestRepo) SQL() (string, []interface{}) {
	return r.sqlBuilder.String(), r.bindings
}

func (r *updateTestRepo) SQL() (string, []interface{}) {
	return r.sqlBuilder.String(), r.bindings
}

func (r *deleteTestRepo) SQL() (string, []interface{}) {
	return r.sqlBuilder.String(), r.bindings
}
