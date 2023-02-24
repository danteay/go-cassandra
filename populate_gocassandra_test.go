package gocassandra

import (
	"github.com/gocql/gocql"

	"github.com/danteay/go-cassandra/config"
)

func populateDefaultTest(cp []string, port int, user, password string) error {
	_, err := getSession(config.Config{ContactPoints: cp, Port: port, Username: user, Password: password})
	if err != nil {
		return err
	}

	return nil
}

func createKeySpace(sess *gocql.Session) error {
	query := `create keyspace testing with replication = {'class':'SimpleStrategy', 'replication_factor' : 1};`
	return sess.Query(query).Exec()
}

func createTestUsersTable(sess *gocql.Session) error {
	query := `CREATE TABLE users (
    id text,
    email text,
	age integer
    created_at timestamp,
	updated_at timestamp,
    PRIMARY KEY (email)
);`

	return sess.Query(query).Exec()
}

func createTestUsersData(sess *gocql.Session) error {
	return nil
}
