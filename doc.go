/*
Package eventchain provides basic building blocks for doing Event Sourcing and Log-Driven Development.

Example:

	import (
		"github.com/reiver/go-eventchain"
	)
	
	// ...
	
	dataStore, err := eventchain.OpenDataStore("database/sql:postgres", db)
	if nil != err {
		return err
	}
	
	
	const eventStoreName = "user_events" // <--------- Your code would likely have a different value here.
	
	evenStore, err := dataStore.OpenEventStore(eventStoreName)
	if nil != err {
		return err
	}
*/
package eventchain
