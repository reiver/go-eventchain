package eventchain


import (
	"errors"
)


var (
	errBadRequest         = errors.New("Bad Reqest")
	errBadRequestNilEvent = errors.New("Bad Request: Nil Event")
	errInternalError      = errors.New("Internal Error")
	errNotFound           = internalNotFoundComplainer{}
	errNilReceiver        = errors.New("Nil Receiver")
)
