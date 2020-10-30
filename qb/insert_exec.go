package qb

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"
)

// Exec execute insert query with args
func (iq *InsertQuery) Exec() error {
	q := iq.build()

	if err := iq.session.Query(q, iq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (iq *InsertQuery) build() string {
	q := qb.Insert(iq.table)
	q.Columns(iq.fields...)

	query, _ := q.ToCql()

	return strings.TrimSpace(query)
}
