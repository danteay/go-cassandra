package qselect

import (
	"github.com/danteay/go-cassandra/qb/query"
	"github.com/gocql/gocql"
)

// Query represents a cassandra select statement and his options
type Query struct {
	session *gocql.Session
	fields  query.Columns
	args    []interface{}
	table   string
	bind    interface{}
	where   []query.WhereStm
	groupBy query.Columns
	orderBy query.Columns
	order   query.Order
	limit   uint
}

// New create a new select query by passing a cassandra session and a variadic bunch of fields that should be
// returned on the binded result
func New(s *gocql.Session, f ...string) *Query {
	return &Query{session: s, fields: f}
}

// From set table for select query
func (q *Query) From(t string) *Query {
	q.table = t
	return q
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (q *Query) Where(f string, op query.WhereOp, v interface{}) *Query {
	q.where = append(q.where, query.WhereStm{Field: f, Op: op})
	q.args = append(q.args, v)
	return q
}

// OrderBy adds `order by` selection statement fields
func (q *Query) OrderBy(ob []string, o query.Order) *Query {
	q.orderBy = ob
	q.order = o
	return q
}

// GroupBy adds `group by` statement fields
func (q *Query) GroupBy(f ...string) *Query {
	q.orderBy = f
	return q
}

// Limit adds `limit` query statement
func (q *Query) Limit(l uint) *Query {
	q.limit = l
	return q
}

// Bind save data struct to bind result query data
func (q *Query) Bind(b interface{}) *Query {
	q.bind = b
	return q
}
