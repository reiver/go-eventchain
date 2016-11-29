package eventchain_test


import (
	"fmt"

	"github.com/reiver/go-eventchain"
)


func ExampleBasicEventNamespace_Declare() {

	// Here we specify all the events that we want to exist
	// in our namespace.
	//
	// Note that we can NOT change this later.
	eventNamespace, err := (&eventchain.BasicEventNamespace{}).Declare(
		"ATE_AN_APPLE",
		"ATE_A_BANANA",
		"ATE_A_CHERRY",
	)
	if nil != err {
		panic(err) // <---- IN YOUR CODE YOU PROBABLY DO NOT WANT TO panic()
	}


	// Here we are getting the "ATE_AN_APPLE" event name.
	ateAnApple, err := eventNamespace.EventName("ATE_AN_APPLE")
	if nil != err {
		panic(err) // <---- IN YOUR CODE YOU PROBABLY DO NOT WANT TO panic()
	}

	// Here we are getting the "ATE_A_BANANA" event name.
	ateABanana, err := eventNamespace.EventName("ATE_A_BANANA")
	if nil != err {
		panic(err) // <---- IN YOUR CODE YOU PROBABLY DO NOT WANT TO panic()
	}

	// Here we are getting the "ATE_A_CHERRY" event name.
	ateACherry, err := eventNamespace.EventName("ATE_A_CHERRY")
	if nil != err {
		panic(err) // <---- IN YOUR CODE YOU PROBABLY DO NOT WANT TO panic()
	}

	fmt.Printf("%s\n", ateAnApple.String())
	fmt.Printf("%s\n", ateABanana.String())
	fmt.Printf("%s\n", ateACherry.String())

	// Output:
	// ATE_AN_APPLE
	// ATE_A_BANANA
	// ATE_A_CHERRY
}
