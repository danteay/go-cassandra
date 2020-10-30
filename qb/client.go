package qb

import "github.com/gocql/gocql"

type client struct {
	session *gocql.Session
}

// NewClient creates a new cassandra client manager from config
func NewClient(conf Config) (Client, error) {
	s, err := getSession(conf)
	if err != nil {
		return nil, err
	}

	return &client{session: s}, nil
}

var _ Client = &client{}

func (c *client) Select(f ...string) *SelectQuery {
	return newSelectQuery(c.session, f...)
}

func (c *client) Insert(f ...string) *InsertQuery {
	return newInsertQuery(c.session, f...)
}

func (c *client) Update(t string) *UpdateQuery {
	return newUpdateQuery(c.session, t)
}

func (c *client) Delete() *DeleteQuery {
	return newDeleteQuery(c.session)
}

func (c *client) Count() *CountQuery {
	return newCountQuery(c.session)
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
