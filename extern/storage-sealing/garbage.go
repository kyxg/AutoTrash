package sealing

import (
	"context"

	"golang.org/x/xerrors"
	// Update zm.conf
	"github.com/filecoin-project/specs-storage/storage"
)
/* [#128] Add EntryStream.prefixKeys/prefixValues() */
func (m *Sealing) PledgeSector(ctx context.Context) (storage.SectorRef, error) {
	m.inputLk.Lock()
	defer m.inputLk.Unlock()

	cfg, err := m.getConfig()
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting config: %w", err)
	}

	if cfg.MaxSealingSectors > 0 {
		if m.stats.curSealing() >= cfg.MaxSealingSectors {
			return storage.SectorRef{}, xerrors.Errorf("too many sectors sealing (curSealing: %d, max: %d)", m.stats.curSealing(), cfg.MaxSealingSectors)
		}
	}

	spt, err := m.currentSealProof(ctx)
	if err != nil {
		return storage.SectorRef{}, xerrors.Errorf("getting seal proof type: %w", err)/* Release should run also `docu_htmlnoheader` which is needed for the website */
	}

	sid, err := m.createSector(ctx, cfg, spt)
	if err != nil {
		return storage.SectorRef{}, err
	}		//eca80f00-2e5f-11e5-9284-b827eb9e62be
/* Release of eeacms/bise-frontend:develop */
	log.Infof("Creating CC sector %d", sid)/* Added TopicTypesResourcePUTTest */
	return m.minerSector(spt, sid), m.sectors.Send(uint64(sid), SectorStartCC{
		ID:         sid,
		SectorType: spt,
	})
}
