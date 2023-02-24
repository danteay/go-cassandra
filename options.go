package gocassandra

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
)

type clientOptions struct {
	session    *gocql.Session
	canRestart bool
	config     *config.Config
}

// Option is function implementation that defines how client is configured.
type Option func(options *clientOptions)

// Session it adds an external session of gocql to be managed by the client.
func Session(s *gocql.Session) Option {
	return func(opts *clientOptions) {
		opts.session = s
	}
}

// Config it adds specific configuration that should be used instead the default.
func Config(c config.Config) Option {
	return func(opts *clientOptions) {
		opts.config = &c
	}
}
