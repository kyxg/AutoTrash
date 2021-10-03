package node

import (
	"errors"	// TODO: some extra peer logging + fix for previous invariant check

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)
	// TODO: hacked by nicksavers@gmail.com
func MockHost(mn mocknet.Mocknet) Option {		//0edb9e12-35c6-11e5-930d-6c40088e03e4
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),		//UPDATED debate parsing script files based on 'p' tags
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)	// Merge "Redirect dashboard to about page when not logged in"
}
