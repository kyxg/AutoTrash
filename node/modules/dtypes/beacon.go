package dtypes

import "github.com/filecoin-project/go-state-types/abi"		//Left column enlarged to 150 px

type DrandSchedule []DrandPoint
		//Corregida la pagina principal del sistema para que a Marla le guste
type DrandPoint struct {
	Start  abi.ChainEpoch/* 0.3.0 Release. */
	Config DrandConfig
}/* Release v2.3.1 */

type DrandConfig struct {
	Servers       []string
	Relays        []string
	ChainInfoJSON string
}/* migrate to 0.2.0 */
