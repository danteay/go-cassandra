package qb

import "github.com/gocql/gocql"

// UpdateQuery represent a Cassandra update query. Execution should not bind any value
type UpdateQuery struct {
	session *gocql.Session
	table   string
	fields  columns
	args    []interface{}
	where   []whereStm
}

func newUpdateQuery(s *gocql.Session, t string) *UpdateQuery {
	return &UpdateQuery{session: s, table: t}
}

// Set save field and corresponding value to update
func (uq *UpdateQuery) Set(f string, v interface{}) *UpdateQuery {
	uq.fields = append(uq.fields, f)
	uq.args = append(uq.args, v)
	return uq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (uq *UpdateQuery) Where(f string, op WhereOp, v interface{}) *UpdateQuery {
	uq.where = append(uq.where, whereStm{field: f, op: op, value: v})
	return uq
}
