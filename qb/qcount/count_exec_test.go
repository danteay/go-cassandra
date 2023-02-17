package qcount

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/qcount/mocks"
	"github.com/danteay/go-cassandra/qb/where"
)

func TestQuery_build(t *testing.T) {
	tt := []struct {
		name           string
		table          string
		column         string
		where          []where.Stm
		expQuery       string
		expArgs        []interface{}
		allowFiltering bool
	}{
		{
			name:   "select count with Eq and Gt filters",
			table:  "test",
			column: "*",
			where: []where.Stm{
				{
					Field: "field1",
					Op:    where.Eq,
					Value: true,
				},
				{
					Field: "field2",
					Op:    where.Gt,
					Value: 123,
				},
			},
			expArgs:  []interface{}{true, 123},
			expQuery: "SELECT count(*) FROM test WHERE field1=? AND field2>?",
		},
		{
			name:   "select count with Eq and GtOrEq filters",
			table:  "test",
			column: "id",
			where: []where.Stm{
				{
					Field: "field1",
					Op:    where.Eq,
					Value: true,
				},
				{
					Field: "field2",
					Op:    where.GtOrEq,
					Value: 123,
				},
			},
			expArgs:  []interface{}{true, 123},
			expQuery: "SELECT count(id) FROM test WHERE field1=? AND field2>=?",
		},
		{
			name:           "select count with Eq and GtOrEq filters and allow filtering",
			table:          "test",
			column:         "id",
			allowFiltering: true,
			where: []where.Stm{
				{
					Field: "field1",
					Op:    where.Eq,
					Value: true,
				},
				{
					Field: "field2",
					Op:    where.GtOrEq,
					Value: 123,
				},
			},
			expArgs:  []interface{}{true, 123},
			expQuery: "SELECT count(id) FROM test WHERE field1=? AND field2>=? ALLOW FILTERING",
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			client := mocks.NewClient(t)
			client.On("Debug").Return(false)

			q := New(client).From(test.table).Column(test.column)

			for _, w := range test.where {
				q = q.Where(w.Field, w.Op, w.Value)
			}

			if test.allowFiltering {
				q = q.AllowFiltering()
			}

			query, err := q.build()

			assert.NoError(t, err)
			assert.Equal(t, test.expQuery, query)
			assert.Equal(t, test.expArgs, q.args)
		})
	}
}
