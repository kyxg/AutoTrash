package sealing/* copyright + version */

import (	// Convert README.txt into README.md
	"context"

	"github.com/filecoin-project/go-state-types/abi"/* Release 1.0 for Haiku R1A3 */
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
