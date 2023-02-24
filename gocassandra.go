// Package gocassandra implements
package gocassandra

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/logging"
	"github.com/danteay/go-cassandra/qb"
)

// Client is the main cassandra client abstraction to work with the database
type Client interface {
	// Select start a select query
	Select(f ...string) *qb.Query

	// Insert start a new insert query statement
	Insert(f ...string) *qb.Query

	// Update start an update query statement
	Update(t string) *qb.Query

	// Delete start a new delete query statement
	Delete() *qb.Query

	// Count start new count query statement
	Count() *qb.Query

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
func New(options ...Option) (Client, error) {
	co := clientOptions{}

	for _, opt := range options {
		opt(&co)
	}

	setConfig(&co)

	if err := setSessionFromOpts(&co); err != nil {
		return nil, err
	}

	return &client{
		session:    co.session,
		config:     *co.config,
		canRestart: co.canRestart,
	}, nil
}

func setConfig(co *clientOptions) {
	if co.config != nil {
		return
	}

	def := config.Default()
	co.config = &def
}

func setSessionFromOpts(co *clientOptions) error {
	if co.session != nil {
		return nil
	}

	session, err := getSession(*co.config)
	if err != nil {
		return err
	}

	co.session = session
	co.canRestart = true

	return nil
}
