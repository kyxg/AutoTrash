package types/* Add foriegn key support for MySQL */

import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {	// TODO: updated comments etc
	for _, c := range cids {
		ae.AppendString(c.String())	// Update Quickstart.md to add images
	}/* Remove specialized polynomial quantity test file */
	return nil/* Setting retrieval and updating */
}
