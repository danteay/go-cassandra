package runner

import (
	"reflect"

	"github.com/avast/retry-go/v4"

	"github.com/danteay/go-cassandra/errors"
)

func (r *Runner) Query(query string, args []interface{}, bind interface{}) error {
	if bind == nil {
		return errors.ErrNilBinding
	}

	rows := make([]map[string]interface{}, 0)

	if err := r.binder.IsBindable(bind, reflect.Slice); err != nil {
		return err
	}

	execFn := func() error {
		if r.client.Session() == nil || r.client.Session().Closed() {
			return errors.ErrClosedConnection
		}

		iter := r.client.Session().Query(query, args...).Consistency(r.getConsistency()).Iter()

		row := make(map[string]interface{})
		for iter.MapScan(row) {
			rows = append(rows, row)
		}

		return iter.Close()
	}

	if err := retry.Do(execFn, r.getRetryOptions()...); err != nil {
		return err
	}

	ib := reflect.Indirect(reflect.ValueOf(bind))

	bv := reflect.ValueOf(ib.Interface())
	bt := bv.Type().Elem()

	for _, row := range rows {
		elem, err := r.binder.MapToStruct(row, bt)
		if err != nil {
			return err
		}

		ib.Set(reflect.Append(ib, reflect.Indirect(elem)))
	}

	return nil
}
