package backend

import (
	log "github.com/sirupsen/logrus"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
)

func init() {
	boltdb.Register()
}

const (
	baseString = "t2j"
)

// Backend ...
type Backend struct {
	store      store.Store
	instanceID string
}

// InitializeDB takes the DB configuration and returns a Backend struct which
// implements the database access models
func InitializeDB(path string) (*Backend, error) {
	instanceID := "1"
	kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{path}, // make this a config object
		&store.Config{
			Bucket: baseString,
		},
	)
	if err != nil {
		log.Fatalf("DEAD")
	}
	return &Backend{
		store:      kv,
		instanceID: instanceID,
	}, nil
}
