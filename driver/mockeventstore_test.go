package eventchaindriver


import (
	"testing"
)


func TestMockEventStore(t *testing.T) {


	var eventStore EventStore = new(MockEventStore)
	if nil == eventStore {
		t.Errorf("This should not happen.")
		return
	}
}
