package qb

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func verifyBind(b interface{}, k reflect.Kind) error {
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)

	if t.Kind() != reflect.Ptr {
		return errors.New("bind value should be a pointer")
	}

	str := reflect.Indirect(v).Interface()
	t = reflect.TypeOf(str)

	if t.Kind() != k {
		return errors.New("bind value needs to be a struct to get one value, or a slice of structs to get many")
	}

	return nil
}

func bindMapToStruct(m map[string]interface{}, st reflect.Value) error {
	indVal := reflect.Indirect(st)
	indTyp := indVal.Type()

	if st.Kind() != reflect.Ptr {
		return errors.New("bind should be a slice of struct pointers with `gocql` tags")
	}

	if indVal.Kind() != reflect.Struct {
		return errors.New("bind should be a slice of struct pointers with `gocql` tags - 2")
	}

	numField := indTyp.NumField()

	for i := 0; i < numField; i++ {
		field := indTyp.Field(i)
		value := indVal.Field(i)

		tagField := field.Tag.Get(modelsTag)
		mv, ok := m[tagField]

		if tagField != "" && ok {
			fmt.Printf("map: %T field: %v\n", mv, field.Type)

			err := castMapValue(mv, field.Type, value)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}

	return nil
}

func castMapValue(mv interface{}, t reflect.Type, v reflect.Value) error {
	newVal := reflect.Indirect(reflect.New(t)).Interface()

	switch newVal.(type) {
	case int64:
		intVal, _ := mv.(int)
		v.SetInt(int64(intVal))
	case float64:
		floatVal, _ := mv.(float64)
		v.SetFloat(floatVal)
	case string:
		stringVal, _ := mv.(string)
		v.SetString(stringVal)
	case bool:
		boolVal, _ := mv.(bool)
		v.SetBool(boolVal)
	default:
		return fmt.Errorf("can't cast value of type %T with value %v, to type %v", mv, mv, t)
	}

	return nil
}

func bindRow(jsonRow []byte, st reflect.Type) (reflect.Value, error) {
	var row map[string]interface{}
	elem := reflect.New(st)

	if err := json.Unmarshal(jsonRow, &row); err != nil {
		return reflect.Value{}, err
	}

	if err := bindMapToStruct(row, elem); err != nil {
		return reflect.Value{}, err
	}

	return elem, nil
}
