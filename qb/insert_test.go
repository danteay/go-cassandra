package qb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/types"
	"github.com/danteay/go-cassandra/runner/mocks"
)

func TestQuery_Insert(t *testing.T) {
	t.Run("new insert query", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Insert(client)

		assert.NotNil(t, q.runner)
		assert.Equal(t, types.Insert, q.queryType)
	})

	t.Run("set into", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Insert(client)

		q.Into("table")

		assert.Equal(t, "table", q.table)
	})

	t.Run("set fields", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Insert(client)

		q.Fields("field1", "field2")

		assert.Equal(t, []string{"field1", "field2"}, q.fields)
	})

	t.Run("set values", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Insert(client)

		q.Values("val1", "val2")

		assert.Equal(t, []interface{}{"val1", "val2"}, q.args)
	})

	t.Run("build to cql", func(t *testing.T) {
		client := mocks.NewClient(t)

		query, args, err := Insert(client).
			Into("table").
			Fields("field1", "field2").
			Values("val1", "val2").
			ToCQL()

		expQuery := `INSERT INTO table (field1,field2) VALUES (?,?)`

		assert.NoError(t, err)
		assert.Equal(t, []interface{}{"val1", "val2"}, args)
		assert.Equal(t, expQuery, query)
	})
}
