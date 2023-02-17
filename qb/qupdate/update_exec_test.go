package qupdate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/qupdate/mocks"
	"github.com/danteay/go-cassandra/qb/where"
)

func TestQuery_build(t *testing.T) {
	type set struct {
		field string
		value interface{}
	}

	tt := []struct {
		name    string
		table   string
		set     []set
		where   []where.Stm
		expArgs []interface{}
		res     string
	}{
		{
			name:  "query update with no set values and Eq filter",
			table: "test_table",
			set:   []set{},
			where: []where.Stm{
				{
					Field: "field1",
					Op:    where.Eq,
					Value: 0,
				},
			},
			expArgs: []interface{}{0},
			res:     "UPDATE test_table SET  WHERE field1=?",
		},
		{
			name:  "query update with 2 set values and no filters",
			table: "test_table",
			set: []set{
				{
					field: "field1",
					value: 1,
				},
				{
					field: "field2",
					value: true,
				},
			},
			where:   []where.Stm{},
			expArgs: []interface{}{1, true},
			res:     "UPDATE test_table SET field1=?,field2=?",
		},
		{
			name:  "query update with 2 set values and Eq filter",
			table: "test_table",
			set: []set{
				{
					field: "field1",
					value: 1,
				},
				{
					field: "field2",
					value: true,
				},
			},
			where: []where.Stm{
				{
					Field: "field1",
					Op:    where.Eq,
					Value: 0,
				},
			},
			expArgs: []interface{}{1, true, 0},
			res:     "UPDATE test_table SET field1=?,field2=? WHERE field1=?",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			client := mocks.NewClient(t)
			client.On("Debug").Return(false)

			q := New(client).Table(test.table)

			for _, s := range test.set {
				q = q.Set(s.field, s.value)
			}

			for _, w := range test.where {
				q = q.Where(w.Field, w.Op, w.Value)
			}

			qs, err := q.build()

			assert.NoError(t, err)
			assert.Equal(t, test.res, qs)
			assert.Equal(t, test.expArgs, q.args)
		})
	}
}
