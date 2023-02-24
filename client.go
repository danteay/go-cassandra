package gocassandra

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
	"github.com/danteay/go-cassandra/errors"
	"github.com/danteay/go-cassandra/logging"
	"github.com/danteay/go-cassandra/qb"
)

type client struct {
	canRestart bool
	config     config.Config
	session    *gocql.Session
}

var _ Client = &client{}

func (c *client) Select(f ...string) *qb.Query {
	return qb.Select(c).Fields(f...)
}

func (c *client) Insert(f ...string) *qb.Query {
	return qb.Insert(c).Fields(f...)
}

func (c *client) Update(t string) *qb.Query {
	return qb.Update(c).Table(t)
}

func (c *client) Delete() *qb.Query {
	return qb.Delete(c)
}

func (c *client) Count() *qb.Query {
	return qb.Count(c)
}

func (c *client) Debug() bool {
	return c.Config().Debug
}

func (c *client) PrintFn() logging.DebugPrint {
	return c.Config().PrintQuery
}

func (c *client) Close() {
	c.session.Close()
}

func (c *client) Session() *gocql.Session {
	return c.session
}

func (c *client) Config() config.Config {
	return c.config
}

func (c *client) Restart() error {
	c.Close()

	if !c.canRestart {
		return errors.ErrUnableToRestart
	}

	session, err := getSession(c.config)
	if err != nil {
		return err
	}

	c.session = session

	return nil
}
