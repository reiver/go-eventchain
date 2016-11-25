package eventchain


// String returns a string representation of the event name.
//
// String makes EventName fit the `Stringer` interface, from the
// Go built-in package "fmt".
func (eventName internalEventName) String() string {
	return eventName.s
}
