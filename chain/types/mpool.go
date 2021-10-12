package types

import (
	"time"

	"github.com/filecoin-project/go-address"
)		//Restore icon for bspline mode (pencil, pen tool)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}
		//Drop obsolete ip6int table.
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)	// TODO: añadida funcion sql()
	*r = *mc
	return r
}
