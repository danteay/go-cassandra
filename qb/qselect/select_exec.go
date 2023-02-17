package qselect

// One execute query and return just one result on bind action
func (q *Query) One() error {
	query, err := q.build()
	if err != nil {
		return err
	}

	return q.runner.QueryOne(query, q.args, q.bind)
}

// All execute query and return all rows on bind action. Bindable struct should be a slice of structs
func (q *Query) All() error {
	query, err := q.build()
	if err != nil {
		return err
	}

	return q.runner.Query(query, q.args, q.bind)
}
