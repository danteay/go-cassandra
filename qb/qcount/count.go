package qcount

import (
	"github.com/danteay/go-cassandra/qb/query"
	"github.com/gocql/gocql"
)

// Query create new qselect qcount query
type Query struct {
	session *gocql.Session
	table   string
	column  string
	where   []query.WhereStm
	args    []interface{}
}

// New create a new qcount query instnace by passing a cassandra session
func New(s *gocql.Session) *Query {
	return &Query{session: s}
}

// Column set qcount column of the query
func (cq *Query) Column(c string) *Query {
	cq.column = c
	return cq
}

// From set table for qselect query
func (cq *Query) From(t string) *Query {
	cq.table = t
	return cq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (cq *Query) Where(f string, op query.WhereOp, v interface{}) *Query {
	cq.where = append(cq.where, query.WhereStm{Field: f, Op: op})
	cq.args = append(cq.args, v)
	return cq
}
