package binder

import (
	"fmt"
	"reflect"
	"time"

	"github.com/danteay/go-cassandra/constants"
	"github.com/danteay/go-cassandra/errors"
)

// MapToStruct bind values of a map into a new struct of the given reflected type Value
func (b *Binder) MapToStruct(m map[string]interface{}, st reflect.Type) (reflect.Value, error) {
	elem := reflect.New(st)

	indVal := reflect.Indirect(elem)
	indTyp := indVal.Type()

	if st.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.ErrNoPtrBinding
	}

	if indVal.Kind() != reflect.Struct {
		return reflect.Value{}, errors.ErrNoSliceOfStructsBinding
	}

	numField := indTyp.NumField()

	for i := 0; i < numField; i++ {
		field := indTyp.Field(i)
		value := indVal.Field(i)

		tagField := field.Tag.Get(constants.Tag)
		mv, ok := m[tagField]

		if tagField != "" && ok {
			if err := b.castMapValue(mv, field.Type, value); err != nil {
				return reflect.Value{}, err
			}
		}
	}

	return elem, nil
}

// castMapValue convert interface value into a compatible value of the reflected type
func (b *Binder) castMapValue(mv interface{}, t reflect.Type, v reflect.Value) error {
	newVal := reflect.Indirect(reflect.New(t)).Interface()

	switch newVal.(type) {
	case int64:
		intVal := int64(mv.(float64))
		v.SetInt(intVal)
	case float64:
		floatVal, _ := mv.(float64)
		v.SetFloat(floatVal)
	case string:
		stringVal, _ := mv.(string)
		v.SetString(stringVal)
	case bool:
		boolVal, _ := mv.(bool)
		v.SetBool(boolVal)
	case time.Time:
		t, _ := time.Parse(b.client.Config().DatetimeLayout, mv.(string))
		v.Set(reflect.ValueOf(t))
	default:
		return fmt.Errorf("go-cassandra: can't cast value of type %T with value %v, to type %v", mv, mv, t)
	}

	return nil
}
