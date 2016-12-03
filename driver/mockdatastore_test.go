package eventchaindriver


import (
	"testing"
)


func TestMockDataStore(t *testing.T) {


	var dataStore DataStore = new(MockDataStore)
	if nil == dataStore {
		t.Errorf("This should not happen.")
		return
	}
}
