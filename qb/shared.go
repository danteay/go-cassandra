package qb

import "github.com/danteay/go-cassandra/qb/where"

// From set table for count query
func (q *Query) From(t string) *Query {
	q.table = t
	return q
}

// AllowFiltering sets a ALLOW FILTERING clause on the query.
func (q *Query) AllowFiltering() *Query {
	q.allowFiltering = true
	return q
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (q *Query) Where(f string, op where.Operator, v interface{}) *Query {
	q.where = append(q.where, where.Stm{Field: f, Op: op, Value: v})
	q.args = append(q.args, v)
	return q
}
