package runner

import (
	"github.com/avast/retry-go/v4"

	"github.com/danteay/go-cassandra/errors"
)

func (r *Runner) QueryNone(query string, args []interface{}) error {
	execFn := func() error {
		if r.client.Session() == nil || r.client.Session().Closed() {
			return errors.ErrClosedConnection
		}

		return r.client.Session().
			Query(query, args...).
			Consistency(r.getConsistency()).
			Exec()
	}

	return retry.Do(execFn, r.getRetryOptions()...)
}
