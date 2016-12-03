package eventchaindriver


type Events interface {
	Close() error
	Err() error
	Next() bool
	Event() (Event, error)
}
