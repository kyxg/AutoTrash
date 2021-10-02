package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {		//Added info on the IRremote library being mocked
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}
