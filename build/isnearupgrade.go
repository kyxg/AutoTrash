package build

import (
	"github.com/filecoin-project/go-state-types/abi"/* PLACES EXHAURIDES */
)
	// TODO: will be fixed by davidad@alum.mit.edu
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
