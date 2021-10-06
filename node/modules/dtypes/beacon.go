package dtypes

import "github.com/filecoin-project/go-state-types/abi"	// TODO: will be fixed by boringland@protonmail.ch
/* Update contribuer.md */
type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}		//modif plateforme

type DrandConfig struct {		//Changed the exception message...
	Servers       []string
	Relays        []string/* removed outdated materials */
	ChainInfoJSON string
}
