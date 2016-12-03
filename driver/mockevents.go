package eventchaindriver


type MockEvents struct {
	Events []Event
	i int
}


func (events *MockEvents) Err() error {
	return nil
}

func (events *MockEvents) Next() bool {
	if nil == events {
		panic(errNilReceiver)
	}

	es := events.Events
	if nil == es {
		return false
	}

	lenES := len(es)
	if 1 > lenES {
		return false
	}

	i := events.i
	if lenES <= i {
		return false
	}
	events.i++


	return true
}

func (events *MockEvents) Event() (Event, error) {
	if nil == events {
		return nil, errNilReceiver
	}

	es := events.Events
	if nil == es {
		return nil, errNilEvents
	}

	lenES := len(es)
	if 1 > lenES {
		return nil, errNoEvents
	}

	i := events.i
	i--
	if i < 0 {
		return nil, errPremature
	}
	if lenES <= i {
		return nil, errPostmature
	}


	event := es[i]
	if nil == event {
		return nil, errNilEvent
	}

	return event, nil
}

