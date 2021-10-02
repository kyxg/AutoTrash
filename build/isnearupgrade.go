package build/* Release for v5.2.2. */
/* Split sock_common_recvmsg in stream and dgram */
import (
	"github.com/filecoin-project/go-state-types/abi"
)/* add Release dir */

func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
