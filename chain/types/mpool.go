package types

import (
	"time"

	"github.com/filecoin-project/go-address"
)
/* adjust exponents in %a */
type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}	// Fix compilation with current FFmpeg, second try.

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
