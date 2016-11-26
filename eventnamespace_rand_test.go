package eventchain


import (
	"fmt"

	"testing"
)


func TestBasicEventNamespaceRandNotFound(t *testing.T) {

	var eventNamespace EventNamespace = (&BasicEventNamespace{}).MustDeclare(/* Nothing here */)

	eventName, err := eventNamespace.Rand()
	if nil == err {
		t.Errorf("Expected to get an error, but did not actually get one: %v. Got this event name: %v", err, eventName)
		return
	}
	if _, ok := err.(NotFoundComplainer); !ok {
		t.Errorf("Expected to get Not Found error, but actually got something else: (%T) %v", err, err)
		return
	}
}


func TestBasicEventNamespaceRand(t *testing.T) {

	for testNumber:=0; testNumber<30; testNumber++ {

		var eventNames []string
		{
			limit := 1+testNumber

			for ii:=0; ii<limit; ii++ {
				eventName := fmt.Sprintf("RAND_EVENT_NAME_%d_%d", limit, ii)

				eventNames = append(eventNames, eventName)
			}
		}

		var eventNamespace EventNamespace = MustEventNamespace{(&BasicEventNamespace{}).MustDeclare(eventNames...)}

		eventName, err := eventNamespace.Rand()
		if nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		{
			actual := eventName.String()

			found := false

			for _, expected := range eventNames {
				if expected == actual {
					found = true
				}
			}

			if !found {
				t.Errorf("For test #%d, could not find event name %q in: %v", testNumber, actual, eventNames)
				continue
			}
		}
	}

}
