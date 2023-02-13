package qupdate

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/qb/query"
)

//go:generate mockery --name=Client --filename=client.go --structname=Client --output=mocks --outpkg=mocks

type Client interface {
	Session() *gocql.Session
	Debug() bool
	Restart() error
	PrintFn() query.DebugPrint
}

// Query represent a Cassandra update query. Execution should not bind any value
type Query struct {
	client Client
	table  string
	fields query.Columns
	args   []interface{}
	where  []query.WhereStm
}

// New create a new update query by passing a cassandra session and the affected table
func New(c Client) *Query {
	return &Query{client: c}
}

// Table set the table name to affect with the update query
func (uq *Query) Table(t string) *Query {
	uq.table = t
	return uq
}

// Set save field and corresponding value to update
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
