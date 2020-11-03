package qdelete

import (
	"strings"

	"github.com/danteay/go-cassandra/qb/query"
	"github.com/scylladb/gocqlx/qb"
)

// Exec execute qdelete query and return error on failure
func (dq *Query) Exec() error {
	q := dq.build()

	if err := dq.session.Query(q, dq.args...).Exec(); err != nil {
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

	return strings.TrimSpace(queryStr)
}
