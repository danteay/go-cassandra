package types

import "golang.org/x/exp/slices"

// QueryType represents all different operations that be query build builder can construct
type QueryType string

const (
	// Select represents SELECT statement
	Select QueryType = "select"

	// Update represents UPDATE statement
	Update QueryType = "update"

	// Insert represents INSERT statement
	Insert QueryType = "insert"

	// Delete represents DELETE statement
	Delete QueryType = "delete"

	// Count represents SELECT COUNT statement
	Count QueryType = "count"
)

var queryTypes = []QueryType{
	Select,
	Update,
	Insert,
	Delete,
	Count,
}

// String convert QueryType to a string.
func (q QueryType) String() string {
	if slices.Contains(queryTypes, q) {
		return string(q)
	}

	return "unknown"
}

// IsValid verifies that the value of the QueryType is valid.
func (q QueryType) IsValid() bool {
	return q.String() != "unknown"
}
