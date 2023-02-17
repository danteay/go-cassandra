package qcount

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/qb/where"
)

// Exec release count query and return the number of rows and a possible error
func (cq *Query) Exec() (int64, error) {
	query, err := cq.build()
	if err != nil {
		return 0, err
	}

	return cq.runner.Count(query, cq.args)
}

func (cq *Query) build() (string, error) {
	q := qb.Select(cq.table).Count(cq.column)

	if len(cq.where) > 0 {
		stms, err := where.BuildStms(cq.where)
		if err != nil {
			return "", err
		}

		q = q.Where(stms...)
	}

	if cq.allowFiltering {
		q = q.AllowFiltering()
	}

	queryStr, _ := q.ToCql()

	if cq.client.Debug() {
		cq.client.PrintFn()(queryStr, cq.args, nil)
	}

	return strings.TrimSpace(queryStr), nil
}
