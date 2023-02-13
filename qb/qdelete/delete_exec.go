package qdelete

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/qb/query"
)

// Exec execute delete query and return error on failure
func (dq *Query) Exec() error {
	q := dq.build()

	if err := dq.client.Session().Query(q, dq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (dq *Query) build() string {
	q := qb.Delete(dq.table)

	if len(dq.where) > 0 {
		q = q.Where(query.BuildWhere(dq.where)...)
	}

	queryStr, _ := q.ToCql()

	if dq.client.Debug() {
		dq.client.PrintFn()(queryStr, dq.args, nil)
	}

	return strings.TrimSpace(queryStr)
}
