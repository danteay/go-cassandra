package qb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/types"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner/mocks"
)

func TestQuery_Update(t *testing.T) {
	t.Run("new update query", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Update(client)

		assert.NotNil(t, q.runner)
		assert.Equal(t, types.Update, q.queryType)
	})

	t.Run("set fields", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Update(client)

		q.Set("field", "value")

		assert.Equal(t, []string{"field"}, q.set)
		assert.Equal(t, []interface{}{"value"}, q.args)
	})

	t.Run("build cql", func(t *testing.T) {
		client := mocks.NewClient(t)
		query, args, err := Update(client).
			From("table").
			Set("field", "value").
			Where("field", where.Eq, "eq_val").
			Where("field2", where.Gt, "gt_val").
			ToCQL()

		expQuery := `UPDATE table SET field=? WHERE field=? AND field2>?`

		assert.NoError(t, err)
		assert.Equal(t, expQuery, query)
		assert.Equal(t, []interface{}{"value", "eq_val", "gt_val"}, args)
	})
}
