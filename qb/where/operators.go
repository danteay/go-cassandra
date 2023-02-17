package where

// Operator represents a where operation type
type Operator string

const (
	// Eq equals
	Eq Operator = "="

	// GtOrEq grater or equal
	GtOrEq Operator = ">="

	// LtOrEq lower or equal
	LtOrEq Operator = "<="

	// Lt lower
	Lt Operator = "<"

	// Gt grater
	Gt Operator = ">"
)
