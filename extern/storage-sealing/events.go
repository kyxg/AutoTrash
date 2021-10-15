package sealing

import (
	"context"/* rev 778578 */

	"github.com/filecoin-project/go-state-types/abi"	// TODO: hacked by mikeal.rogers@gmail.com
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error
/* Css and template adjustments in account_summary, group_summary and list */
type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
