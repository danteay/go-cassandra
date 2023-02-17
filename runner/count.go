package runner

import (
	"github.com/avast/retry-go/v4"

	"github.com/danteay/go-cassandra/errors"
)

func (r *Runner) Count(query string, args []interface{}) (int64, error) {
	var count int64

	execFn := func() error {
		if r.client.Session() == nil || r.client.Session().Closed() {
			return errors.ErrClosedConnection
		}

		return r.client.Session().Query(query, args...).Consistency(r.getConsistency()).Scan(&count)
	}

	if err := retry.Do(execFn, r.getRetryOptions()...); err != nil {
		return 0, err
	}

	return count, nil
}
