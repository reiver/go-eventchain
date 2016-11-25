package eventchain


import (
	"fmt"
)


func (eventName *internalEventName) Scan(src interface{}) error {
	if nil == eventName {
		return errNilReceiver
	}

	if bs, ok := src.([]byte); ok {
		src = string(bs)
	}

	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("Cannot scan something into %T that is not of type string or []byte: %T", eventName, src)
	}

	eventNamespace := eventName.eventNamespace
	if nil == eventNamespace {
		return errInternalError
	}

	if err := eventNamespace.Validate(s); nil != err {
		if errNotFound == err {
			return fmt.Errorf("%q is not a valid event name.", s)
		}
		return err
	}

	eventName.s = s

	return nil
}
