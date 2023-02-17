package qdelete

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/logging"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner"
)

type Client interface {
	Session() *gocql.Session
	Config() config.Config
	Restart() error
	Debug() bool
	PrintFn() logging.DebugPrint
}

type Runner interface {
	QueryNone(string, []interface{}) error
}

// Query represents a Cassandra delete query. Execution should not bind any value
type Query struct {
	client Client
	runner Runner
	table  string
	where  []where.Stm
	args   []interface{}
}

// New create a new delete query instance by passing a cassandra session
func New(c Client) *Query {
	return &Query{
		client: c,
		runner: runner.New(c),
	}
}

// From set table where be data deleted
func (dq *Query) From(t string) *Query {
	dq.table = t
	return dq
}

// Where set where conditions that can be nested to delete data
func (dq *Query) Where(f string, op where.Operator, v interface{}) *Query {
	dq.where = append(dq.where, where.Stm{Field: f, Op: op, Value: v})
	dq.args = append(dq.args, v)
	return dq
}
