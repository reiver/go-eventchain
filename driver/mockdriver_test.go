package eventchaindriver


import (
	"testing"
)


func TestMockDriver(t *testing.T) {


	var driver Driver = new(MockDriver)
	if nil == driver {
		t.Errorf("This should not happen.")
		return
	}
}
