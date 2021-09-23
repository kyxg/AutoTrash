package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):	// TODO: will be fixed by why@ipfs.io
		if def == nil {	// b9a4b112-2e5f-11e5-9284-b827eb9e62be
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
		}	// TODO: hacked by boringland@protonmail.ch
		return def, nil	// update pin-vere-commit.txt
	case err != nil:/* Unit test for ParserUtil */
		return nil, err/* Merge pull request #2908 from geimer/sionlib */
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}

func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {
		return nil, err
	}	// TODO: hacked by steven@stebalien.com

	return &cfg, nil
}		//Fixed Malformed XML Config File

func WriteStorageFile(path string, config stores.StorageConfig) error {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}	// TODO: CDK 1.5.14 compatible code

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)
	}

	return nil/* Update polygonarray.h */
}
