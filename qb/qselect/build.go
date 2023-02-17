package qselect

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/constants"
	"github.com/danteay/go-cassandra/qb/where"
)

func (q *Query) build() (string, error) {
	sb := qb.Select(q.table)

	if len(q.fields) > 0 {
		sb = sb.Columns(q.fields...)
	}

	if len(q.where) > 0 {
		stms, err := where.BuildStms(q.where)
		if err != nil {
			return "", err
		}

		sb = sb.Where(stms...)
	}

	if q.limit > 0 {
		sb = sb.Limit(q.limit)
	}

	if q.allowFiltering {
		sb = sb.AllowFiltering()
	}

	if len(q.orderBy) > 0 {
		for _, column := range q.orderBy {
			var order qb.Order = qb.DESC

			switch q.order {
			case constants.Asc:
				order = qb.ASC
			case constants.Desc:
				order = qb.ASC
			}

			sb = sb.OrderBy(column, order)
		}
	}

	if len(q.groupBy) > 0 {
		sb = sb.GroupBy(q.groupBy...)
	}

	queryStr, _ := sb.ToCql()

	if q.client.Debug() {
		q.client.PrintFn()(queryStr, q.args, nil)
	}

	return strings.TrimSpace(queryStr), nil
}
