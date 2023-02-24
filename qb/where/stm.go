package where

import (
	"github.com/scylladb/gocqlx/qb"

	"github.com/danteay/go-cassandra/errors"
)

// Stm defines a Where condition holding filed, operation and value of the condition
type Stm struct {
	Field string
	Op    Operator
	Value interface{}
}

func (ws Stm) Build() (qb.Cmp, error) {
	switch ws.Op {
	case Eq:
		return qb.Eq(ws.Field), nil
	case GtOrEq:
		return qb.GtOrEq(ws.Field), nil
	case LtOrEq:
		return qb.LtOrEq(ws.Field), nil
	case Gt:
		return qb.Gt(ws.Field), nil
	case Lt:
		return qb.Lt(ws.Field), nil
	case Like:
		return qb.Like(ws.Field), nil
	default:
		return qb.Cmp{}, errors.ErrInvalidWhereOperator
	}
}
