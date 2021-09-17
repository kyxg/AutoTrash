package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint

type DrandPoint struct {
	Start  abi.ChainEpoch	// Add a line to Ignore ".idea/compiler.xml"
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string	// TODO: hacked by xaber.twt@gmail.com
	Relays        []string
	ChainInfoJSON string
}
