package backend

import (
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"

	"fmt"
)

func init() {
	// this needs to be done to initialize the lower level data stores that we are supporting
	// for now we are only going to support BoltDB in order to keep it simple. Others would be added
	// here and then handled during the initialization function
	boltdb.Register()
}

const (
	// boltBaseString is to ensure we can exist in a multi-tenant K/V store environment. We don't want
	// to conflict with other services with similar keys
	boltBaseString = "t2j"
)

// Backend interface encapsulates all the functions to put/get domain objects
type Database interface {
	UserBackend
	TaskBackend
	ProjectBackend
}

// Backend ...
type Backend struct {
	// this is the libkv reference that we are wrapping with our domain objects
	store store.Store
	// instanceID is storing the unique ID of this specific instance so we can support multiple t2j
	// instances in the same datastore
	instanceID string
}

// InitializeDB takes the DB configuration and returns a Backend struct which
// implements the database access models
func InitializeDB(path string) (*Backend, error) {
	// TODO: make this a passed in value from configuration or add some other mechanism to ensure uniqueness
	instanceID := "1"
	// we only support Bolt to start but if we wanted other we would need to create a switch or if/else to
	// properly setup the configuration for the right backend store
	kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{path}, // make this a config object
		&store.Config{
			Bucket: boltBaseString,
		},
	)
	// if we get an error, we need to blow up but let the higher levels handle this
	if err != nil {
		return &Backend{}, fmt.Errorf("Failure to create datastore: %+v", err)
	}
	// if we got here, we are good.
	return &Backend{
		store:      kv,
		instanceID: instanceID,
	}, nil
}
