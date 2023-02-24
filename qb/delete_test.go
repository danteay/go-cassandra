package qb

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/danteay/go-cassandra/qb/types"
	"github.com/danteay/go-cassandra/qb/where"
	"github.com/danteay/go-cassandra/runner/mocks"
)

func TestQuery_Delete(t *testing.T) {
	t.Run("new delete query", func(t *testing.T) {
		client := mocks.NewClient(t)
		q := Delete(client)

		assert.NotNil(t, q.runner)
		assert.Equal(t, types.Delete, q.queryType)
	})

	t.Run("build cql", func(t *testing.T) {
		client := mocks.NewClient(t)
		query, args, err := Delete(client).
			From("table").
			Where("field", where.Eq, "val").
			Where("field2", where.Gt, 10).
			ToCQL()

		expQuery := `DELETE FROM table WHERE field=? AND field2>?`

		assert.NoError(t, err)
		assert.Equal(t, expQuery, query)
		assert.Equal(t, []interface{}{"val", 10}, args)
	})
}
