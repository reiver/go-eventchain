package eventchainregistry


import (
	"errors"
)


var (
	errConflict    = errors.New("Conflict")
	errNilDriver   = errors.New("Nil Driver")
	errNilReceiver = errors.New("Nil Receiver")
	errNotFound    = errors.New("Not Found")
)
