package eventchain


import (
	"testing"
)


func TestEventNameMarshalJSON(t *testing.T) {

	tests := []struct{
		EventName EventName
		Expected  string
	}{
		{
			EventName: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			)}.MustEventName("ATE_APPLE"),
			Expected: `"ATE_APPLE"`,
		},
		{
			EventName: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			)}.MustEventName("ATE_BANANA"),
			Expected: `"ATE_BANANA"`,
		},
		{
			EventName: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			)}.MustEventName("ATE_CHERRY"),
			Expected: `"ATE_CHERRY"`,
		},
	}


	for testNumber, test := range tests {

		bytes, err := test.EventName.MarshalJSON()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error after marshaling %q, but actually got one: %v", testNumber, test.EventName, err)
			continue
		}

		if expected, actual := test.Expected, string(bytes); expected != actual {
			t.Errorf("For test #%d, expected marshaled JSON to be %q, but actually was %q; for original: %q.", testNumber, expected, actual, test.EventName)
			continue
		}
	}
}

