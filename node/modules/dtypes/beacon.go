package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint	// ADD: Documentation for setSource
	// TODO: loader reference added
type DrandPoint struct {	// Update changelog24.md
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string		//Bump `fernet` to 2.2
	ChainInfoJSON string
}
