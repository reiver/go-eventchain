package eventchaindriver


type EventStore interface {
	LoadEvents(string) (Events, error)
}
