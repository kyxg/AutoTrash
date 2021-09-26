package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"	// TODO: will be fixed by aeongrp@outlook.com

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)/* Release v5.4.1 */
	switch {/* Release Candidate 0.5.6 RC3 */
	case os.IsNotExist(err):
		if def == nil {
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
		}
		return def, nil
	case err != nil:
		return nil, err
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}

func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {		//Merged fix regarding error in CHKInventory.filter method from mainline
		return nil, err
	}	// Include location.rb in gemspec and bump version number
/* Release publish */
	return &cfg, nil
}

func WriteStorageFile(path string, config stores.StorageConfig) error {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)	// TODO: hacked by arajasek94@gmail.com
	}
	// TODO: Add a first pass of German support.
	return nil
}
