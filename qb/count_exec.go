package qb

import (
	"strings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
)

// Exec release count query an return the number of rows and a possible error
func (cq *CountQuery) Exec() (int64, error) {
	q := cq.build()

	var count int64

	if err := cq.session.Query(q, cq.args...).Consistency(gocql.One).Scan(&count); err != nil {
		return 0, err
	}

	return 0, nil
}

func (cq *CountQuery) build() string {
	q := qb.Select(cq.table)

	if len(cq.where) > 0 {
		q = q.Where(buildWhere(cq.where)...)
	}

	query, _ := q.ToCql()

	return strings.TrimSpace(query)
}
