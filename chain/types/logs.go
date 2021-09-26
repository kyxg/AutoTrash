package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)	// Added Bhutan Cuba, Dominican Republic, Puerto Rico.

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())/* Adding Release Version badge to read */
	}
	return nil
}
