package qb

import "github.com/gocql/gocql"

type CountQuery struct {
	session *gocql.Session
	table   string
	where   []whereStm
}

func newCountQuery(s *gocql.Session) *CountQuery {
	return &CountQuery{session: s}
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
