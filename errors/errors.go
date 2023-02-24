package errors

import "errors"

var (
	ErrNilBinding              = errors.New("go-cassandra: nil bind is not allowed")
	ErrNoPtrBinding            = errors.New("go-cassandra: bind value should be a pointer")
	ErrNoStructOrSliceBinding  = errors.New("go-cassandra: bind value should be a struct or slice")
	ErrNoSliceOfStructsBinding = errors.New("go-cassandra: bind value should be a slice of structs")
	ErrClosedConnection        = errors.New("go-cassandra: can execute on closed connection")
	ErrInvalidWhereOperator    = errors.New("go-cassandra: invalid where operator")
	ErrUnableToRestart         = errors.New("go-cassandra: unable to restart session")
	ErrInvalidQueryType        = errors.New("go-cassandra: invalid query type")
)
