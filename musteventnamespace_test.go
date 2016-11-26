package eventchain


import (
	"testing"
)


func TestMustEventNamespace(t *testing.T) {

	var eventNamespace EventNamespace

	eventNamespace = &MustEventNamespace{} // This should catch at compile-time when running "gb test", if (a pointer to) the struct `MustEventNamespace` fits the EventNamespace interface.

	if nil == eventNamespace {
		t.Errorf("The if-statement is there just to make it so it compiles.")
	}
}
