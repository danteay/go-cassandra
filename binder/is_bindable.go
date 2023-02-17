package binder

import (
	"reflect"

	"github.com/danteay/go-cassandra/errors"
)

// IsBindable verify if an interface is bindable or not by checking it is a Ptr kind
func (b *Binder) IsBindable(target interface{}, kind reflect.Kind) error {
	t := reflect.TypeOf(target)
	v := reflect.ValueOf(target)

	if t.Kind() != reflect.Ptr {
		return errors.ErrNoPtrBinding
	}

	str := reflect.Indirect(v).Interface()
	t = reflect.TypeOf(str)

	if t.Kind() != kind {
		return errors.ErrNoStructOrSliceBinding
	}

	return nil
}
