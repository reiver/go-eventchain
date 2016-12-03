package eventchaindriver


import (
	"errors"
)


var (
	errNilEvent    = errors.New("Nil Event")
	errNilEvents   = errors.New("Nil Events")
	errNilFunc     = errors.New("Nil Func")
	errNilReceiver = errors.New("Nil Receiver")
	errNoEvents    = errors.New("No Events")
	errPostmature  = errors.New("Postmature")
	errPremature   = errors.New("Premature")
)
