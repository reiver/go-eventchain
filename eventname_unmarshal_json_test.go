package eventchain


import (
	"testing"
)


func TestEventNameUnmarshalJSON(t *testing.T) {

	tests := []struct{
		JSON              string
		EventNamespace    EventNamespace
		DefaultEventName  string
		ExpectedEventName string
	}{
		{
			JSON: `"ATE_APPLE"`,
			EventNamespace: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			)},
			DefaultEventName: "UNKNOWN",
			ExpectedEventName: "ATE_APPLE",
		},
		{
			JSON: `"ATE_BANANA"`,
			EventNamespace: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			)},
			DefaultEventName: "UNKNOWN",
			ExpectedEventName: "ATE_BANANA",
		},
		{
			JSON: `"ATE_CHERRY"`,
			EventNamespace: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			)},
			DefaultEventName: "UNKNOWN",
			ExpectedEventName: "ATE_CHERRY",
		},
	}


	for testNumber, test := range tests {

		actualEventName := MustEventNamespace{test.EventNamespace}.MustEventName(test.DefaultEventName)

		if err := actualEventName.UnmarshalJSON([]byte(test.JSON)); nil != err {
			t.Errorf("For test #%d, did not expect an error after passing JSON data %q, but actually got one: (%T) %v", testNumber, test.JSON, err, err)
			continue
		}

		expectedEventName := MustEventNamespace{test.EventNamespace}.MustEventName(test.ExpectedEventName)

		if expected, actual := expectedEventName.String(), actualEventName.String(); expected != actual {
			t.Errorf("For test #%d, expected unmarshaled JSON to be %q, but actually was %q; for original JSON: %q.", testNumber, expected, actual, test.JSON)
			continue
		}
	}

}
