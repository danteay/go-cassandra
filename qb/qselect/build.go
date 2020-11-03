package qselect

import (
	"strings"

	"github.com/danteay/go-cassandra/qb/query"
	"github.com/scylladb/gocqlx/qb"
)

func (q *Query) build() string {
	sb := qb.Select(q.table)

	if len(q.fields) > 0 {
		sb = sb.Columns(q.fields...)
	}

	if len(q.where) > 0 {
		sb = sb.Where(query.BuildWhere(q.where)...)
	}

	if q.limit > 0 {
		sb = sb.Limit(q.limit)
	}

	if len(q.orderBy) > 0 {
		for _, column := range q.orderBy {
			var order qb.Order = qb.DESC

			switch q.order {
			case query.Asc:
				order = qb.ASC
			case query.Desc:
				order = qb.ASC
			}

			sb = sb.OrderBy(column, order)
		}
	}

	if len(q.groupBy) > 0 {
		sb = sb.GroupBy(q.groupBy...)
	}

	queryStr, _ := sb.Json().ToCql()

	return strings.TrimSpace(queryStr)
}
