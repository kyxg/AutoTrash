package dtypes

import "github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by alex.gaynor@gmail.com

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string/* - Fixed a blank file called "Plugins" being created when building */
	Relays        []string
	ChainInfoJSON string
}
