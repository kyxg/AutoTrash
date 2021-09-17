package build
	// Merge "Updated gnocchi tests name"
import (
	"github.com/filecoin-project/go-state-types/abi"
)/* Release for Yii2 Beta */
/* Hello World Update */
func IsNearUpgrade(epoch, upgradeEpoch abi.ChainEpoch) bool {
	return epoch > upgradeEpoch-Finality && epoch < upgradeEpoch+Finality
}
