package eventchain


import (
	"encoding/json"
)


// UnmarshalJSON akes EventName fit the Unmarshaler interface from the "encoding/json" package.
func (eventName *internalEventName) UnmarshalJSON(src []byte) error {
	if nil == eventName {
		return errNilReceiver
	}

	var s string
	if err := json.Unmarshal(src, &s); nil != err {
		return err
	}

	if err := eventName.Scan(s); nil != err {
		return err
	}

	return nil
}
