package eventchain


import (
	"database/sql"
	sqldriver "database/sql/driver"
	"encoding/json"
	"fmt"
)


// EventName represents the name of an event.
//
// EventName exists, for representing event names, rather that just
// using (the built-in Go type) `string`, to prevent invalid event
// names from being used. (Either through a typo, a bug, or
// intentionally.)
//
// It is compatible with the built-in Golang "database/sql"
// library, in that it fits the `Scanner` interface
// from "database/sql", and fits the `Valuer` interface
// from "database/sql/driver".
//
// (That means you can write values into it with the sql.Rows.Scan
// and sql.Row.Scan methods.)
//
// (And means you can read values from it with the sql.DB.Exec, sql.DB.Query,
// sql.DB.QueryRow, sql.Stmt.Exec, sql.Stmt.Query, sql.Stmt.QueryRow,
// sql.Tx.Exec, sql.Tx.Query, and sql.Tx.QueryRow methods.)
//
// It is also compatible with the built-in Golang "encoding/json"
// library, in that it fits the `Marshaler` interface and the
// `Unmarshaler` interface.
//
// (That means you can write values int it with the json.Unmarshal functions.)
//
// (And means you can read values from it with the sql.Marshal functions.)
//
// In addition to that, it is also compatible with the fmt.Printf,
// fmt.Sprintf, etc family of functions, in it fits the `Stringer`
// interface from "fmt".
type EventName interface {
	fmt.Stringer
	json.Marshaler
	json.Unmarshaler
	sql.Scanner
	sqldriver.Valuer
}


type internalEventName struct {
	s string
	eventNamespace EventNamespace
}
