package logging

// DebugPrint defines a callback that prints query values
type DebugPrint func(q string, args []interface{}, err error)
