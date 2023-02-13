package qinsert

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"
)

// Exec execute insert query with args
func (iq *Query) Exec() error {
	q := iq.build()

	if err := iq.client.Session().Query(q, iq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (iq *Query) build() string {
	q := qb.Insert(iq.table)
	q.Columns(iq.fields...)

	queryStr, _ := q.ToCql()

	if iq.client.Debug() {
		iq.client.PrintFn()(queryStr, iq.args, nil)
	}

	return strings.TrimSpace(queryStr)
}
