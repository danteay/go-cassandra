package qb

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"
)

func (uq *UpdateQuery) Exec() error {
	q := uq.build()

	if err := uq.session.Query(q, uq.args...).Exec(); err != nil {
		return err
	}

	return nil
}

func (uq *UpdateQuery) build() string {
	q := qb.Update(uq.table)

	if len(uq.fields) > 0 {
		q = q.Set(uq.fields...)
	}

	if len(uq.where) > 0 {
		if len(uq.where) > 0 {
			q = q.Where(buildWhere(uq.where)...)
		}
	}

	query, _ := q.ToCql()

	return strings.TrimSpace(query)
}
