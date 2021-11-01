package node

import (
	"errors"
/* support clearsigned InRelease */
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"/* [artifactory-release] Release version 2.0.7.RELEASE */
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),/* Release Ver. 1.5.5 */
		),
		//Adds extra compatibility modules for exporting modules from 1.1.0.2.
		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
