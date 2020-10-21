package qb

import "github.com/scylladb/gocqlx/qb"

// WhereOp represents a where operation type
type WhereOp string

const (
	// Eq equals
	Eq WhereOp = "="

	// Gt grater or equal
	Ge WhereOp = ">="

	// Lt lower or equal
	Le WhereOp = "<="

	// L lower
	L WhereOp = "<"

	// G grater
	G WhereOp = ">"
)

type whereStm struct {
	field string
	op    WhereOp
}

func buildWhere(stms []whereStm) []qb.Cmp {
	var ops []qb.Cmp

	for _, op := range stms {
		switch op.op {
		case Eq:
			ops = append(ops, qb.Eq(op.field))
		case Ge:
			ops = append(ops, qb.GtOrEq(op.field))
		case Le:
			ops = append(ops, qb.GtOrEq(op.field))
		case G:
			ops = append(ops, qb.GtOrEq(op.field))
		case L:
			ops = append(ops, qb.GtOrEq(op.field))
		}
	}

	return ops
}
