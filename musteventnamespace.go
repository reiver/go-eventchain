package eventchain


// MustEventNamespace wraps an EventNamespace, and provides "Must" methods for it.
//
// ("Must" methods panic() when there is an error, instead of returning the error.)
//
// These are sometimes useful when writing unit tests.
//
// And useful in (other) cases where you want to panic() on an error.
type MustEventNamespace struct {
	EventNamespace EventNamespace

}


// EventName is the same as with the EventNamespace interface.
func (ns MustEventNamespace) EventName(s string) (EventName, error) {
	return ns.EventNamespace.EventName(s)
}

// Rand is the same as with the EventNamespace interface.
func (ns MustEventNamespace) Rand() (EventName, error) {
	return ns.EventNamespace.Rand()
}

// Validate is the same as with the EventNamespace interface.
func (ns MustEventNamespace) Validate(s string) error {
	return ns.EventNamespace.Validate(s)
}


// MustEventName is like EventName, except panic()s on error.
func (ns MustEventNamespace) MustEventName(s string) EventName {
	eventName, err := ns.EventName(s)
	if nil != err {
		panic(err)
	}

	return eventName
}

// MustRand is like Rand, except panic()s on error.
func (ns MustEventNamespace) MustRand() EventName {
	eventName, err := ns.Rand()
	if nil != err {
		panic(err)
	}

	return eventName
}

// MustValidate is like Validate, except panic()s on error.
func (ns MustEventNamespace) MustValidate(s string) {
	err := ns.Validate(s)
	if nil != err {
		panic(err)
	}

	return
}
