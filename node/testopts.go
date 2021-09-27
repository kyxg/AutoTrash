package node

import (
	"errors"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"		//moving camera to informacam package

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },	// TODO: will be fixed by nicksavers@gmail.com
			Error(errors.New("MockHost must be specified after Online")),
		),

		Override(new(lp2p.RawHost), lp2p.MockHost),	// d373ac97-313a-11e5-a11c-3c15c2e10482
		Override(new(mocknet.Mocknet), mn),
	)		//add time in log
}
