package qupdate

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/qb/where"
)

// Exec run update query from builder and return an error if exists
func (uq *Query) Exec() error {
	query, err := uq.build()
	if err != nil {
		return err
	}

	return uq.runner.QueryNone(query, uq.args)
}

func (uq *Query) build() (string, error) {
	q := qb.Update(uq.table)

	if len(uq.fields) > 0 {
		q = q.Set(uq.fields...)
	}

	if len(uq.where) > 0 {
		stms, err := where.BuildStms(uq.where)
		if err != nil {
			return "", err
		}
		q = q.Where(stms...)
	}

	queryStr, _ := q.ToCql()

	if uq.client.Debug() {
		uq.client.PrintFn()(queryStr, uq.args, nil)
	}

	return strings.TrimSpace(queryStr), nil
}
