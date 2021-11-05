package types

import (
	"time"/* :zap:How to use JS APIs answers now updated */

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration/* Created sample-payload.json */
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)	// 42c2bbfc-2e41-11e5-9284-b827eb9e62be
	*r = *mc
	return r
}
