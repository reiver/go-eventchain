package eventchain


import (
	"fmt"
)


// EventNamespace represents a set of event names.
//
// To get an event name value, one would use code like the follow:
//
//	eventName, err := eventNamespace.EventName("ATE_A_CHERRY")
//
// But before you can do that, you need to create an EventNamespace.
//
// Likely one would use BasicEventNamespace to create a EventNamespace.
//
// For example:
//
//	var (
//		EventNamespace EventNamespace
//	)
//	
//	func init() {
//		var basicEventNamespace BasicEventNamespace
//	
//		// Here we specify what events exist.
//		eventNamespace, err := basicEventNamespace.Declare(
//			"ATE_AN_APPLE",
//			"ATE_A_BANANA",
//			"ATE_A_CHERRY",
//		)
//		if nil != err {
//			panic(err) // <---- IN YOUR CODE YOU MIGHT NOT WANT TO panic()
//		}
//	
//		EventNamespace = eventNamespace
//	}
//
// In that example, our event namespace has 3 events: ATE_AN_APPLE, ATE_A_BANANA, ATE_A_CHERRY.
type EventNamespace interface {
	EventName(string) (EventName, error)
	Rand() (EventName, error)
	Validate(string) error
}


// BasicEventNamespace is used to create a EventNamespace.
//
// For example:
//
//	var (
//		EventNamespace EventNamespace
//	)
//	
//	func init() {
//		var basicEventNamespace BasicEventNamespace
//	
//		// Here we specify what events exist.
//		eventNamespace, err := basicEventNamespace.Declare(
//			"ATE_AN_APPLE",
//			"ATE_A_BANANA",
//			"ATE_A_CHERRY",
//		)
//		if nil != err {
//			panic(err) // <---- IN YOUR CODE YOU MIGHT NOT WANT TO panic()
//		}
//	
//		EventNamespace = eventNamespace
//	}
//
// In that example, our event namespace has 3 events: ATE_AN_APPLE, ATE_A_BANANA, ATE_A_CHERRY.
type BasicEventNamespace struct {
	names map[string]struct{}
}


// MustDeclare is like Declare, except it panic()s on error.
//
// Example usage:
//
//	import (
//		"github.com/reiver/go-eventchain"
//	)
//	
//	var (
//		RestaurantEventNames EventNamespace = (&BasicNameSpace{}).MustDeclare(
//			"FOOD_ORDERED",
//			"FOOD_COOKED",
//			"FOOD_DELIVERED",
//		)
//	)
func (ns *BasicEventNamespace) MustDeclare(names ...string) EventNamespace {
	namespace, err := ns.Declare(names...)
	if nil != err {
		panic(err)
	}

	return namespace
}


// Declare initializes a namespace, by declaring the names that exist in it.
//
// Note that Declare only "allows" you to specify the names in the namespace
// once.
//
// On all invocations after the first one, it will return an error.
func (ns *BasicEventNamespace) Declare(names ...string) (EventNamespace, error) {
	if nil == ns {
		return nil, errNilReceiver
	}

	if previousNames := ns.names; nil != previousNames {
		return nil, fmt.Errorf("You may not declare the names in the namespace more than once! %T already declared, with names: %v. Cannot re-declare with names: %v", *ns, previousNames, names)
	}

	ns.names = map[string]struct{}{}
	for _, name := range names {
		ns.names[name] = struct{}{}
	}

	return ns, nil
}
