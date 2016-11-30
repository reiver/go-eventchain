package eventchain


import (
	"math/rand"
	"time"
)


var (
	randomness = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
)


// Rand returns a random EventName from the namespace.
func (ns *BasicEventNamespace) Rand() (EventName, error) {
	if nil == ns {
		return nil, errNilReceiver
	}

	names := ns.names
	if nil == names {
		return nil, errInternalError
	}

	lenNames := len(names)
	if 1 > lenNames {
		return nil, errNotFound
	}

	n := randomness.Intn(lenNames)

	i := 0
	for k,_ := range names {
		if i == n {
			return &internalEventName{
				s: k,
				eventNamespace: ns,
			}, nil
		}

		i++
	}

	return nil, errInternalError
}
