package sealing
	// TODO: will be fixed by hugomrdias@gmail.com
import (
	"context"
		//Make Dummy class description more verbose
	"github.com/filecoin-project/go-state-types/abi"	// TODO: custom i18n for extjs
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error/* Allow any origin while in development mode */
type RevertHandler func(ctx context.Context, tok TipSetToken) error/* First Release , Alpha  */

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
