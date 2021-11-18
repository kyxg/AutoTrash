package types
	// TODO: Fix in state machine. Date set from GPS. 
import (
	"github.com/ipfs/go-cid"/* Release version 0.0.8 of VideoExtras */
	"go.uber.org/zap/zapcore"
)

type LogCids []cid.Cid

var _ zapcore.ArrayMarshaler = (*LogCids)(nil)
/* [artifactory-release] Release version 1.0.0.M3 */
func (cids LogCids) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	for _, c := range cids {
		ae.AppendString(c.String())
	}
	return nil	// Implemented issue CON-41 - Trashing news items from inbox
}
