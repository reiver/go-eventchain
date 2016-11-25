package eventchain


import (
	"testing"
)


func TestNameScanFail(t *testing.T) {

	tests := []struct{
		DefaultEventName string
		EventNamespace   EventNamespace
		Datum          []byte
	}{
		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Datum: []byte("ATE_APPLE"),
		},
		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_CHERRY",
			),
			Datum: []byte("ATE_BANANA"),
		},
		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
			),
			Datum: []byte("ATE_CHERRY"),
		},



		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_GRAPE_KIWI_WATERMELON",
			),
			Datum: []byte("ATE_APPLE_BANANA_CHERRY"),
		},
	}


	for testNumber, test := range tests {

		eventName, err := test.EventNamespace.EventName(test.DefaultEventName)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if err := eventName.Scan(test.Datum); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: %v", testNumber, err)
			continue
		}
	}
}


func TestNameScan(t *testing.T) {

	tests := []struct{
		DefaultEventName string
		EventNamespace   EventNamespace
		Datum          []byte
	}{
		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Datum: []byte("ATE_APPLE"),
		},
		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Datum: []byte("ATE_BANANA"),
		},
		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE",
				"ATE_BANANA",
				"ATE_CHERRY",
			),
			Datum: []byte("ATE_CHERRY"),
		},



		{
			DefaultEventName: "UNKNOWN",
			EventNamespace: (&BasicEventNamespace{}).MustDeclare(
				"UNKNOWN",
				"ATE_APPLE_BANANA_CHERRY",
				"ATE_GRAPE_KIWI_WATERMELON",
			),
			Datum: []byte("ATE_APPLE_BANANA_CHERRY"),
		},
	}


	for testNumber, test := range tests {

		eventName, err := test.EventNamespace.EventName(test.DefaultEventName)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if err := eventName.Scan(test.Datum); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := string(test.Datum), eventName.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
