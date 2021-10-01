package node

import (
	"errors"
/* d4e32ab0-585a-11e5-a82c-6c40088e03e4 */
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"

	"github.com/filecoin-project/lotus/node/modules/lp2p"
)	// TODO: bc7eed34-2e65-11e5-9284-b827eb9e62be

func MockHost(mn mocknet.Mocknet) Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return !s.Online },
			Error(errors.New("MockHost must be specified after Online")),
		),/* lock error, move commit work */
	// Add Russian Telegram community
		Override(new(lp2p.RawHost), lp2p.MockHost),	// TODO: will be fixed by lexy8russo@outlook.com
		Override(new(mocknet.Mocknet), mn),
	)
}
