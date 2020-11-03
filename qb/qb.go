package qb

import (
	"github.com/danteay/go-cassandra/qb/query"

	"github.com/danteay/go-cassandra/qb/qcount"
	delete2 "github.com/danteay/go-cassandra/qb/qdelete"
	"github.com/danteay/go-cassandra/qb/qinsert"
	_select "github.com/danteay/go-cassandra/qb/qselect"
	"github.com/danteay/go-cassandra/qb/qupdate"
	"github.com/gocql/gocql"
)

// Config is the main cassandra configuration needed
type Config struct {
	Port          int      `yaml:"port" json:"port"`
	KeyspaceName  string   `yaml:"keyspace_name" json:"keyspace_name"`
	Username      string   `yaml:"username" json:"username"`
	Password      string   `yaml:"password" json:"password"`
	ContactPoints []string `yaml:"contact_points" json:"contact_points"`
	Debug         bool     `yaml:"debug" json:"debug"`
	PrintQuery    query.DebugPrint
}

// Client is the main cassandra client abstraction to work with the database
type Client interface {
	// Select start a select query
	Select(f ...string) *_select.Query

	// Insert start a new insert query statement
	Insert(f ...string) *qinsert.Query

	// Update start an update query statement
	Update(t string) *qupdate.Query

	// Delete start a new delete query statement
	Delete() *delete2.Query

	// Count start new count query statement
	Count() *qcount.Query

	// Session return the plain session object to build some direct query
	Session() *gocql.Session

	// Close close cassandra connection pool
	Close()
}
