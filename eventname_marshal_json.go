package eventchain


import (
	"encoding/json"
)


// MarshalJSON makes EventName fit the Marshaler interface from the "encoding/json" package.
func (eventName internalEventName) MarshalJSON() ([]byte, error) {
	s := eventName.String()
	return json.Marshal(s)
}
