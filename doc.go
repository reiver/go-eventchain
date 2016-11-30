/*
Package eventchain provides basic building blocks for doing Event Sourcing and Log-Driven Development.

Two concepts that this package has are: "data stores" and "event stores".

At the high level, an 'event store' contains the events for a single thing, and a 'data store' contains
one or more 'event stores'.

So for example, you might have 'event stores' for "user events", "purchase events", "warehouse inventory events", etc.

A little more low level....

If this was being backed by something like a Postgres database, then in that specific case the 'data store'
is the Postgres database, and an 'event store' is a specifc table in that Postgres database.

Alternatively, if one is working log files , a 'data store' could be a directory on a server (containing log files),
and an 'event store' could be a single log file.

Here is an example of connecting to a Postgres backed 'data store', and then using that 'data store' to
get access to an 'event store' within it.

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
