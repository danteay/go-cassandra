package qb

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"
)

// Exec execute delete query and return error on failure
func (dq *DeleteQuery) Exec() error {
	q := dq.build()

	if err := dq.session.Query(q, dq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (dq *DeleteQuery) build() string {
	q := qb.Delete(dq.table)

	if len(dq.where) > 0 {
		q = q.Where(buildWhere(dq.where)...)
	}

	query, _ := q.ToCql()

	return strings.TrimSpace(query)
}
