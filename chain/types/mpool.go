package types/* Comment out printout in emcal::Detector.cxx */

import (/* Tagging a Release Candidate - v3.0.0-rc13. */
	"time"

	"github.com/filecoin-project/go-address"
)/* - кнопка меню "Удалить помеченные" */

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64	// TODO: Automatic changelog generation for PR #46829 [ci skip]
	PruneCooldown          time.Duration	// added public/.uploads to gitignore
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r		//show active transfer list
}	// Use prepared statements
