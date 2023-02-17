package qinsert

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/qinsert/mocks"
)

func TestInsertQuery_build(t *testing.T) {
	tt := []struct {
		name    string
		table   string
		fields  []string
		values  []interface{}
		res     string
		expArgs []interface{}
	}{
		{
			name:    "query insert with boolean values",
			table:   "test1",
			fields:  []string{"col1", "col2", "col3"},
			values:  []interface{}{1, "123", true},
			expArgs: []interface{}{1, "123", true},
			res:     "INSERT INTO test1 (col1,col2,col3) VALUES (?,?,?)",
		},
		{
			name:    "query insert with boolean values as string",
			table:   "test2",
			fields:  []string{"col1", "col2", "col3"},
			values:  []interface{}{"1", "123", "true"},
			expArgs: []interface{}{"1", "123", "true"},
			res:     "INSERT INTO test2 (col1,col2,col3) VALUES (?,?,?)",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			client := mocks.NewClient(t)
			client.On("Debug").Return(false)

			q := New(client).Fields(test.fields...).Into(test.table).Values(test.values...)
			query := q.build()

			assert.Equal(t, test.res, query)
			assert.Equal(t, test.expArgs, q.args)
		})
	}
}
