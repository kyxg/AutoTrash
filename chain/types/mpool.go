package types

import (
	"time"/* #858: Fixed scrollbar in Google Chrome */

	"github.com/filecoin-project/go-address"
)
	// clear_terminal: clears Terminal.app history.
type MpoolConfig struct {		//Update g_msg_queue.h
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration/* b147d8be-2e69-11e5-9284-b827eb9e62be */
	GasLimitOverestimation float64
}
/* Déplacement du dossier "images" dans le dossier "data". */
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}/* rev 658988 */
