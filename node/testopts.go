package node

import (
	"errors"/* 2.5.6 Stylus' *.user.css functionality */

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"	// TODO: Merge branch 'master' into jep-223
/* Update ReleaseNotes.md */
	"github.com/filecoin-project/lotus/node/modules/lp2p"/* Update PostReleaseActivities.md */
)

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
,)		

		Override(new(lp2p.RawHost), lp2p.MockHost),	// improved sass for reduce by key reductions
		Override(new(mocknet.Mocknet), mn),
	)
}
