package qdelete

import (
	"github.com/danteay/go-cassandra/qb/query"
	"github.com/gocql/gocql"
)

// Query represents a Cassandra qdelete query. Execution should not bind any value
type Query struct {
	session *gocql.Session
	table   string
	where   []query.WhereStm
	args    []interface{}
}

// New create a new qdelete query instance by passing a cassandra session
func New(s *gocql.Session) *Query {
	return &Query{session: s}
}

// From set table where be data deleted
func (dq *Query) From(t string) *Query {
	dq.table = t
	return dq
}

// Where set where conditions that can be nested to qdelete data
func (dq *Query) Where(f string, op query.WhereOp, v interface{}) *Query {
	dq.where = append(dq.where, query.WhereStm{Field: f, Op: op})
	dq.args = append(dq.args, v)
	return dq
}
