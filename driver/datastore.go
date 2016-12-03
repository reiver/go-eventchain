package eventchaindriver


type DataStore interface {
	OpenEventStore(name string, eventNames ...string) (EventStore, error)
}
