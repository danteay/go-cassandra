package gocassandra

type Consistency uint16

const (
	ConsistencyAny         Consistency = 0x00
	ConsistencyOne         Consistency = 0x01
	ConsistencyTwo         Consistency = 0x02
	ConsistencyThree       Consistency = 0x03
	ConsistencyQuorum      Consistency = 0x04
	ConsistencyAll         Consistency = 0x05
	ConsistencyLocalQuorum Consistency = 0x06
	ConsistencyEachQuorum  Consistency = 0x07
	ConsistencyLocalOne    Consistency = 0x0A
)
