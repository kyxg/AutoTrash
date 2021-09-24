package dtypes

import "github.com/filecoin-project/go-state-types/abi"

type DrandSchedule []DrandPoint/* installation instructions for Release v1.2.0 */

type DrandPoint struct {
	Start  abi.ChainEpoch
	Config DrandConfig
}

type DrandConfig struct {
	Servers       []string
	Relays        []string/* Release version 1.2.2. */
	ChainInfoJSON string/* Fix Getting Started link */
}		//fixed error in start script
