package eventchain


import (
	"database/sql"
)


type DataStore interface {
	OpenEventStore(string) (EventStore, error)
}


type internalDataStore struct {
	db *sql.DB
}

func OpenDataStore(driverSpec string, args ...interface{}) (DataStore, error) {

	const onlySupportedDriverSpec = "database/sql:postgres"

	if "" == driverSpec {
		return nil, errNotFound
	}

	if onlySupportedDriverSpec != driverSpec {
		return nil, errNotFound
	}

	lenArgs := len(args)
	if 1 > lenArgs {
		return nil, errBadRequest
	}

	arg0 := args[0]

	db, ok := arg0.(*sql.DB)
	if !ok {
		return nil, errBadRequest
	}

	dataStore := internalDataStore{
		db: db,
	}

	return &dataStore, nil
}
