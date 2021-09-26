package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)		//Update jackknife.jl

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {	// TODO: Set units visible whenever any units entered in InputField
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
