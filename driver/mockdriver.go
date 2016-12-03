package  eventchaindriver


type MockDriver struct{
	OpenDataStoreFunc func(args ...interface{})(DataStore, error)
}


func (driver *MockDriver) OpenDataStore(args ...interface{}) (DataStore, error) {
	if nil == driver {
		return nil, errNilReceiver
	}

	fn := driver.OpenDataStoreFunc
	if nil == fn {
		return nil, errNilFunc
	}

	dataStore, err := fn(args...)
	if nil != err {
		return nil, err
	}

	return dataStore, nil
}
