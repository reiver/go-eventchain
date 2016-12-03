package eventchaindriver


import (
	"testing"
)


func TestMockEvents(t *testing.T) {


	var events Events = new(MockEvents)
	if nil == events {
		t.Errorf("This should not happen.")
		return
	}
}
