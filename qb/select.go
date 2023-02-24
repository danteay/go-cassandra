package qb

import "strings"

// Distinct adds a DISTINCT clause to select fields.
func (q *Query) Distinct() *Query {
	q.distinct = true
	return q
}

// OrderBy adds `order by` selection statement fields.
func (q *Query) OrderBy(ob []string, o string) *Query {
	q.orderBy = ob
	q.order = strings.ToUpper(o)
	return q
}

// GroupBy adds `group by` statement fields.
func (q *Query) GroupBy(f ...string) *Query {
	q.orderBy = f
	return q
}

// Limit adds `limit` query statement.
func (q *Query) Limit(l uint) *Query {
	q.limit = l
	return q
}
