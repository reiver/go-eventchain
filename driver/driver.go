package eventchaindriver


type Driver interface {
	OpenDataStore(args ...interface{}) (DataStore, error)
}
