package eventchain


type Iterator interface {
	Next() bool
	Error() error
	Event() Event
}
