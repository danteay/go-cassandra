package qb

import (
	"github.com/danteay/go-cassandra/qb/types"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner"
)

type Query struct {
	runner         *runner.Runner
	queryType      types.QueryType
	table          string
	countColumn    string
	distinct       bool
	fields         []string
	set            []string
	where          []where.Stm
	allowFiltering bool
	args           []interface{}
	bind           interface{}
	groupBy        []string
	orderBy        []string
	order          string
	limit          uint
}

// NewQuery returns a New Query pointer instance
func NewQuery(qt types.QueryType, c runner.Client) *Query {
	return &Query{
		queryType: qt,
		runner:    runner.New(c),
		args:      make([]interface{}, 0),
	}
}

// Select create a Query for a SELECT statement.
func Select(c runner.Client) *Query {
	return NewQuery(types.Select, c)
}

func Count(c runner.Client) *Query {
	return NewQuery(types.Count, c)
}

func Insert(c runner.Client) *Query {
	return NewQuery(types.Insert, c)
}

func Update(c runner.Client) *Query {
	return NewQuery(types.Update, c)
}

func Delete(c runner.Client) *Query {
	return NewQuery(types.Delete, c)
}
