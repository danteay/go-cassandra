package qb

import "github.com/gocql/gocql"

// Config is the main cassandra configuration needed
type Config struct {
	Port          int      `yaml:"port"`
	KeyspaceName  string   `yaml:"keyspace_name"`
	Username      string   `yaml:"username"`
	Password      string   `yaml:"password"`
	ContactPoints []string `yaml:"contact_points"`
}

type columns []string

// Order definition for type or order on a select query
type Order string

const (
	// Desc represents DESC order filter
	Desc Order = "DESC"

	// Asc represents ASC order filter
	Asc Order = "ASC"
)

const modelsTag = "gocql"

// Client is the main cassandra client abstraction to work with the database
type Client interface {
	// Select start a select query
	Select(f ...string) *SelectQuery

	// Insert start a new insert query statement
	Insert(f ...string) *InsertQuery

	// Update start an update query statement
	Update() *UpdateQuery

	// Delete start a new delete query statement
	Delete() *DeleteQuery

	// Count start new count query statement
	Count() *CountQuery

	// Session return the plain session object to build some direct query
	Session() *gocql.Session

	// Close close cassandra connection pool
	Close()
}
