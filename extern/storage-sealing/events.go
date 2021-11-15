package sealing
/* Add OSU multi latency test in demos */
import (
	"context"/* Issue #44 Fixed append location bug on Journal recovery. */
/* #19 GIBS-542 Added support for classifications with horizontal legends  */
	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error	// TODO: fix n error when loading the station entries
}
