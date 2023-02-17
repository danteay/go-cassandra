package qdelete

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/qb/where"
)

// Exec execute delete query and return error on failure
func (dq *Query) Exec() error {
	q, err := dq.build()
	if err != nil {
		return err
	}

	return dq.runner.QueryNone(q, dq.args)
}

func (dq *Query) build() (string, error) {
	q := qb.Delete(dq.table)

	if len(dq.where) > 0 {
		stms, err := where.BuildStms(dq.where)
		if err != nil {
			return "", err
		}

		q = q.Where(stms...)
	}

	queryStr, _ := q.ToCql()

	if dq.client.Debug() {
		dq.client.PrintFn()(queryStr, dq.args, nil)
	}

	return strings.TrimSpace(queryStr), nil
}
