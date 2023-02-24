package config

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/danteay/go-cassandra/constants"
	"github.com/danteay/go-cassandra/logging"
)

// Default return an instance of Config obtaining his values from environment variables and in case these vars has
// no values, a default value will be used in place for each field.
func Default() Config {
	return Config{
		Port:                     getDefaultPort(),
		KeyspaceName:             getKeyspaceName(),
		Username:                 getUsername(),
		Password:                 getPassword(),
		ContactPoints:            getContactPoints(),
		Debug:                    getDebug(),
		ProtoVersion:             getProtoVersion(),
		Consistency:              constants.ConsistencyAny,
		CaPath:                   getCaPath(),
		DisableInitialHostLookup: getDisableInitialHostLookup(),
		Timeout:                  getTimeout(),
		ConnectTimeout:           getConnectionTimeout(),
		PrintQuery:               logging.DefaultDebugPrint,
		NoHostRetries:            getNoHostRetries(),
		DatetimeLayout:           getDatetimeLayout(),
	}
}

func getDefaultPort() int {
	port := 9042

	if osPort := os.Getenv("CASSANDRA_PORT"); osPort != "" {
		aux, err := strconv.Atoi(osPort)
		if err == nil {
			return aux
		}
	}

	return port
}

func getKeyspaceName() string {
	keyspaceName := ""

	if osKeyspaceName := os.Getenv("CASSANDRA_KEYSPACE_NAME"); osKeyspaceName != "" {
		return osKeyspaceName
	}

	return keyspaceName
}

func getUsername() string {
	username := ""

	if osUsername := os.Getenv("CASSANDRA_USERNAME"); osUsername != "" {
		return osUsername
	}

	return username
}

func getPassword() string {
	password := ""

	if osPassword := os.Getenv("CASSANDRA_PASSWORD"); osPassword != "" {
		return osPassword
	}

	return password
}

func getContactPoints() []string {
	contactPoints := []string{"127.0.0.1"}
	osContactPoints := os.Getenv("CASSANDRA_CONTACT_POINTS")

	if osContactPoints == "" {
		return contactPoints
	}

	return strings.Split(osContactPoints, ",")
}

func getDebug() bool {
	osDebug := os.Getenv("CASSANDRA_DEBUG")

	if osDebug != "" {
		return strings.ToLower(osDebug) == "true"
	}

	return false
}

func getProtoVersion() int {
	protoVersion := 4
	osProtoVersion := os.Getenv("CASSANDRA_PROTO_VERSION")

	if osProtoVersion == "" {
		return protoVersion
	}

	cast, err := strconv.Atoi(osProtoVersion)
	if err != nil {
		return protoVersion
	}

	return cast
}

func getCaPath() string {
	caPath := ""

	if osCaPath := os.Getenv("CASSANDRA_CA_PATH"); osCaPath != "" {
		return osCaPath
	}

	return caPath
}

func getDisableInitialHostLookup() bool {
	osDisableInitialHostLookup := os.Getenv("CASSANDRA_DISABLE_INITIAL_HOST_LOOKUP")

	if osDisableInitialHostLookup != "" {
		return strings.ToLower(osDisableInitialHostLookup) == "true"
	}

	return false
}

func getTimeout() time.Duration {
	timeout := time.Second * 30
	osTimeout := os.Getenv("CASSANDRA_TIMEOUT")

	if osTimeout == "" {
		return timeout
	}

	cast, err := strconv.Atoi(osTimeout)
	if err != nil {
		return timeout
	}

	return time.Second * time.Duration(cast)
}

func getConnectionTimeout() time.Duration {
	timeout := time.Second * 30
	osTimeout := os.Getenv("CASSANDRA_CONNECTION_TIMEOUT")

	if osTimeout == "" {
		return timeout
	}

	cast, err := strconv.Atoi(osTimeout)
	if err != nil {
		return timeout
	}

	return time.Second * time.Duration(cast)
}

func getNoHostRetries() int {
	noHostRetries := 2
	osNoHostRetries := os.Getenv("CASSANDRA_NO_HOST_RETRIES")

	if osNoHostRetries == "" {
		return noHostRetries
	}

	cast, err := strconv.Atoi(osNoHostRetries)
	if err != nil {
		return noHostRetries
	}

	return cast
}

func getDatetimeLayout() string {
	datetimeLayout := constants.DatetimeLayout

	if osDatetimeLayout := os.Getenv("CASSANDRA_DATETIME_LAYOUT"); osDatetimeLayout != "" {
		return osDatetimeLayout
	}

	return datetimeLayout
}
