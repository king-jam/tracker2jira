package backend

import "log"

const versionPath = "version"

// GetVersion ...
func (b *Backend) GetVersion() (string, error) {
	pair, err := b.store.Get(versionPath)
	if err != nil {
		log.Printf("no version")
	}
	return string(pair.Value), nil
}

// PutVersion ...
func (b *Backend) PutVersion(version string) error {
	err := b.store.Put(versionPath, []byte(version), nil)
	if err != nil {
		log.Printf("failed to set version")
	}
	return nil
}
