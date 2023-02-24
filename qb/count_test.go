package qb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/types"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner/mocks"
)

func TestQuery_Count(t *testing.T) {
	t.Run("new count query", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Count(client)

		assert.NotNil(t, q.runner)
		assert.Equal(t, types.Count, q.queryType)
	})

	t.Run("set from table", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Count(client)

		q.From("test2")
		assert.Equal(t, "test2", q.table)
	})

	t.Run("set count column", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Count(client)

		q.Column("test")
		assert.Equal(t, "test", q.countColumn)
	})

	t.Run("set where clauses", func(t *testing.T) {
		tt := []struct {
			name    string
			field   string
			op      where.Operator
			value   interface{}
			expArgs []interface{}
			expStm  []where.Stm
		}{
			{
				name:    "eq",
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
				name:    "gt",
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
				name:    "lw",
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
			t.Run(test.name, func(t *testing.T) {
				client := mocks.NewClient(t)
				q := Count(client)

				q.Where(test.field, test.op, test.value)

				assert.Equal(t, test.expArgs, q.args)
				assert.Equal(t, test.expStm, q.where)
			})
		}
	})

	t.Run("build cql with just table specified", func(t *testing.T) {
		client := mocks.NewClient(t)

		query, args, err := Count(client).From("table").ToCQL()

		expQuery := `SELECT count(*) FROM table`

		assert.NoError(t, err)
		assert.Equal(t, expQuery, query)
		assert.Equal(t, []interface{}{}, args)
	})

	t.Run("to cql build", func(t *testing.T) {
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
			{
				name:           "select count all with Eq and GtOrEq filters and allow filtering",
				table:          "test",
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
				expQuery: "SELECT count(*) FROM test WHERE field1=? AND field2>=? ALLOW FILTERING",
			},
		}

		for _, test := range tt {
			t.Run(test.name, func(t *testing.T) {
				client := mocks.NewClient(t)

				q := Count(client).From(test.table).Column(test.column)

				for _, w := range test.where {
					q = q.Where(w.Field, w.Op, w.Value)
				}

				if test.allowFiltering {
					q = q.AllowFiltering()
				}

				query, args, err := q.ToCQL()

				assert.NoError(t, err)
				assert.Equal(t, test.expQuery, query)
				assert.Equal(t, test.expArgs, args)
			})
		}
	})
}
