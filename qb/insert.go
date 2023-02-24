package qb

// Fields save query fields that should be used for insert query
func (q *Query) Fields(f ...string) *Query {
	q.fields = f
	return q
}

// Into set table to insert query
func (q *Query) Into(t string) *Query {
	q.table = t
	return q
}

// Values set values as query arguments for insert statement
func (q *Query) Values(v ...interface{}) *Query {
	q.args = v
	return q
}
