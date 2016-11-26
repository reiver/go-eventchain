package eventchain


func (ns BasicEventNamespace) Validate(name string) error {
	names := ns.names
	if nil == names {
		return errInternalError
	}

	if _, ok := names[name]; !ok {
		return errNotFound
	}

	return nil
}
