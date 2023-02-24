package qb

// Exec build CQL query and execute it. Returns error in case of build or execution errors.
func (q *Query) Exec() error {
	query, args, err := q.ToCQL()
	if err != nil {
		return err
	}

	return q.runner.QueryNone(query, args)
}

// One receive a pointer to a struct with `cql` tags to bind result first result of the query. Return error in case of
// build, execution or bind error.
func (q *Query) One(b interface{}) error {
	query, args, err := q.ToCQL()
	if err != nil {
		return err
	}

	return q.runner.QueryOne(query, args, b)
}

// All receive a pointer to a slice of structs with `cql` tags to bind result first result of the query. Return error in
// case of build, execution or bind error.
func (q *Query) All(b interface{}) error {
	query, args, err := q.ToCQL()
	if err != nil {
		return err
	}

	return q.runner.Query(query, args, b)
}
