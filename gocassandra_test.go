package gocassandra

import (
	"fmt"
	"os"
	"testing"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/cassandra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoCassandraSuite struct {
	suite.Suite
	cassandraContainer *gnomock.Container
}

func (s *GoCassandraSuite) SetupTest() {
	cp := cassandra.Preset(
		cassandra.WithVersion("4.1"),
	)

	s.T().Log("Starting cassandra...")

	container, err := gnomock.Start(cp)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Log("Cassandra started")

	s.cassandraContainer = container
}

func (s *GoCassandraSuite) TearDownSuite() {
	if err := gnomock.Stop(s.cassandraContainer); err != nil {
		s.T().Log(err)
	}
}

func (s *GoCassandraSuite) BeforeTest(_, _ string) {
	s.T().Log("before testing adding environment")

	s.T().Setenv("CASSANDRA_CONTACT_POINTS", s.cassandraContainer.DefaultAddress())
	s.T().Setenv("CASSANDRA_PORT", fmt.Sprintf("%d", s.cassandraContainer.DefaultPort()))
	s.T().Setenv("CASSANDRA_USERNAME", cassandra.DefaultUser)
	s.T().Setenv("CASSANDRA_PASSWORD", cassandra.DefaultPassword)
}

func (s *GoCassandraSuite) TestSetupGnomock() {
	addrs := os.Getenv("CASSANDRA_CONTACT_POINTS")
	assert.NotEmpty(s.T(), addrs)
}

func TestGoCassandraSuite(t *testing.T) {
	suite.Run(t, new(GoCassandraSuite))
}
