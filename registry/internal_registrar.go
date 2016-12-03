package eventchainregistry


import (
	"github.com/reiver/go-eventchain/driver"
)


type internalRegistrar struct{
	m map[string]eventchaindriver.Driver

}


func (r *internalRegistrar) Register(driver eventchaindriver.Driver, name string) error {
	if nil == r {
		return errNilReceiver
	}

	if nil == r.m {
		r.m = map[string]eventchaindriver.Driver{}
	}

	if _, exists := r.m[name]; exists {
		return errConflict
	}

	if nil == driver {
		return errNilDriver
	}

	r.m[name] = driver

	return nil
}


func (r *internalRegistrar) Driver(name string) (eventchaindriver.Driver, error) {
	if nil == r {
		return nil, errNilReceiver
	}

	m := r.m
	if nil == m {
		return nil, errNotFound
	}

	driver, found := m[name]
	if ! found {
		return nil, errNotFound
	}

	if nil == driver {
		return nil, errNilDriver
	}

	return driver, nil
}
