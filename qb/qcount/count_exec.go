package qcount

import (
	"strings"

	"github.com/danteay/go-cassandra/qb/query"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/qb"
)

// Exec release count query an return the number of rows and a possible error
func (cq *Query) Exec() (int64, error) {
	q := cq.build()

	var count int64

	if err := cq.ctx.Session.Query(q, cq.args...).Consistency(gocql.One).Scan(&count); err != nil {
		return 0, err
	}

	return 0, nil
}

func (cq *Query) build() string {
	q := qb.Select(cq.table).Count(cq.column)

	if len(cq.where) > 0 {
		q = q.Where(query.BuildWhere(cq.where)...)
	}

	queryStr, _ := q.ToCql()

	if cq.ctx.Debug {
		cq.ctx.PrintQuery(queryStr, cq.args)
	}

	return strings.TrimSpace(queryStr)
}
