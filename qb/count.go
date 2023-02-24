package qb

// Column set count countColumn of the query
func (q *Query) Column(c string) *Query {
	if c == "" {
		c = "*"
	}

	q.countColumn = c
	return q
}
