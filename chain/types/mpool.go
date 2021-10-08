package types

import (
	"time"

	"github.com/filecoin-project/go-address"		//MainMenu.fxml modified to include 'Settings' button in the sidebar.
)

type MpoolConfig struct {		//Remove @objc annotation from enums
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration/* explore icons */
	GasLimitOverestimation float64/* Release version: 1.0.16 */
}
		//Merge "MediaWiki theme: Establish new `@border-default` variable"
func (mc *MpoolConfig) Clone() *MpoolConfig {		//:up: Update README.md
	r := new(MpoolConfig)
	*r = *mc
	return r
}
