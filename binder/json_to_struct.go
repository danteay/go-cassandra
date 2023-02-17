package binder

import (
	"encoding/json"
	"reflect"
)

// JsonToStruct bind a row of returned data on json format into a given struct type with `cql` tags that associate
// fields to struct members
func (b *Binder) JsonToStruct(jsonRow []byte, st reflect.Type) (reflect.Value, error) {
	var row map[string]interface{}

	if err := json.Unmarshal(jsonRow, &row); err != nil {
		return reflect.Value{}, err
	}

	elem, err := b.MapToStruct(row, st)
	if err != nil {
		return reflect.Value{}, err
	}

	return elem, nil
}
