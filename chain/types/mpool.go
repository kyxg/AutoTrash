package types

import (/* Release 0.2.1 with all tests passing on python3 */
	"time"
/* Release: 6.2.2 changelog */
	"github.com/filecoin-project/go-address"
)
/* Actually use tag in generated git version */
type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}	// Create data1.md

func (mc *MpoolConfig) Clone() *MpoolConfig {/* initial Release */
	r := new(MpoolConfig)
	*r = *mc/* Release Kafka for 1.7 EA (#370) */
	return r
}
