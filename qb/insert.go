package qb

import "github.com/gocql/gocql"

// InsertQuery represent a Cassandra insert query. Execution should not bind any value
type InsertQuery struct {
	session *gocql.Session
	table   string
	fields  columns
	args    []interface{}
}

func newInsertQuery(s *gocql.Session, f ...string) *InsertQuery {
	return &InsertQuery{session: s, fields: f}
}

// Into set table to insert query
func (iq *InsertQuery) Into(t string) *InsertQuery {
	iq.table = t
	return iq
}

// Values set values as query arguments for insert statement
func (iq *InsertQuery) Values(v ...interface{}) *InsertQuery {
	iq.args = v
	return iq
}
