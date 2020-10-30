package qb

import "github.com/gocql/gocql"

// CountQuery create new select count query
type CountQuery struct {
	session *gocql.Session
	table   string
	column  string
	where   []whereStm
	args    []interface{}
}

func newCountQuery(s *gocql.Session) *CountQuery {
	return &CountQuery{session: s}
}

// Column set count column of the query
func (cq *CountQuery) Column(c string) *CountQuery {
	cq.column = c
	return cq
}

// From set table for select query
func (cq *CountQuery) From(t string) *CountQuery {
	cq.table = t
	return cq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (cq *CountQuery) Where(f string, op WhereOp, v interface{}) *CountQuery {
	cq.where = append(cq.where, whereStm{field: f, op: op})
	cq.args = append(cq.args, v)
	return cq
}
