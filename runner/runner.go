package runner

import (
	"github.com/avast/retry-go/v4"
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra"
	"github.com/danteay/go-cassandra/errors"
	"github.com/danteay/go-cassandra/qb/query"
)

type Client interface {
	Session() *gocql.Session
	Config() gocassandra.Config
	Restart() error
	Debug() bool
	PrintFn() query.DebugPrint
}

type Runner struct {
	client Client
}

func (r *Runner) Query() {}

func (r *Runner) QueryOne() {}

func (r *Runner) QueryNone(query string, args []interface{}) error {
	execFn := func() error {
		if r.client.Session() == nil || r.client.Session().Closed() {
			return errors.ErrClosedConnection
		}

		return r.client.Session().Query(query, args...).Exec()
	}

	opts := []retry.Option{
		retry.Attempts(uint(r.client.Config().NoHostRetries)),
		retry.RetryIf(func(err error) bool {
			switch err {
			case gocql.ErrNoConnections:
				return true
			default:
				return false
			}
		}),
		retry.OnRetry(func(n uint, err error) {
			errRestart := r.client.Restart()

			if r.client.Debug() {
				r.client.PrintFn()("", nil, errRestart)
			}
		}),
	}

	return retry.Do(execFn, opts...)
}

func New(c Client) *Runner {
	return &Runner{client: c}
}
