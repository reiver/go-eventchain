package eventchain


// String returns a string representation of the event name.
//
// String makes EventName fit the fmt.Stringer interface.
func (eventName internalEventName) String() string {
	return eventName.s
}
