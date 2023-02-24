package runner

import (
	"reflect"

	"github.com/avast/retry-go/v4"
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/binder"
	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/logging"
)

//go:generate mockery --name=Client --filename=client.go --structname=Client --output=mocks --outpkg=mocks
//go:generate mockery --name=Binder --filename=binder.go --structname=Binder --output=mocks --outpkg=mocks

type Client interface {
	Session() *gocql.Session
	Config() config.Config
	Restart() error
	Debug() bool
	PrintFn() logging.DebugPrint
}

type Binder interface {
	MapToStruct(map[string]interface{}, reflect.Type) (reflect.Value, error)
	IsBindable(interface{}, reflect.Kind) error
}

type Runner struct {
	client Client
	binder Binder
}

// New creates a new Runner instance from a Client interface
func New(c Client) *Runner {
	return &Runner{client: c, binder: binder.New(c)}
}

func (r *Runner) getRetryOptions() []retry.Option {
	return []retry.Option{
		retry.Attempts(uint(r.client.Config().NoHostRetries)),
		retry.RetryIf(func(err error) bool {
			return err == gocql.ErrNoConnections
		}),
		retry.OnRetry(func(n uint, err error) {
			_ = r.client.Restart()
			// TODO: add logging
		}),
	}
}

func (r *Runner) getConsistency() gocql.Consistency {
	return gocql.Consistency(r.client.Config().Consistency)
}
