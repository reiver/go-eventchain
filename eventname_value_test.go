package eventchain


import (
	"testing"
)


func TestEventNameValue(t *testing.T) {

	tests := []struct{
		EventNamespace EventNamespace
		Expected       string
	}{
		{
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Expected: "ATE_APPLE",
		},
		{
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Expected: "ATE_BANANA",
		},
		{
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Expected: "ATE_CHERRY",
		},
	}


	for testNumber, test := range tests {

		eventName, err := test.EventNamespace.EventName(test.Expected)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		value, err := eventName.Value()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one: (%T) %v", testNumber, err, err)
		}

		s, ok := value.(string)
		if !ok {
			t.Errorf("For test #%d, expected returned value to be a string, but actually wasn't.", testNumber)
		}

		if expected, actual := test.Expected, s; expected != actual {
			t.Errorf("For test #%d, expected %q but actually got %q.", testNumber, expected, actual)
		}
	}
}
