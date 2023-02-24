package qb

import (
	"strings"

	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/constants"
	"github.com/danteay/go-cassandra/errors"
	"github.com/danteay/go-cassandra/qb/types"
	"github.com/danteay/go-cassandra/qb/where"
)

// ToCQL builds query to create the configured query type, returns his arguments or an error if something goes wrong.
func (q *Query) ToCQL() (string, []interface{}, error) {
	var sql string
	var err error

	switch q.queryType {
	case types.Count:
		sql, err = q.buildCount()
	case types.Delete:
		sql, err = q.buildDelete()
	case types.Insert:
		sql, err = q.buildInsert()
	case types.Update:
		sql, err = q.buildUpdate()
	case types.Select:
		sql, err = q.buildSelect()
	default:
		err = errors.ErrInvalidQueryType
	}

	if err != nil {
		return "", []interface{}{}, err
	}

	return sql, q.args, nil
}

func (q *Query) buildCount() (string, error) {
	if q.countColumn == "" {
		q.countColumn = "*"
	}

	cql := qb.Select(q.table).Count(q.countColumn)

	if len(q.where) > 0 {
		stms, err := where.BuildStms(q.where)
		if err != nil {
			return "", err
		}

		cql = cql.Where(stms...)
	}

	if q.allowFiltering {
		cql = cql.AllowFiltering()
	}

	queryStr, _ := cql.ToCql()

	return strings.TrimSpace(queryStr), nil
}

func (q *Query) buildDelete() (string, error) {
	cql := qb.Delete(q.table)

	if len(q.where) > 0 {
		stms, err := where.BuildStms(q.where)
		if err != nil {
			return "", err
		}

		cql = cql.Where(stms...)
	}

	queryStr, _ := cql.ToCql()

	return strings.TrimSpace(queryStr), nil
}

func (q *Query) buildInsert() (string, error) {
	cql := qb.Insert(q.table)
	cql.Columns(q.fields...)

	queryStr, _ := cql.ToCql()

	return strings.TrimSpace(queryStr), nil
}

func (q *Query) buildUpdate() (string, error) {
	cql := qb.Update(q.table)

	if len(q.set) > 0 {
		cql = cql.Set(q.set...)
	}

	if len(q.where) > 0 {
		stms, err := where.BuildStms(q.where)
		if err != nil {
			return "", err
		}

		cql = cql.Where(stms...)
	}

	queryStr, _ := cql.ToCql()

	return strings.TrimSpace(queryStr), nil
}

func (q *Query) buildSelect() (string, error) {
	sb := qb.Select(q.table)

	if len(q.fields) > 0 {
		if q.distinct {
			sb = sb.Distinct(q.fields...)
		} else {
			sb = sb.Columns(q.fields...)
		}
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

	return strings.TrimSpace(queryStr), nil
}
