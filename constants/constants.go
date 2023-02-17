package constants

// DatetimeLayout is the default datetime layout format that is expected from cassandra datetime values. This layout
// will be used on binder to fill time.Time fields.
const DatetimeLayout string = "2006-01-02 15:04:05.000Z"

// Desc represents DESC order filter
const Desc string = "DESC"

// Asc represents ASC order filter
const Asc string = "ASC"

// Tag represents the struct tag that should be placed on the structs that will be bound as results.
const Tag string = "cql"
