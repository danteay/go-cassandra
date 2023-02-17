package qcount

import (
	"reflect"
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/qcount/mocks"
	"github.com/danteay/go-cassandra/qb/where"
)

func TestNew(t *testing.T) {
	s := &gocql.Session{}

	client := mocks.NewClient(t)
	client.On("Session").Return(s)

	q := New(client)

	assert.Same(t, q.client.Session(), s)
}

func TestQuery_From(t *testing.T) {
	q := &Query{}

	q.From("test")
	assert.Equal(t, "test", q.table)

	q.From("test2")
	assert.Equal(t, "test2", q.table)
}

func TestQuery_Column(t *testing.T) {
	q := &Query{}

	q.Column("field")
	if q.column != "field" {
		t.Errorf("exp: field got: %v", q.column)
	}

	q.Column("field2")
	if q.column != "field2" {
		t.Errorf("exp: field2 got: %v", q.column)
	}
}

func TestQuery_Where(t *testing.T) {
	tt := []struct {
		field   string
		op      where.Operator
		value   interface{}
		expArgs []interface{}
		expStm  []where.Stm
	}{
		{
			field:   "field1",
			op:      where.Eq,
			value:   nil,
			expArgs: []interface{}{nil},
			expStm: []where.Stm{
				{
					Field: "field1",
					Op:    where.Eq,
					Value: nil,
				},
			},
		},
		{
			field:   "field2",
			op:      where.Gt,
			value:   1,
			expArgs: []interface{}{1},
			expStm: []where.Stm{
				{
					Field: "field2",
					Op:    where.Gt,
					Value: 1,
				},
			},
		},
		{
			field:   "field3",
			op:      where.Lt,
			value:   1,
			expArgs: []interface{}{1},
			expStm: []where.Stm{
				{
					Field: "field3",
					Op:    where.Lt,
					Value: 1,
				},
			},
		},
	}

	for _, test := range tt {
		q := &Query{}

		q.Where(test.field, test.op, test.value)

		if !reflect.DeepEqual(q.args, test.expArgs) {
			t.Errorf("exp: %v got: %v", test.expArgs, q.args)
		}

		if !reflect.DeepEqual(q.where, test.expStm) {
			t.Errorf("exp: %v got: %v", test.expStm, q.where)
		}
	}
}
