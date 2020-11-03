package qupdate

import (
	"github.com/danteay/go-cassandra/qb/query"
	"github.com/gocql/gocql"
)

// Query represent a Cassandra qupdate query. Execution should not bind any value
type Query struct {
	session *gocql.Session
	table   string
	fields  query.Columns
	args    []interface{}
	where   []query.WhereStm
}

// New create a new qupdate query by passing a cassandra session and the affected table
func New(s *gocql.Session, t string) *Query {
	return &Query{session: s, table: t}
}

// Set save field and corresponding value to qupdate
func (uq *Query) Set(f string, v interface{}) *Query {
	uq.fields = append(uq.fields, f)
	uq.args = append(uq.args, v)
	return uq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (uq *Query) Where(f string, op query.WhereOp, v interface{}) *Query {
	uq.where = append(uq.where, query.WhereStm{Field: f, Op: op, Value: v})
	uq.args = append(uq.args, v)
	return uq
}
