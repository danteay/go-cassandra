package qinsert

import (
	"github.com/danteay/go-cassandra/qb/query"
	"github.com/gocql/gocql"
)

// Query represent a Cassandra qinsert query. Execution should not bind any value
type Query struct {
	session *gocql.Session
	table   string
	fields  query.Columns
	args    []interface{}
}

// New creates a new insert query by passing a cassandra session and a variadic bunch of affected fields
func New(s *gocql.Session, f ...string) *Query {
	return &Query{session: s, fields: f}
}

// Into set table to qinsert query
func (iq *Query) Into(t string) *Query {
	iq.table = t
	return iq
}

// Values set values as query arguments for qinsert statement
func (iq *Query) Values(v ...interface{}) *Query {
	iq.args = v
	return iq
}
