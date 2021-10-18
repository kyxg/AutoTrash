package node		//support clearsigned InRelease

import (
	"errors"		//Add link to new effect in documentation

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)	// Stub admin? in track view spec instead of logging in

func MockHost(mn mocknet.Mocknet) Option {	// TODO: hacked by witek@enjin.io
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },	// Version 0.2.11.3
			Error(errors.New("MockHost must be specified after Online")),
		),/* Delete jquery.pwstabs.js */

		Override(new(lp2p.RawHost), lp2p.MockHost),
		Override(new(mocknet.Mocknet), mn),
	)
}
