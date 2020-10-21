package qb

import (
	"github.com/gocql/gocql"
)

// SelectQuery represents a cassandra select statement and his options
type SelectQuery struct {
	session *gocql.Session
	fields  columns
	args    []interface{}
	table   string
	bind    interface{}
	where   []whereStm
	groupBy columns
	orderBy columns
	order   Order
	limit   uint
}

func newSelectQuery(s *gocql.Session, f ...string) *SelectQuery {
	return &SelectQuery{session: s, fields: f}
}

// From set table for select query
func (sq *SelectQuery) From(t string) *SelectQuery {
	sq.table = t
	return sq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (sq *SelectQuery) Where(f string, op WhereOp) *SelectQuery {
	sq.where = append(sq.where, whereStm{
		field: f, op: op,
	})

	return sq
}

// OrderBy adds `order by` selection statement fields
func (sq *SelectQuery) OrderBy(ob []string, o Order) *SelectQuery {
	sq.orderBy = ob
	sq.order = o
	return sq
}

// GroupBy adds `group by` statement fields
func (sq *SelectQuery) GroupBy(f ...string) *SelectQuery {
	sq.orderBy = f
	return sq
}

// Limit adds `limit` query statement
func (sq *SelectQuery) Limit(l uint) *SelectQuery {
	sq.limit = l
	return sq
}

// Args set replacement arg values on result query to be processed
func (sq *SelectQuery) Args(args ...interface{}) *SelectQuery {
	sq.args = args
	return sq
}

// Bind save data struct to bind result query data
func (sq *SelectQuery) Bind(b interface{}) *SelectQuery {
	sq.bind = b
	return sq
}
