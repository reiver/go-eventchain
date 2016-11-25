package eventchain


import (
	"errors"
)


var (
	errInternalError = errors.New("Internal Error")
	errNotFound      = internalNotFoundComplainer{}
	errNilReceiver   = errors.New("Nil Receiver")
)
