package config

import (
	"time"

	"github.com/danteay/go-cassandra/constants"
	"github.com/danteay/go-cassandra/logging"
)

// Config is the main cassandra configuration needed
type Config struct {
	Port                     int
	KeyspaceName             string
	Username                 string
	Password                 string
	ContactPoints            []string
	Debug                    bool
	ProtoVersion             int
	Consistency              constants.Consistency
	CaPath                   string
	DisableInitialHostLookup bool
	Timeout                  time.Duration
	ConnectTimeout           time.Duration
	PrintQuery               logging.DebugPrint
	NoHostRetries            int
	DatetimeLayout           string
}
