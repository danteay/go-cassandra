package qselect

import (
	"strings"

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
	Query(string, []interface{}, interface{}) error
	QueryOne(string, []interface{}, interface{}) error
}

// Query represents a cassandra select statement and his options
type Query struct {
	client         Client
	runner         Runner
	fields         []string
	args           []interface{}
	table          string
	bind           interface{}
	where          []where.Stm
	groupBy        []string
	orderBy        []string
	order          string
	limit          uint
	allowFiltering bool
}

// New create a new select query by passing a cassandra session and debug options
func New(c Client) *Query {
	return &Query{client: c, runner: runner.New(c)}
}

// Fields save query fields that should be used for select query
func (q *Query) Fields(f ...string) *Query {
	q.fields = f
	return q
}

// From set table for select query
func (q *Query) From(t string) *Query {
	q.table = t
	return q
}

// Where adds single where conditional. If more are needed, concatenate more calls to this functions
func (q *Query) Where(f string, op where.Operator, v interface{}) *Query {
	q.where = append(q.where, where.Stm{Field: f, Op: op})
	q.args = append(q.args, v)
	return q
}

// OrderBy adds `order by` selection statement fields
func (q *Query) OrderBy(ob []string, o string) *Query {
	q.orderBy = ob
	q.order = strings.ToUpper(o)
	return q
}

// GroupBy adds `group by` statement fields
func (q *Query) GroupBy(f ...string) *Query {
	q.orderBy = f
	return q
}

// Limit adds `limit` query statement
func (q *Query) Limit(l uint) *Query {
	q.limit = l
	return q
}

// AllowFiltering sets a ALLOW FILTERING clause on the query.
func (q *Query) AllowFiltering() *Query {
	q.allowFiltering = true
	return q
}

// Bind save data struct to bind result query data
func (q *Query) Bind(b interface{}) *Query {
	q.bind = b
	return q
}
