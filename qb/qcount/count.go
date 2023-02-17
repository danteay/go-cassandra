package qcount

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/logging"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner"
)

//go:generate mockery --name=Client --filename=client.go --structname=Client --output=mocks --outpkg=mocks
//go:generate mockery --name=Runner --filename=runner.go --structname=Runner --output=mocks --outpkg=mocks

type Client interface {
	Session() *gocql.Session
	Config() config.Config
	Restart() error
	Debug() bool
	PrintFn() logging.DebugPrint
}

type Runner interface {
	Count(string, []interface{}) (int64, error)
}

// Query create new select count query
type Query struct {
	client         Client
	runner         Runner
	table          string
	column         string
	where          []where.Stm
	allowFiltering bool
	args           []interface{}
}

// New create a new count query instance by passing a cassandra session
func New(c Client) *Query {
	return &Query{
		client: c,
		runner: runner.New(c),
	}
}

// Column set count column of the query
func (cq *Query) Column(c string) *Query {
	cq.column = c
	return cq
}

// From set table for count query
func (cq *Query) From(t string) *Query {
	cq.table = t
	return cq
}

// AllowFiltering sets a ALLOW FILTERING clause on the query.
func (cq *Query) AllowFiltering() *Query {
	cq.allowFiltering = true
	return cq
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (cq *Query) Where(f string, op where.Operator, v interface{}) *Query {
	cq.where = append(cq.where, where.Stm{Field: f, Op: op, Value: v})
	cq.args = append(cq.args, v)
	return cq
}
