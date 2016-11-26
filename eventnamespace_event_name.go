package eventchain


func (ns *BasicEventNamespace) EventName(name string) (EventName, error) {
	if nil == ns {
		return nil, errNilReceiver
	}

	if _, ok := ns.names[name]; !ok {
		return nil, errNotFound
	}

	eventName := internalEventName{
		s: name,
		eventNamespace: ns,
	}

	return &eventName, nil
}
