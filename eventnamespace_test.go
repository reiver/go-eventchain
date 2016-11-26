package eventchain


import (
	"testing"
)


func TestBasicEventNamespaceDeclareFailOnTwice(t *testing.T) {

	var basicEventNamespace BasicEventNamespace

	_, err := basicEventNamespace.Declare(
		"ATE_APPLE",
		"ATE_BANANA",
		"ATE_CHERRY",
	)
	if nil != err {
		t.Errorf("Did not expect an error, but actually got one: (%T) %v", err, err)
		return
	}

	_, err = basicEventNamespace.Declare(
		"ATE_GRAPE",
		"ATE_KIWI",
		"ATE_WATERMELON",
	)
	if nil == err {
		t.Errorf("Expected an error, but did not actually got one: %v", err)
		return
	}

}


func TestBasicEventNamespace(t *testing.T) {

	var eventNamespace EventNamespace

	eventNamespace = &BasicEventNamespace{} // This should catch at compile-time when running "gb test", if (a pointer to) the struct `BasicEventNamespace` fits the EventNamespace interface.

	if nil == eventNamespace {
		t.Errorf("The if-statement is there just to make it so it compiles.")
	}
}
