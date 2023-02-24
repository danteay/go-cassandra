package qb

// Table alias of From method.
func (q *Query) Table(t string) *Query {
	return q.From(t)
}

// Set save field and corresponding value to update.
func (q *Query) Set(f string, v interface{}) *Query {
	q.set = append(q.set, f)
	q.args = append(q.args, v)
	return q
}
