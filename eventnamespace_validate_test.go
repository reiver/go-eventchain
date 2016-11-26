package eventchain


import (
	"testing"
)


func TestBasicEventNamespaceValidate(t *testing.T) {

	tests := []struct{
		EventNamespace      EventNamespace
		ValidEventNames   []string
		InvalidEventNames []string

	}{
		{
			EventNamespace: MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(
				"ATE_AN_APPLE",
				"ATE_A_BANANA",
				"ATE_A_CHERRY",
			)},
			ValidEventNames: []string{
				"ATE_AN_APPLE",
				"ATE_A_BANANA",
				"ATE_A_CHERRY",
			},
			InvalidEventNames: []string{
				"ATE_A_GRAPE",
				"ATE_A_KIWI",
				"ATE_A_WATERMELON",
			},
		},
	}


	for testNumber, test := range tests {

		for validEventNameNumber, validEventName := range test.ValidEventNames {
			if err := test.EventNamespace.Validate(validEventName); nil != err {
				t.Errorf("For test #%d and valid event name #%d, did not expect an error, but actually got one: (%T) %v", testNumber, validEventNameNumber, err, err)
				continue
			}
		}

		for invalidEventNameNumber, invalidEventName := range test.InvalidEventNames {
			err := test.EventNamespace.Validate(invalidEventName)
			if nil == err {
				t.Errorf("For test #%d and invalid event name #%d, expected an error, but did not actually get one: %v", testNumber, invalidEventNameNumber, err)
				continue
			}
			if _, ok := err.(NotFoundComplainer); !ok {
				t.Errorf("For test #%d and invalid event name #%d, expected an Not Found error, but did not actually get that: (%T) %v", testNumber, invalidEventNameNumber, err, err)
				continue
			}
		}

	}
}
