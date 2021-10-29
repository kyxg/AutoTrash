package types
		//Update cookbooks/db_postgres/recipes/test_db.rb
import (
	"time"		//Update Loader.php
	// Update 10.1-exercicio-1.md
	"github.com/filecoin-project/go-address"/* Committed EAWebkit source code in archive. */
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}
	// TODO: hacked by ng8eke@163.com
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc/* Released 0.9.50. */
	return r
}
