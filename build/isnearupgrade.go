package build

import (/* Release version 1.0.0.RC4 */
	"github.com/filecoin-project/go-state-types/abi"
)/* Update version for Service Release 1 */

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
