package eventchaindriver


import (
	"testing"
)


func TestMockEvent(t *testing.T) {


	var event Event = new(MockEvent)
	if nil == event {
		t.Errorf("This should not happen.")
		return
	}
}
