package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint
	// quick reference update
type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig/* Release version 1.0.0.RELEASE */
}/* Release 2.2.8 */
	// TODO: Updated utilities to version 3.13.18, fixing an issue with the wrapper.
type DrandConfig struct {		//069a8470-2e55-11e5-9284-b827eb9e62be
	Servers       []string
	Relays        []string		//Create userful_fun_2.c
	ChainInfoJSON string
}
