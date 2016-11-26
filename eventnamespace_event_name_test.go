package eventchain


import (
	"testing"
)


func TestBasicEventNameEventName(t *testing.T) {

	tests := []struct{
		EventNames []string
		EventName   string
	}{
		{
			EventNames: []string{
				"ATE_AN_APPLE",
				"ATE_A_BANANA",
				"ATE_A_CHERRY",
			},
			EventName: "ATE_AN_APPLE",
		},
		{
			EventNames: []string{
				"ATE_AN_APPLE",
				"ATE_A_BANANA",
				"ATE_A_CHERRY",
			},
			EventName: "ATE_A_BANANA",
		},
		{
			EventNames: []string{
				"ATE_AN_APPLE",
				"ATE_A_BANANA",
				"ATE_A_CHERRY",
			},
			EventName: "ATE_A_CHERRY",
		},
	}


	for testNumber, test := range tests {

		var eventNamespace EventNamespace = (&BasicEventNamespace{}).MustDeclare(test.EventNames...)

		eventName, err := eventNamespace.EventName(test.EventName)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v; for event name %q, and event names %v", testNumber, err, err, test.EventName, test.EventNames)
			continue
		}

		if expected, actual := test.EventName, eventName.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
