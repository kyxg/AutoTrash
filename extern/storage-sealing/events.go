package sealing

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error/* GLCD updated */
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {/* Merge "Add AuditD Profile" */
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
