package qupdate

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/logging"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner"
)

//go:generate mockery --name=Client --filename=client.go --structname=Client --output=mocks --outpkg=mocks

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

// Query represent a Cassandra update query. Execution should not bind any value
type Query struct {
	client Client
	runner Runner
	table  string
	fields []string
	args   []interface{}
	where  []where.Stm
}

// New create a new update query by passing a cassandra session and the affected table
func New(c Client) *Query {
	return &Query{client: c, runner: runner.New(c)}
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
func (uq *Query) Where(f string, op where.Operator, v interface{}) *Query {
	uq.where = append(uq.where, where.Stm{Field: f, Op: op, Value: v})
	uq.args = append(uq.args, v)
	return uq
}
