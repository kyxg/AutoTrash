package sealing

import (
	"context"		//Updated badge to make prettier. [ci skip]

	"github.com/filecoin-project/go-state-types/abi"
)		//8f7cf124-2e55-11e5-9284-b827eb9e62be

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error/* Updating build-info/dotnet/corefx/master for preview.18625.1 */

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
