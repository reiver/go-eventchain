package eventchain



func (dataStore *internalDataStore) OpenEventStore(name string) (EventStore, error) {
	if nil == dataStore {
		return nil, errNilReceiver
	}

	eventStore := internalEventStore{
		dataStore: dataStore,
		name: name,
	}

	return &eventStore, nil
}
