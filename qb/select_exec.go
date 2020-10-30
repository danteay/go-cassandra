package qb

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
)

// One return just one result on bind action
func (sq *SelectQuery) One() error {
	if sq.bind == nil {
		return errors.New("nil bind is not allowed, use None() functions instead One()")
	}

	if err := verifyBind(sq.bind, reflect.Struct); err != nil {
		return err
	}

	q := sq.build()

	var jsonRow string

	if err := sq.session.Query(q, sq.args...).Consistency(gocql.One).Scan(&jsonRow); err != nil {
		return err
	}

	ib := reflect.Indirect(reflect.ValueOf(sq.bind))

	bv := reflect.ValueOf(ib.Interface())
	bt := bv.Type()

	elem, err := bindRow([]byte(jsonRow), bt)
	if err != nil {
		return err
	}

	ib.Set(reflect.Indirect(elem))

	return nil
}

// All return all rows on bind action. Bind should be a slice of structs
func (sq *SelectQuery) All() error {
	if sq.bind == nil {
		return errors.New("nil bind is not allowed, use None() function instead All()")
	}

	if err := verifyBind(sq.bind, reflect.Slice); err != nil {
		return err
	}

	q := sq.build()

	var jsonRow string

	iter := sq.session.Query(q, sq.args...).Iter()
	defer func() { _ = iter.Close() }()

	ib := reflect.Indirect(reflect.ValueOf(sq.bind))

	bv := reflect.ValueOf(ib.Interface())
	bt := bv.Type().Elem()

	for iter.Scan(&jsonRow) {
		elem, err := bindRow([]byte(jsonRow), bt)
		if err != nil {
			return err
		}

		ib.Set(reflect.Append(ib, reflect.Indirect(elem)))
	}

	return nil
}

// None execute a select query without returning values
func (sq *SelectQuery) None() error {
	q := sq.build()

	if err := sq.session.Query(q, sq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (sq *SelectQuery) build() string {
	q := qb.Select(sq.table)

	if len(sq.fields) > 0 {
		q = q.Columns(sq.fields...)
	}

	if len(sq.where) > 0 {
		q = q.Where(buildWhere(sq.where)...)
	}

	if sq.limit > 0 {
		q = q.Limit(sq.limit)
	}

	if len(sq.orderBy) > 0 {
		for _, column := range sq.orderBy {
			var order qb.Order = qb.DESC

			switch sq.order {
			case Asc:
				order = qb.ASC
			case Desc:
				order = qb.ASC
			}

			q = q.OrderBy(column, order)
		}
	}

	if len(sq.groupBy) > 0 {
		q = q.GroupBy(sq.groupBy...)
	}

	query, _ := q.Json().ToCql()

	return strings.TrimSpace(query)
}
