// Package gocassandra implements
package gocassandra

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/logging"
	"github.com/danteay/go-cassandra/qb/qcount"
	"github.com/danteay/go-cassandra/qb/qdelete"
	"github.com/danteay/go-cassandra/qb/qinsert"
	"github.com/danteay/go-cassandra/qb/qselect"
	"github.com/danteay/go-cassandra/qb/qupdate"
)

// Client is the main cassandra client abstraction to work with the database
type Client interface {
	// Select start a select query
	Select(f ...string) *qselect.Query

	// Insert start a new insert query statement
	Insert(f ...string) *qinsert.Query

	// Update start an update query statement
	Update(t string) *qupdate.Query

	// Delete start a new delete query statement
	Delete() *qdelete.Query

	// Count start new count query statement
	Count() *qcount.Query

	// Session return the plain session object to build some direct query
	Session() *gocql.Session

	// Debug return an assertion for debugging
	Debug() bool

	// PrintFn return the configured debug print function.
	PrintFn() logging.DebugPrint

	// Restart should close and start a new connection.
	Restart() error

	// Config return current client configuration
	Config() config.Config

	// Close ends cassandra connection pool
	Close()
}

// New creates a new cassandra client manager from config
func New(conf config.Config) (Client, error) {
	session, err := getSession(conf)
	if err != nil {
		return nil, err
	}

	return &client{
		session:    session,
		config:     conf,
		canRestart: true,
	}, nil
}

// NewWithSession creates a new cassandra client manager from a given gocql session.
func NewWithSession(session *gocql.Session, conf config.Config) (Client, error) {
	return &client{
		session:    session,
		config:     conf,
		canRestart: false,
	}, nil
}
