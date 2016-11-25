package eventchain


import (
	"database/sql"
	sqldriver "database/sql/driver"
	"encoding/json"
	"fmt"
)


// EventName represents the name of an event.
//
// It is compatible with the built-in Golang "database/sql"
// library, in that it fits the `Scanner` interface
// from "database/sql", and fits the `Valuer` interface
// from "database/sql/driver".
//
// It is also compatible with the built-in Golang "encoding/json"
// library, in that it fits the `Marshaler` interface and the
// `Unmarshaler` interface.
//
// In addition to that, it is also compatible with the fmt.Printf,
// fmt.Sprintf, etc family of functions, in it fits the `Stringer`
// interface from "fmt".
type EventName interface {
	sql.Scanner
	sqldriver.Valuer
	fmt.Stringer
}


type internalEventName struct {
	s string
	eventNamespace EventNamespace
}
