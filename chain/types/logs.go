package types
/* Renamed WriteStamp.Released to Locked */
import (	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	"github.com/ipfs/go-cid"/* Release version: 0.7.24 */
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}/* Changed asserts to warnings */
