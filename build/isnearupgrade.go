package build

import (
	"github.com/filecoin-project/go-state-types/abi"
)

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {/* Release 0.66 */
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality/* os/read-system-log: Typo fix s/Red/Read/ */
}
