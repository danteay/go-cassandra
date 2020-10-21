package qb

import "github.com/gocql/gocql"

// CountQuery create new select count query
type CountQuery struct {
	session *gocql.Session
	table   string
	fields  columns
	where   []whereStm
}

func newCountQuery(s *gocql.Session) *CountQuery {
	return &CountQuery{session: s}
}

// Columns set columns to count
func (cq *CountQuery) Columns(f ...string) *CountQuery {
	cq.fields = f
	return cq
}

// From set table for select query
func (cq *CountQuery) From(t string) *CountQuery {
	cq.table = t
	return cq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (cq *CountQuery) Where(f string, op WhereOp) *CountQuery {
	cq.where = append(cq.where, whereStm{
		field: f, op: op,
	})

	return cq
}
