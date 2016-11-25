package eventchain


import (
	sqldriver "database/sql/driver"
)


// Value makes EventName fit the `Valuer` interface from "database/sql/driver".
func (eventName internalEventName) Value() (sqldriver.Value, error) {
	s := eventName.String()
	return s, nil
}
