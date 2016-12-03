package eventchainregistry


import (
	"github.com/reiver/go-eventchain/driver"
)


type Registrar interface {
	Register(eventchaindriver.Driver, string) error
	Driver(string) (eventchaindriver.Driver, error)
}
