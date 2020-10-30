package qb

import "github.com/gocql/gocql"

// DeleteQuery represents a Cassandra delete query. Execution should not bind any value
type DeleteQuery struct {
	session *gocql.Session
	table   string
	where   []whereStm
	args    []interface{}
}

func newDeleteQuery(s *gocql.Session) *DeleteQuery {
	return &DeleteQuery{session: s}
}

// From set table where be data deleted
func (dq *DeleteQuery) From(t string) *DeleteQuery {
	dq.table = t
	return dq
}

// Where set where conditions that can be nested to delete data
func (dq *DeleteQuery) Where(f string, op WhereOp, v interface{}) *DeleteQuery {
	dq.where = append(dq.where, whereStm{field: f, op: op})
	dq.args = append(dq.args, v)
	return dq
}
