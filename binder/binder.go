package binder

import (
	"github.com/danteay/go-cassandra/config"
)

//go:generate mockery --name=Client --filename=client.go --structname=Client --output=mocks --outpkg=mocks

// Client represents the main client  that holds cassandra configuration
type Client interface {
	Config() config.Config
}

type Binder struct {
	client Client
}

// New Creates a new instance of a Binder pinter
func New(c Client) *Binder {
	return &Binder{client: c}
}
