package  eventchaindriver


type MockDataStore struct{
	OpenEventStoreFunc func(name string, eventNames ...string)(EventStore, error)
}


func (dataStore *MockDataStore) OpenEventStore(name string, eventNames ...string) (EventStore, error) {
	if nil == dataStore {
		return nil, errNilReceiver
	}

	fn := dataStore.OpenEventStoreFunc
	if nil == fn {
		return nil, errNilFunc
	}

	eventStore, err := fn(name, eventNames...)
	if nil != err {
		return nil, err
	}

	return eventStore, nil
}
