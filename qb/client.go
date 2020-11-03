package qb

import (
	"github.com/danteay/go-cassandra/qb/qcount"
	"github.com/danteay/go-cassandra/qb/qdelete"
	"github.com/danteay/go-cassandra/qb/qinsert"
	"github.com/danteay/go-cassandra/qb/qselect"
	"github.com/danteay/go-cassandra/qb/qupdate"
	"github.com/gocql/gocql"
)

type client struct {
	session *gocql.Session
}

// NewClient creates a new cassandra client manager from config
func NewClient(conf Config) (Client, error) {
	session, err := getSession(conf)
	if err != nil {
		return nil, err
	}

	return &client{session: session}, nil
}

var _ Client = &client{}

func (c *client) Select(f ...string) *qselect.Query {
	return qselect.New(c.session, f...)
}

func (c *client) Insert(f ...string) *qinsert.Query {
	return qinsert.New(c.session, f...)
}

func (c *client) Update(t string) *qupdate.Query {
	return qupdate.New(c.session, t)
}

func (c *client) Delete() *qdelete.Query {
	return qdelete.New(c.session)
}

func (c *client) Count() *qcount.Query {
	return qcount.New(c.session)
}

// Close finish cassandra session
func (c *client) Close() {
	c.session.Close()
}

func (c *client) Session() *gocql.Session {
	return c.session
}

func getSession(c Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(c.ContactPoints...)
	cluster.Keyspace = c.KeyspaceName
	cluster.Consistency = gocql.Quorum

	if c.Username != "" && c.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: c.Username,
			Password: c.Password,
		}
	}

	return cluster.CreateSession()
}
