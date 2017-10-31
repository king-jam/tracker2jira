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

// DB ...
var DB *Backend

// Backend ...
type Backend struct {
	store      store.Store
	instanceID string
}

// GetDB ...
func GetDB() (*Backend, error) {
	if DB != nil {
		return DB, nil
	}
	instanceID := "1"
	kv, err := libkv.NewStore(
		store.BOLTDB,
		[]string{"/tmp/not_exist_dir/__boltdbtest"}, // make this a config object
		&store.Config{
			Bucket: baseString,
		},
	)
	if err != nil {
		log.Fatalf("DEAD")
	}
	DB = &Backend{
		store:      kv,
		instanceID: instanceID,
	}
	return DB, nil
}
