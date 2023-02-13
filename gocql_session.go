package gocassandra

import "github.com/gocql/gocql"

func getSession(c Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(c.ContactPoints...)
	cluster.Keyspace = c.KeyspaceName
	cluster.Consistency = gocql.Consistency(c.Consistency)
	cluster.ProtoVersion = c.ProtoVersion

	if c.Port != 0 {
		cluster.Port = c.Port
	}

	if c.DisableInitialHostLookup {
		cluster.DisableInitialHostLookup = c.DisableInitialHostLookup
	}

	if c.Username != "" && c.Password != "" {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: c.Username,
			Password: c.Password,
		}
	}

	if c.CaPath != "" {
		cluster.SslOpts = &gocql.SslOptions{
			CaPath: c.CaPath,
		}
	}

	if c.Timeout != 0 {
		cluster.Timeout = c.Timeout
	}

	if c.ConnectTimeout != 0 {
		cluster.ConnectTimeout = c.ConnectTimeout
	}

	return cluster.CreateSession()
}
