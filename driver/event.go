package eventchaindriver


import (
	"time"
)


type Event interface {
	ID() string
	WhenCreated() time.Time
	PrevID() (string, error)
	InteractionIdentifier() (string, error)
	WhenHappened() time.Time
	EventName() string
	Version() string
	Data() interface{}
}
