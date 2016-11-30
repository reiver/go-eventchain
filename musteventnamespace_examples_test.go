package eventchain_test


import (
	"fmt"

	"github.com/reiver/go-eventchain"
)


func ExampleMustEventNamespace() {

	// Here we specify all the events that we want to exist
	// in our namespace.
	//
	// Note that we can NOT change this later.
	//
	// NOTE THAT THIS WILL panic() IF THERE IS AN ERROR!
	eventNamespace := eventchain.MustEventNamespace{(&eventchain.BasicEventNamespace{}).MustDeclare(
		"ATE_AN_APPLE",
		"ATE_A_BANANA",
		"ATE_A_CHERRY",
	)}


	// Here we are getting the "ATE_AN_APPLE" event name.
	//
	// NOTE THAT THIS WILL panic() IF THERE IS AN ERROR!
	ateAnApple := eventNamespace.MustEventName("ATE_AN_APPLE")

	// Here we are getting the "ATE_A_BANANA" event name.
	//
	// NOTE THAT THIS WILL panic() IF THERE IS AN ERROR!
	ateABanana := eventNamespace.MustEventName("ATE_A_BANANA")

	// Here we are getting the "ATE_A_CHERRY" event name.
	//
	// NOTE THAT THIS WILL panic() IF THERE IS AN ERROR!
	ateACherry := eventNamespace.MustEventName("ATE_A_CHERRY")

	fmt.Printf("%s\n", ateAnApple.String())
	fmt.Printf("%s\n", ateABanana.String())
	fmt.Printf("%s\n", ateACherry.String())

	// Output:
	// ATE_AN_APPLE
	// ATE_A_BANANA
	// ATE_A_CHERRY
}
