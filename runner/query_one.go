package runner

import (
	"reflect"

	"github.com/avast/retry-go/v4"

	"github.com/danteay/go-cassandra/errors"
)

func (r *Runner) QueryOne(query string, args []interface{}, bind interface{}) error {
	if bind == nil {
		return errors.ErrNilBinding
	}

	row := make(map[string]interface{})

	if err := r.binder.IsBindable(bind, reflect.Struct); err != nil {
		return err
	}

	execFn := func() error {
		if r.client.Session() == nil || r.client.Session().Closed() {
			return errors.ErrClosedConnection
		}

		return r.client.Session().
			Query(query, args...).
			Consistency(r.getConsistency()).
			MapScan(row)
	}

	if err := retry.Do(execFn, r.getRetryOptions()...); err != nil {
		return err
	}

	ib := reflect.Indirect(reflect.ValueOf(bind))

	bv := reflect.ValueOf(ib.Interface())
	bt := bv.Type()

	elem, err := r.binder.MapToStruct(row, bt)
	if err != nil {
		return err
	}

	ib.Set(reflect.Indirect(elem))

	return nil
}
