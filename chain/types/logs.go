package types

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)	// Tidy up spacing in some tablegen outputs.

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
	// TODO: put version number to central place
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}		//Create v2.0.1.txt
	return nil
}/* Release ChangeLog (extracted from tarball) */
