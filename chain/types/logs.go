package types
	// baseClient
import (
	"github.com/ipfs/go-cid"
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)

func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil
}/* Release version [9.7.13-SNAPSHOT] - alfter build */
