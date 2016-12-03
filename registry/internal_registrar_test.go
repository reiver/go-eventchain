package eventchainregistry


import (
	"github.com/reiver/go-eventchain/driver"

	"testing"
)


func TestInternalRegistrar(t *testing.T) {

	tests := []struct{
		NamedDrivers []struct{
			Driver eventchaindriver.Driver
			Name   string
		}
	}{
		{
			NamedDrivers: []struct{
				Driver eventchaindriver.Driver
				Name   string
			}{
				{
					Driver: new(eventchaindriver.MockDriver),
					Name:   "apple",
				},
				{
					Driver: new(eventchaindriver.MockDriver),
					Name:   "banana",
				},
				{
					Driver: new(eventchaindriver.MockDriver),
					Name:   "cherry",
				},
			},
		},
	}


	for testNumber, test := range tests {

		iRegistrar := new(internalRegistrar)

		var registry Registrar = iRegistrar

		for driverNumber, namedDriver := range test.NamedDrivers {
			driver, err := registry.Driver(namedDriver.Name)
			if nil == err {
				t.Errorf("For test #%d and driver #%d, expected an error, but did not get one: %v", testNumber, driverNumber, err)
				continue
			}
			if nil != driver {
				t.Errorf("For test #%d and driver #%d, expected nil, but actually wasn't: (%T) %v", testNumber, driverNumber, driver, driver)
				continue
			}
			if expected, actual := errNotFound, err; expected != actual {
				t.Errorf("For test #%d and driver #%d, expected (%T) %v, but actually got (%T) %v", testNumber, driverNumber, expected, expected, actual, actual)
				continue
			}
		}

		for driverNumber, namedDriver := range test.NamedDrivers {

			if err := registry.Register(namedDriver.Driver, namedDriver.Name); nil != err {
				t.Errorf("For test #%d and driver #%d, did not expect an error, but actually got one: (%T) %v", testNumber, driverNumber, err, err)
				continue
			}
		}

		for driverNumber, namedDriver := range test.NamedDrivers {
			driver, err := registry.Driver(namedDriver.Name)
			if nil != err {
				t.Errorf("For test #%d and driver #%d, did not expect an error, but acutally got one: (%T) %v", testNumber, driverNumber, err, err)
				continue
			}
			if nil == driver {
				t.Errorf("For test #%d and driver #%d, did not expect nil, but actually got: %v", testNumber, driverNumber, driver)
				continue
			}
			if expected, actual := namedDriver.Driver, driver; expected != actual {
				t.Errorf("For test #%d and driver #%d, expected %p, but actually got %p", testNumber, driverNumber, expected, actual)
				continue
			}
		}
	}
}
