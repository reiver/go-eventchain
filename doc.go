/*
Package eventchain provides basic building blocks for doing Event Sourcing and Log-Driven Development.

Two concepts that this package has are: "data stores" and "event stores".

At the high level, an 'event store' contains the events for a single thing, and a 'data store' contains
one or more 'event stores'.

So for example, you might have 'event stores' for "user events", "purchase events", "warehouse inventory events", etc.

A little more low level....

If this was being backed by something like a Postgres database, then in that specific case the 'data store'
is the Postgres database.

And an 'event store' is a specifc table in that Postgres database. (Maybe names something like "user_events",
"purchase_events", "warehouse_inventory_events", etc.)

Alternatively, if one is working log files, a 'data store' could be a directory on a server (containing all log files),
and an 'event store' could be a single log file name pattern (which would include the rotated out files).

Here is an example of connecting to a Postgres backed 'data store'.

	import (
		"github.com/reiver/go-eventchain"
	)
	
	// ...
	
	dataStore, err := eventchain.OpenDataStore("database/sql:postgres", db)
	if nil != err {
		return err
	}

Now that we have that 'data store', we can use it to get access to an 'event store' within it.

Note that we need to declare a event names that are part of this `event store`.

	eventNamespace, err := eventchain.basicEventNamespace.Declare(
		"ATE_AN_APPLE", //  \
		"ATE_A_BANANA", //  | <--------- You could would likely have different value here.
		"ATE_A_CHERRY", //  /
	)
	if nil != err {
		return err
	}
	
	
	const eventStoreName = "hungry_human_events" // <--------- Your code would likely have a different value here.
	
	evenStore, err := dataStore.OpenEventStore(eventStoreName, eventNamespace)
	if nil != err {
		return err
	}

In that example, we grabbed the "hungry_human_events" event store.

And that 'event store' has exactly 3 possible events: "ATE_AN_APPLE", "ATE_A_BANANA", and "ATE_A_CHERRY".

This would contain the events for all the "hungry humans" we have.

We would then want to get the events for just one "hungry human".

We would do that with code like the following.

	const eventChainID = "W5_BmUq.RVIdbSgdhrezFWJHU_Uh3jT.Oe20rRUe.MKhH"
	
	eventChain, err := eventStore.LoadEventChain(eventChainID)
	if nil != err {
		return err
	}


*/
package eventchain
