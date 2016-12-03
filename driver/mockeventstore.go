package eventchaindriver


type MockEventStore struct {
	LoadEventsFunc func(string) (Events, error)
}


func (eventStore *MockEventStore) LoadEvents(id string) (Events, error) {
	if nil == eventStore {
		return nil, errNilReceiver
	}

	fn := eventStore.LoadEventsFunc
	if nil == fn {
		return nil, errNilFunc
	}

	events, err := fn(id)
	if nil != err {
		return nil, err
	}

	return events, nil
}
