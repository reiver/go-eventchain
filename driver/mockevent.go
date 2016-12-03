package eventchaindriver


import (
	"time"
)


type MockEvent struct {
	MockedID string

	MockedWhenCreated time.Time

	MockedPrevID    string
	MockedPrevIDErr error

	MockedInteractionIdentifier    string
	MockedInteractionIdentifierErr error

	MockedWhenHappened time.Time

	MockedEventName string

	MockedVersion string

	MockedData interface{}
}


func (event *MockEvent) ID() string {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedID
}

func (event *MockEvent) WhenCreated() time.Time {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedWhenCreated
}

func (event *MockEvent) PrevID() (string, error) {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedPrevID, event.MockedPrevIDErr
}

func (event *MockEvent) InteractionIdentifier() (string, error) {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedInteractionIdentifier, event.MockedInteractionIdentifierErr
}

func (event *MockEvent) WhenHappened() time.Time {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedWhenHappened
}

func (event *MockEvent) EventName() string {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedEventName
}

func (event *MockEvent) Version() string {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedVersion
}

func (event *MockEvent) Data() interface{} {
	if nil == event {
		panic(errNilReceiver)
	}

	return event.MockedData
}

