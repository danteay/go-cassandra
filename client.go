package gocassandra

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/qb/qcount"
	"github.com/danteay/go-cassandra/qb/qdelete"
	"github.com/danteay/go-cassandra/qb/qinsert"
	"github.com/danteay/go-cassandra/qb/qselect"
	"github.com/danteay/go-cassandra/qb/query"
	"github.com/danteay/go-cassandra/qb/qupdate"
)

type client struct {
	canRestart bool
	config     Config
	session    *gocql.Session
	printQuery query.DebugPrint
}

var _ Client = &client{}

func (c *client) Select(f ...string) *qselect.Query {
	return qselect.New(c).Fields(f...)
}

func (c *client) Insert(f ...string) *qinsert.Query {
	return qinsert.New(c).Fields(f...)
}

func (c *client) Update(t string) *qupdate.Query {
	return qupdate.New(c).Table(t)
}

func (c *client) Delete() *qdelete.Query {
	return qdelete.New(c)
}

func (c *client) Count() *qcount.Query {
	return qcount.New(c)
}

func (c *client) Debug() bool {
	return c.config.Debug
}

func (c *client) PrintFn() query.DebugPrint {
	return c.printQuery
}

func (c *client) Close() {
	c.session.Close()
}

func (c *client) Session() *gocql.Session {
	return c.session
}

func (c *client) Config() Config {
	return c.config
}

func (c *client) Restart() error {
	c.Close()

	session, err := getSession(c.config)
	if err != nil {
		return err
	}

	c.session = session

	return nil
}
