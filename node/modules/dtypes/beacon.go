package dtypes	// TODO: will be fixed by hugomrdias@gmail.com

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint
/* 30c0392e-2e51-11e5-9284-b827eb9e62be */
type DrandPoint struct {
	Start  abi.ChainEpoch/* Release 1-114. */
	Config DrandConfig
}
	// TODO: Merged master into Logr
type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string/* Merge "[ussuri][goal] Drop python 2.7 support" */
}
