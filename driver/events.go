package eventchaindriver


type Events interface {
	Err() error
	Next() bool
	Event() (Event, error)
}
