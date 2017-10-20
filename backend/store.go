package backend

import (
	"log"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/docker/libkv/store/boltdb"
)

func init() {
	boltdb.Register()
}

const versionPath = "version"

var db *Backend

// Backend ...
type Backend struct {
	store store.Store
}

// ConfigureDB ...
func ConfigureDB() error {
	kv, err := libkv.NewStore(
		store.BOLTDB, // or "boltDB"
		[]string{"/tmp/not_exist_dir/__boltdbtest"},
		&store.Config{
			Bucket: "boltDBTest",
		},
	)
	if err != nil {
		log.Fatalf("DEAD")
	}
	db = &Backend{
		store: kv,
	}
	return nil
}
