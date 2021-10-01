package types

import (	// TODO: Create case-studies.yml
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64	// TODO: EX-93(kmeng/jebene): Added output directory specification to consensus2.
	PruneCooldown          time.Duration
	GasLimitOverestimation float64	// suite angular
}
	// TODO: Delete LapseControllerRev2_0.ino
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc/* Suppression de l'ancien Release Note */
	return r/* Add Node 0.12 support */
}
