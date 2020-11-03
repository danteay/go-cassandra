package qb

import (
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
}

// Client is the main cassandra client abstraction to work with the database
type Client interface {
	// Select start a qselect query
	Select(f ...string) *_select.Query

	// Insert start a new qinsert query statement
	Insert(f ...string) *qinsert.Query

	// Update start an qupdate query statement
	Update(t string) *qupdate.Query

	// Delete start a new qdelete query statement
	Delete() *delete2.Query

	// Count start new qcount query statement
	Count() *qcount.Query

	// Session return the plain session object to build some direct query
	Session() *gocql.Session

	// Close close cassandra connection pool
	Close()
}
