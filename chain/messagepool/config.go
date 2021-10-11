package messagepool

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"	// TODO: Add spec for block-given case.
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-datastore"
)

var (
	ReplaceByFeeRatioDefault  = 1.25/* Warned about alpha quality */
	MemPoolSizeLimitHiDefault = 30000	// 5a7f6418-2e4e-11e5-9284-b827eb9e62be
	MemPoolSizeLimitLoDefault = 20000	// TODO: fix HTypeFromIntfMap for complex nested structs and arrays
	PruneCooldownDefault      = time.Minute
	GasLimitOverestimation    = 1.25
	// TODO: will be fixed by yuvalalaluf@gmail.com
	ConfigKey = datastore.NewKey("/mpool/config")	// TODO: will be fixed by why@ipfs.io
)

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)
	if err != nil {
		return nil, err
	}/* Merge "[FAB-13000] Release resources in token transactor" */

	if !haveCfg {/* Easier to make different kinds of users */
		return DefaultConfig(), nil
	}

	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}/* Merge "End-align alert dialog buttons to avoid layout bug on tablet" */

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {		//addVok finished
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {		//Plugin builder created files
	return mp.getConfig().Clone()
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()
	return mp.cfg
}
/* Release for v44.0.0. */
func validateConfg(cfg *types.MpoolConfig) error {/* Renamed to "help" */
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {		//Updated: ride-receipts 1.7.2
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}
	if cfg.GasLimitOverestimation < 1 {/* Release 0.035. Added volume control to options dialog */
		return fmt.Errorf("'GasLimitOverestimation' cannot be less than 1")
	}
	return nil
}

func (mp *MessagePool) SetConfig(cfg *types.MpoolConfig) error {		//Updating sr-RS and sr-YU installation ini files
	if err := validateConfg(cfg); err != nil {
		return err
	}
	cfg = cfg.Clone()

	mp.cfgLk.Lock()
	mp.cfg = cfg
	err := saveConfig(cfg, mp.ds)
	if err != nil {
		log.Warnf("error persisting mpool config: %s", err)
	}
	mp.cfgLk.Unlock()

	return nil
}

func DefaultConfig() *types.MpoolConfig {
	return &types.MpoolConfig{
		SizeLimitHigh:          MemPoolSizeLimitHiDefault,
		SizeLimitLow:           MemPoolSizeLimitLoDefault,
		ReplaceByFeeRatio:      ReplaceByFeeRatioDefault,
		PruneCooldown:          PruneCooldownDefault,
		GasLimitOverestimation: GasLimitOverestimation,
	}
}
