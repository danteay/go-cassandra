package qinsert

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"
)

// Exec execute qinsert query with args
func (iq *Query) Exec() error {
	q := iq.build()

	if err := iq.session.Query(q, iq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (iq *Query) build() string {
	q := qb.Insert(iq.table)
	q.Columns(iq.fields...)

	query, _ := q.ToCql()

	return strings.TrimSpace(query)
}
